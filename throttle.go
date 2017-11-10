package throttle

import "time"

type Throttle chan bool

// create throttle
func NewThrottle(limit uint) Throttle {
	var limiter Throttle

	// 0 is unlimited
	if limit == 0 {
		limiter = nil
	} else {
		limiter = make(Throttle, limit)
	}

	return limiter
}

// increase throttle count
func (t Throttle) Increase() {
	if t == nil {
		return
	}
	t <- true
}

// decrease throttle count
func (t Throttle) Decrease() {
	if t == nil {
		return
	}
	<-t
}

// keep current throttle count for milliseconds
func (t Throttle) Keep(msec uint) {
	if t == nil || msec == 0 {
		return
	}
	time.Sleep(time.Millisecond * time.Duration(msec))
}

// close channel of throttle
func (t Throttle) Close() {
	if t == nil {
		return
	}
	close(t)
}

