// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveGrtnDurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveGrtnDurationResponseBody
	GetRequestId() *string
	SetStreamDetailData(v *DescribeLiveGrtnDurationResponseBodyStreamDetailData) *DescribeLiveGrtnDurationResponseBody
	GetStreamDetailData() *DescribeLiveGrtnDurationResponseBodyStreamDetailData
}

type DescribeLiveGrtnDurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4B460F8B-993C-4F48-B98A-910811DEBFEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the co-streaming usage data.
	StreamDetailData *DescribeLiveGrtnDurationResponseBodyStreamDetailData `json:"StreamDetailData,omitempty" xml:"StreamDetailData,omitempty" type:"Struct"`
}

func (s DescribeLiveGrtnDurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveGrtnDurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveGrtnDurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveGrtnDurationResponseBody) GetStreamDetailData() *DescribeLiveGrtnDurationResponseBodyStreamDetailData {
	return s.StreamDetailData
}

func (s *DescribeLiveGrtnDurationResponseBody) SetRequestId(v string) *DescribeLiveGrtnDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBody) SetStreamDetailData(v *DescribeLiveGrtnDurationResponseBodyStreamDetailData) *DescribeLiveGrtnDurationResponseBody {
	s.StreamDetailData = v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveGrtnDurationResponseBodyStreamDetailData struct {
	StreamData []*DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData `json:"StreamData,omitempty" xml:"StreamData,omitempty" type:"Repeated"`
}

func (s DescribeLiveGrtnDurationResponseBodyStreamDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveGrtnDurationResponseBodyStreamDetailData) GoString() string {
	return s.String()
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailData) GetStreamData() []*DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	return s.StreamData
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailData) SetStreamData(v []*DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) *DescribeLiveGrtnDurationResponseBodyStreamDetailData {
	s.StreamData = v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData struct {
	// The ID of the application.
	//
	// example:
	//
	// 4346289a-a790-4869-9e23-22766d5e****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The co-streaming duration. Unit: minutes.
	//
	// example:
	//
	// 30
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The media specification. Valid values:
	//
	// 	- 0: audio-only. This is a basic specification.
	//
	// 	- 480P: standard definition (SD). The video resolution is 640 × 480 or lower.
	//
	// 	- 720P: high definition (HD). The video resolution is 1280 × 720 or lower.
	//
	// 	- 1080P: full HD. The video resolution is 1920 × 1080 or lower.
	//
	// example:
	//
	// 480P
	MediaProfile *string `json:"MediaProfile,omitempty" xml:"MediaProfile,omitempty"`
	// The media type. Valid values:
	//
	// 	- audio
	//
	// 	- video
	//
	// example:
	//
	// audio
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2022-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GoString() string {
	return s.String()
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GetMediaProfile() *string {
	return s.MediaProfile
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GetMediaType() *string {
	return s.MediaType
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) SetAppId(v string) *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	s.AppId = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) SetDuration(v int64) *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	s.Duration = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) SetMediaProfile(v string) *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	s.MediaProfile = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) SetMediaType(v string) *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	s.MediaType = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) SetTimeStamp(v string) *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveGrtnDurationResponseBodyStreamDetailDataStreamData) Validate() error {
	return dara.Validate(s)
}
