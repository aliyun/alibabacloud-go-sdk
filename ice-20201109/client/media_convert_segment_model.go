// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertSegment interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *MediaConvertSegment
	GetDuration() *int32
	SetForceSegTime(v string) *MediaConvertSegment
	GetForceSegTime() *string
}

type MediaConvertSegment struct {
	// The length of the segment.
	//
	// 	- Unit: seconds.
	//
	// 	- Valid values: [1,60].
	//
	// 	- Default value: 10. A value of 10 will create segments at the 10s, 20s, 30s, and 40s marks of the video.
	//
	// example:
	//
	// 10
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The points in time at which the video is forcibly segmented. Separate multiple points with commas (,). You can specify up to 10 points.
	//
	// 	- Format: {Point in time},{Point in time},{Point in time}.
	//
	// 	- Type: decimal. This parameter supports up to three decimal places.
	//
	// 	- Unit: seconds.
	//
	// 	- Example: 1,2,4,6,10,14,18, which specifies that the video is forcibly segmented at the 1st, 2nd, 4th, 6th, 10th, 14th, and 18th seconds.
	//
	// example:
	//
	// 1,2
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s MediaConvertSegment) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertSegment) GoString() string {
	return s.String()
}

func (s *MediaConvertSegment) GetDuration() *int32 {
	return s.Duration
}

func (s *MediaConvertSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *MediaConvertSegment) SetDuration(v int32) *MediaConvertSegment {
	s.Duration = &v
	return s
}

func (s *MediaConvertSegment) SetForceSegTime(v string) *MediaConvertSegment {
	s.ForceSegTime = &v
	return s
}

func (s *MediaConvertSegment) Validate() error {
	return dara.Validate(s)
}
