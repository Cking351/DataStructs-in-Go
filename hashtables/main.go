package main

import "fmt"

const ArraySize = 7

// HashTable
type HashTable struct {
	array [ArraySize]*Bucket
}
// Bucket struct
type Bucket struct {
	head *bucketNode
}
// Bucketnode struct
type bucketNode struct {
	key string
	next *bucketNode
}
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
} 
// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
// Delete
func (h (HashTable)) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}
// Bucket insert
func (b *Bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already exists!")
	}

}
// Bucket search
func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// Bucket delete
func (b *Bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			//delete
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}


func main() {
	hashTable := Init()
	list := []string {
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	
	for _, value := range list {
		hashTable.Insert(value)
	}

	hashTable.Delete("STAN")
	fmt.Println("STAN", hashTable.Search("STAN"))
	fmt.Println("KENNY", hashTable.Search("KENNY"))
}