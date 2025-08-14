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
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
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
