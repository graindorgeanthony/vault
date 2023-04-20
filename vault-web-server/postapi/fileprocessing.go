package postapi

import "strings"
import "log"

type Chunk struct {
	Start int
	End   int
	Title string
	Text  string
}

func CreateChunks(fileContent string, window int, stride int, title string) []Chunk {
	sentences := strings.Split(fileContent, ".") // assuming sentences end with a period
	newData := make([]Chunk, 0)

	log.Println("length: ", len(sentences))
	log.Println("window: ", window)
	log.Println("stride: ",stride)

	if(len(sentences)-window <= 0) {
		newData = append(newData, Chunk{
			Start: 0,
			End:   len(fileContent),
			Title: title,
			Text:  fileContent,
		})
	} else {
		for i := 0; i < len(sentences)-window; i += stride {
			iEnd := i + window
			text := strings.Join(sentences[i:iEnd], ". ")
			start := 0
			end := 0
	
			if i > 0 {
				start = len(strings.Join(sentences[:i], ". ")) + 2 // +2 for the period and space
			}
	
			end = len(strings.Join(sentences[:iEnd], ". "))
	
			newData = append(newData, Chunk{
				Start: start,
				End:   end,
				Title: title,
				Text:  text,
			})
		}
	}

	

	return newData
}
