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
	SetLabels(v []*Label) *SceneElement
	GetLabels() []*Label
	SetTimeRange(v []*int64) *SceneElement
	GetTimeRange() []*int64
	SetVideoStreamIndex(v int64) *SceneElement
	GetVideoStreamIndex() *int64
}

type SceneElement struct {
	// The time points of frames that match the searched content within the element. Unit: milliseconds.
	FrameTimes []*int64 `json:"FrameTimes,omitempty" xml:"FrameTimes,omitempty" type:"Repeated"`
	Labels     []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The time range of the element. The array length is fixed at 2. One is the start time and the other is the end time. Unit: milliseconds.
	TimeRange        []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
	VideoStreamIndex *int64   `json:"VideoStreamIndex,omitempty" xml:"VideoStreamIndex,omitempty"`
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

func (s *SceneElement) GetLabels() []*Label {
	return s.Labels
}

func (s *SceneElement) GetTimeRange() []*int64 {
	return s.TimeRange
}

func (s *SceneElement) GetVideoStreamIndex() *int64 {
	return s.VideoStreamIndex
}

func (s *SceneElement) SetFrameTimes(v []*int64) *SceneElement {
	s.FrameTimes = v
	return s
}

func (s *SceneElement) SetLabels(v []*Label) *SceneElement {
	s.Labels = v
	return s
}

func (s *SceneElement) SetTimeRange(v []*int64) *SceneElement {
	s.TimeRange = v
	return s
}

func (s *SceneElement) SetVideoStreamIndex(v int64) *SceneElement {
	s.VideoStreamIndex = &v
	return s
}

func (s *SceneElement) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
