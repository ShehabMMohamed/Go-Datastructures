package main

import (
	"fmt"
	"go-datastructures/internal/BinarySearchTrees"
	"go-datastructures/internal/Graphs"
	"go-datastructures/internal/HashTables"
	"go-datastructures/internal/Heaps"
	"go-datastructures/internal/Queues"
	"go-datastructures/internal/SinglyLinkedList"
	"go-datastructures/internal/Stacks"
	"go-datastructures/internal/Tries"
)

func main() {
	// Uncomment to see a datastructure simulation
	// Run_SinglyList()
	// Run_Stack()
	// Run_BST()
	// Run_Graph()
	// Run_Queue()
	// Run_Trie()
	// Run_MaxHeap()
	// Run_Map()
}

func Run_SinglyList() {
	fmt.Println("~~~ Singly Linked List ...")
	mylist := SinglyLinkedList.New()
	mylist.AddNodeAtTail(10)
	mylist.AddNodeAtTail(20)
	mylist.AddNodeAtTail(30)
	mylist.AddNodeAtHead(100)
	mylist.AddNodeAtHead(200)
	mylist.AddNodeAtTail(40)
	mylist.AddNode(-50, 3)
	mylist.Display()
	mylist.Reverse()
	mylist.Display()
	mylist.DeleteNode(5)
	mylist.Display()
}

func Run_BST() {
	fmt.Println("~~~ Binary Search Trees ...")
	myBST := BinarySearchTrees.New()
	/*
			10
		   /   \
		  5		50
			   /
			  30
			/   \
		   20   40 
	*/
	myBST.Add(10)
	myBST.Add(50)
	myBST.Add(30)
	myBST.Add(40)
	myBST.Add(20)
	myBST.Add(5)
	fmt.Printf("Searching for key %d, Result: %v.\n", 10, myBST.Search(10))
	fmt.Printf("Searching for key %d, Result: %v.\n", 20, myBST.Search(20))
	fmt.Printf("Searching for key %d, Result: %v.\n", 2, myBST.Search(2))
	fmt.Printf("Searching for key %d, Result: %v.\n", 5, myBST.Search(5))
	myBST.Delete(20)
	myBST.Delete(25)
	fmt.Printf("Searching for key %d, Result: %v.\n", 20, myBST.Search(20))
	fmt.Printf("Searching for key %d, Result: %v.\n", 25, myBST.Search(25))
	myBST.InOrder()
	myBST.PreOrder()
	myBST.PostOrder()
}

func Run_Stack() {
	fmt.Println("~~~ Stack ...")
	myStack := Stacks.New(3)
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Push("A")
	myStack.Push("B")
	myStack.Push("C")
	myStack.Push("D")
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
}

func Run_Queue() {
	fmt.Println("~~~ Queue ...")
	myQueue := Queues.New(100)
	myQueue.Enqueue("Customer 1")
	myQueue.Enqueue("Customer 2")
	myQueue.Enqueue("Customer 3")
	myQueue.Enqueue("Customer 4")
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
	myQueue.Dequeue()
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
	myQueue.Dequeue()
	myQueue.Enqueue("Customer 5")
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
}

func Run_Graph() {
	fmt.Println("~~~ Directed Graph ...")
	myGraph := Graphs.New(5)
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 4)
	myGraph.AddEdge(1, 5)
	myGraph.AddEdge(2, 3)
	myGraph.AddEdge(3, 4)
	myGraph.AddEdge(5, 6)
	myGraph.AddEdge(6, 7)
	myGraph.AddEdge(7, 8)
	myGraph.AddEdge(8, 9)
	myGraph.AddEdge(3, 10)

	fmt.Println("BFS...")
	myGraph.BFS(1)
	fmt.Println("DFS...")
	myGraph.DFS(1)
}
/*        2 - - > 3 -> 10
         /        |
		1  -> 4 <-
         \ 
		  5 -> 6 -> 7 -> 8 -> 9
*/

func Run_Trie() {
	fmt.Println("~~~ Trie ...")
	myTrie := Tries.New()
	myTrie.AddKey("She")
	myTrie.AddKey("Shehab")
	fmt.Printf("\"%s\" found ? %v\n", "She", myTrie.SearchKey("She"))
	fmt.Printf("\"%s\" found ? %v\n", "Shehab", myTrie.SearchKey("Shehab"))
	fmt.Printf("\"%s\" found ? %v\n", "he", myTrie.SearchKey("he"))
	myTrie.DeleteKey("She")
	fmt.Printf("\"%s\" found ? %v\n", "She", myTrie.SearchKey("She"))
}

func Run_MaxHeap() {
	fmt.Println("~~~ Max Heap ...")
	myMaxHeap := Heaps.New()
	A := []int{20, 2, 15, 3, 45}
	for i := 0; i < len(A); i++ {
		fmt.Printf("Pushing element %d to Heap.\n", A[i])
		myMaxHeap.Push(A[i])
	}
	fmt.Printf("Current size of Heap is %d.\n", myMaxHeap.GetSize())
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
	myMaxHeap.Pop()
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
	myMaxHeap.Pop()
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
	myMaxHeap.Pop()
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
	myMaxHeap.Pop()
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
	myMaxHeap.Pop()
	fmt.Printf("Current Top element in Heap: %d.\n", myMaxHeap.Top())
}

func Run_Map() {
	myMap := HashTables.New(10)
	myMap.Insert(10)
	myMap.Insert(7)
	myMap.Insert(5)
	myMap.Insert(2)
	myMap.Insert(30)
	myMap.Insert(3)
	myMap.Display()
	myMap.Remove(3)
	myMap.Remove(3)
	myMap.Remove(5)
	myMap.Display()
}
