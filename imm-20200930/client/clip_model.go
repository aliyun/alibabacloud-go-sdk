// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClip interface {
	dara.Model
	String() string
	GoString() string
	SetTimeRange(v []*int64) *Clip
	GetTimeRange() []*int64
}

type Clip struct {
	TimeRange []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
}

func (s Clip) String() string {
	return dara.Prettify(s)
}

func (s Clip) GoString() string {
	return s.String()
}

func (s *Clip) GetTimeRange() []*int64 {
	return s.TimeRange
}

func (s *Clip) SetTimeRange(v []*int64) *Clip {
	s.TimeRange = v
	return s
}

func (s *Clip) Validate() error {
	return dara.Validate(s)
}
