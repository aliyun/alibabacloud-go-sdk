// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPPTDetails interface {
	dara.Model
	String() string
	GoString() string
	SetImagePath(v string) *PPTDetails
	GetImagePath() *string
	SetPPTShotIndex(v int64) *PPTDetails
	GetPPTShotIndex() *int64
	SetStartTime(v int64) *PPTDetails
	GetStartTime() *int64
}

type PPTDetails struct {
	// The URL of the captured slide image, which is stored in an Object Storage Service (OSS) bucket.
	ImagePath *string `json:"ImagePath,omitempty" xml:"ImagePath,omitempty"`
	// The zero-based index of the slide in the sequence of detected slides.
	PPTShotIndex *int64 `json:"PPTShotIndex,omitempty" xml:"PPTShotIndex,omitempty"`
	// The start time of the slide, in milliseconds, from the beginning of the video.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PPTDetails) String() string {
	return dara.Prettify(s)
}

func (s PPTDetails) GoString() string {
	return s.String()
}

func (s *PPTDetails) GetImagePath() *string {
	return s.ImagePath
}

func (s *PPTDetails) GetPPTShotIndex() *int64 {
	return s.PPTShotIndex
}

func (s *PPTDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *PPTDetails) SetImagePath(v string) *PPTDetails {
	s.ImagePath = &v
	return s
}

func (s *PPTDetails) SetPPTShotIndex(v int64) *PPTDetails {
	s.PPTShotIndex = &v
	return s
}

func (s *PPTDetails) SetStartTime(v int64) *PPTDetails {
	s.StartTime = &v
	return s
}

func (s *PPTDetails) Validate() error {
	return dara.Validate(s)
}
