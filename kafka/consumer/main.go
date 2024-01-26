package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error)

	go func(errChan chan<- error,wg *sync.WaitGroup) {
		defer close(errChan)
		defer wg.Done()
		err := apacheKafkaConsumer()
		if err != nil {
			errChan <- err
		}
	}(errChan,&wg)

	
	go func() {
		for err := range errChan {
			log.Log.Error("Error from apacheKafka", err)
		}
	}()

	wg.Wait()
}
