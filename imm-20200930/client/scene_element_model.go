// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSceneElement interface {
	dara.Model
	String() string
	GoString() string
	SetFrameTimes(v []*int64) *SceneElement
	GetFrameTimes() []*int64
	SetTimeRange(v []*int64) *SceneElement
	GetTimeRange() []*int64
}

type SceneElement struct {
	FrameTimes []*int64 `json:"FrameTimes,omitempty" xml:"FrameTimes,omitempty" type:"Repeated"`
	TimeRange  []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
}

func (s SceneElement) String() string {
	return dara.Prettify(s)
}

func (s SceneElement) GoString() string {
	return s.String()
}

func (s *SceneElement) GetFrameTimes() []*int64 {
	return s.FrameTimes
}

func (s *SceneElement) GetTimeRange() []*int64 {
	return s.TimeRange
}

func (s *SceneElement) SetFrameTimes(v []*int64) *SceneElement {
	s.FrameTimes = v
	return s
}

func (s *SceneElement) SetTimeRange(v []*int64) *SceneElement {
	s.TimeRange = v
	return s
}

func (s *SceneElement) Validate() error {
	return dara.Validate(s)
}
