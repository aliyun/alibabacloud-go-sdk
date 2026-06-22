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
	// The timestamps of the frames within the current video element that match the search content. The unit is milliseconds.
	FrameTimes []*int64 `json:"FrameTimes,omitempty" xml:"FrameTimes,omitempty" type:"Repeated"`
	// The label information.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The time range of the video element. The array has a fixed length of 2. The two values represent the start time and end time in milliseconds.
	TimeRange []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
	// The index of the video stream to which the current video scene element belongs. This corresponds to the index in the [VideoStreams](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-detectmediameta?spm=a2c4g.11186623.0.0.463e600fIDdM8r#api-detail-40) array.
	//
	// example:
	//
	// 0
	VideoStreamIndex *int64 `json:"VideoStreamIndex,omitempty" xml:"VideoStreamIndex,omitempty"`
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
