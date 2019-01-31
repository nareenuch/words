package words

import (
	// "fmt"
	"strings"
)

func WordCount(s string) map[string]int { // HL
	m :=map[string]int{}
s = strings.Replace(s,",","",-1)
s = strings.Replace(s,".","",-1)
arrs := strings.Split(s," ")
uarrs :=unique(arrs)
for _, word:= range uarrs {
	cntword :=0
	// fmt.Println("in loop 1", word)
	for _, word2:= range arrs {
		//fmt.Println("in loop", word)
		if word == word2 {
			cntword+=1
		}		
	}
	m[word] = cntword
}
	return m
} // HL

func unique(strSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range strSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}