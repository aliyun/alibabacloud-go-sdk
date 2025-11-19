// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayVideoStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePlayVideoStatisResponseBody
	GetRequestId() *string
	SetVideoPlayStatisDetails(v *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) *DescribePlayVideoStatisResponseBody
	GetVideoPlayStatisDetails() *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails
}

type DescribePlayVideoStatisResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A92D3600-A3E7-43D6-****-B6E3B4A1FE6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The daily playback statistics on the video.
	VideoPlayStatisDetails *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails `json:"VideoPlayStatisDetails,omitempty" xml:"VideoPlayStatisDetails,omitempty" type:"Struct"`
}

func (s DescribePlayVideoStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayVideoStatisResponseBody) GetVideoPlayStatisDetails() *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails {
	return s.VideoPlayStatisDetails
}

func (s *DescribePlayVideoStatisResponseBody) SetRequestId(v string) *DescribePlayVideoStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBody) SetVideoPlayStatisDetails(v *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) *DescribePlayVideoStatisResponseBody {
	s.VideoPlayStatisDetails = v
	return s
}

func (s *DescribePlayVideoStatisResponseBody) Validate() error {
	if s.VideoPlayStatisDetails != nil {
		if err := s.VideoPlayStatisDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails struct {
	VideoPlayStatisDetail []*DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail `json:"VideoPlayStatisDetail,omitempty" xml:"VideoPlayStatisDetail,omitempty" type:"Repeated"`
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) GetVideoPlayStatisDetail() []*DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	return s.VideoPlayStatisDetail
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) SetVideoPlayStatisDetail(v []*DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails {
	s.VideoPlayStatisDetail = v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) Validate() error {
	if s.VideoPlayStatisDetail != nil {
		for _, item := range s.VideoPlayStatisDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail struct {
	// The date. The time follows the ISO 8601 standard in the *YYYY-MM-DD*T*hh:mm:ss	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 20170120
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The total playback duration. Unit: milliseconds.
	//
	// example:
	//
	// 967277
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	// The distribution of the playback duration.
	//
	// example:
	//
	// <=1m:79.2%;>1<=5m:16.7%;>5<=10m:4.2%
	PlayRange *string `json:"PlayRange,omitempty" xml:"PlayRange,omitempty"`
	// The video title.
	//
	// example:
	//
	// Four streams (one stream encrypted): LD-HLS + SD-MP4 + HD-HLS-encrypted + UHD-MP4
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The number of unique visitors.
	//
	// example:
	//
	// 1
	UV *string `json:"UV,omitempty" xml:"UV,omitempty"`
	// The number of video views.
	//
	// example:
	//
	// 24
	VV *string `json:"VV,omitempty" xml:"VV,omitempty"`
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetDate() *string {
	return s.Date
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetPlayDuration() *string {
	return s.PlayDuration
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetPlayRange() *string {
	return s.PlayRange
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetTitle() *string {
	return s.Title
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetUV() *string {
	return s.UV
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GetVV() *string {
	return s.VV
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetDate(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.Date = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetPlayDuration(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetPlayRange(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.PlayRange = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetTitle(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.Title = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetUV(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.UV = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetVV(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.VV = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) Validate() error {
	return dara.Validate(s)
}
