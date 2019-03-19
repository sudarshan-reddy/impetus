package impetus

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	const Count = 10
	Delta := 100 * time.Millisecond
	ticker := NewImmediateTicker(Delta)
	t0 := time.Now()
	for i := 0; i < Count; i++ {
		<-ticker.C
	}
	ticker.Stop()
	t1 := time.Now()
	dt := t1.Sub(t0)
	target := Delta * Count
	slop := target * 2 / 10
	if dt < target-slop || dt > target+slop {
		t.Fatalf("%d %s ticks took %s, expected [%s,%s]", Count, Delta, dt, target-slop, target+slop)
	}

	select {
	case <-ticker.C:
		t.Fatal("tick did not stop")
	default:
	}
}
