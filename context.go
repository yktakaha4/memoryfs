package memoryfs

import "time"

type context struct {
	provideTime func() time.Time
}
