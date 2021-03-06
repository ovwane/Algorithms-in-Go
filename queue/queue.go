package queue

import "sync"

//Queue 是一个具有队列api的接口。
//线程安全
type Queue interface {
	//Enqueue 往队列内添加数据。
	Enqueue(interface{})

	//Dequeue 删除队列中最早的数据。
	Dequeue() interface{}

	//IsEmpty 检查队列是否为空
	IsEmpty() bool

	//Size 返回队列内元素的数量。
	Size() int

	//Iterator 返回一个迭代器
	Iterator() Iterator
}

// New 返回一个队列接口，空的哟
func New() Queue {
	return &queue{}
}

type queue struct {
	sync.RWMutex
	first *node //指向最早的node
	last  *node //指向最新的node
	n     int   //记录长度
}

func (q *queue) Enqueue(item interface{}) {
	q.Lock()
	defer q.Unlock()
	oldLast := q.last
	q.last = newNode(item)
	if q.n == 0 { //当空队列enqueue入第一个node的时候。
		q.first = q.last //first和last都指向同一个node
	} else {
		oldLast.next = q.last
	}
	q.n++
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	q.Lock()
	defer q.Unlock()
	item := q.first.item
	q.first = q.first.next
	if q.first == nil { //如果连first都为nil，说明队列空了。
		q.last = nil //需要把last设置为nil
	}
	q.n--
	return item
}

func (q *queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.first == nil
}

func (q *queue) Size() int {
	q.RLock()
	defer q.RUnlock()
	return q.n
}

func (q *queue) Iterator() Iterator {
	q.RLock()
	defer q.RUnlock()
	return &chain{node: q.first}
}
