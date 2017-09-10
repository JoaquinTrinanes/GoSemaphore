//Package semaphore provides a very basic semaphore.
//
//
//Usage:
//
//	s := semaphore.SemInit(0) //sets semaphore to 0
//
//	s.Up() //increments count, n == 1
//
//	s.Down() //decrements count, n == 0
//
//	s.Down() //n == 0, so the goroutine blocks
package semaphore

import (
	"runtime"
	"sync"
)

//Semaphore is a struct containing the necessary (private) data.
type Semaphore struct {
	mutex sync.Mutex
	n     uint32
}

//SemInit creates a semaphore with the given value and returns a pointer to it.
//If not initialized, n equals 0.
func SemInit(N uint32) *Semaphore {
	/*var s *Semaphore = new(Semaphore)
	s.n = N*/
	return &Semaphore{n: N}
}

//Up increments the semaphore value, unlocking every possible goroutine previously blocked by it.
func (s *Semaphore) Up() {
	s.mutex.Lock()
	defer s.mutex.Unlock() //whatever happens, will always unlock at the end
	s.n++
}

//Down will block the caller goroutine if the semaphore value is 0, or decrement it otherwise. A blocked goroutine will attempt to perform the Down operation once it's unlocked.
func (s *Semaphore) Down() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for {
		if s.n > 0 {
			s.n--
			break
		} else {
			s.mutex.Unlock()
			//we give CPU usage to the other goroutines, and then check again the value
			runtime.Gosched() //thread yield
			s.mutex.Lock()
		}
	}

}

//TryDown will attempt to decrement the semaphore value. If the operation was a success
//it will return true. Otherwise, it won't do anything and return false. Note that the goroutine won't block under any circunstances when calling this method.
func (s *Semaphore) TryDown() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.n > 0 {
		s.n--
		return true
	}
	return false
}

//Value returns the semaphore current value.
func (s *Semaphore) Value() uint32 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.n
}
