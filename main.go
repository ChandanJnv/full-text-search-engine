package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/ChandanJnv/full-text-search-engine/utils"
)

func main() {
	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "small wild cat", "search query")

	flag.Parse()
	log.Println("Full text search is in progress")

	start := time.Now()

	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal("failed to load documents: ", err)
	}

	log.Printf("Loaded %d document in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %v document in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
