// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMediaPlayDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *DescribeVodMediaPlayDataResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeVodMediaPlayDataResponseBody
	GetPageSize() *int64
	SetQoeInfoList(v []*DescribeVodMediaPlayDataResponseBodyQoeInfoList) *DescribeVodMediaPlayDataResponseBody
	GetQoeInfoList() []*DescribeVodMediaPlayDataResponseBodyQoeInfoList
	SetRequestId(v string) *DescribeVodMediaPlayDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeVodMediaPlayDataResponseBody
	GetTotalCount() *int64
}

type DescribeVodMediaPlayDataResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data returned.
	QoeInfoList []*DescribeVodMediaPlayDataResponseBodyQoeInfoList `json:"QoeInfoList,omitempty" xml:"QoeInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodMediaPlayDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMediaPlayDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodMediaPlayDataResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeVodMediaPlayDataResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodMediaPlayDataResponseBody) GetQoeInfoList() []*DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	return s.QoeInfoList
}

func (s *DescribeVodMediaPlayDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodMediaPlayDataResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeVodMediaPlayDataResponseBody) SetPageNo(v int64) *DescribeVodMediaPlayDataResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBody) SetPageSize(v int64) *DescribeVodMediaPlayDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBody) SetQoeInfoList(v []*DescribeVodMediaPlayDataResponseBodyQoeInfoList) *DescribeVodMediaPlayDataResponseBody {
	s.QoeInfoList = v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBody) SetRequestId(v string) *DescribeVodMediaPlayDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBody) SetTotalCount(v int64) *DescribeVodMediaPlayDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodMediaPlayDataResponseBodyQoeInfoList struct {
	// The number of visits to the audio or video per day.
	//
	// example:
	//
	// 5
	DAU *float32 `json:"DAU,omitempty" xml:"DAU,omitempty"`
	// The ID of the media file (VideoId).
	//
	// example:
	//
	// 9ae2af636ca6c10412f44891fc****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The total playback duration of the audio or video. Unit: seconds.
	//
	// example:
	//
	// 2400
	PlayDuration *float32 `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	// The average playback duration of the audio or video per viewer. Unit: seconds.
	//
	// example:
	//
	// 120
	PlayDurationPerUv *float32 `json:"PlayDurationPerUv,omitempty" xml:"PlayDurationPerUv,omitempty"`
	// The average number of times that the audio or video was played per viewer.
	//
	// example:
	//
	// 4
	PlayPerVv *float32 `json:"PlayPerVv,omitempty" xml:"PlayPerVv,omitempty"`
	// The total number of times the audio or video has been played.
	//
	// example:
	//
	// 20
	PlaySuccessVv *float32 `json:"PlaySuccessVv,omitempty" xml:"PlaySuccessVv,omitempty"`
	// The duration of the audio or video file. Unit: seconds.
	//
	// example:
	//
	// 246
	VideoDuration *float32 `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	// The name of the audio or video file.
	//
	// example:
	//
	// title
	VideoTitle *string `json:"VideoTitle,omitempty" xml:"VideoTitle,omitempty"`
}

func (s DescribeVodMediaPlayDataResponseBodyQoeInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMediaPlayDataResponseBodyQoeInfoList) GoString() string {
	return s.String()
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetDAU() *float32 {
	return s.DAU
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetPlayDuration() *float32 {
	return s.PlayDuration
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetPlayDurationPerUv() *float32 {
	return s.PlayDurationPerUv
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetPlayPerVv() *float32 {
	return s.PlayPerVv
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetPlaySuccessVv() *float32 {
	return s.PlaySuccessVv
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetVideoDuration() *float32 {
	return s.VideoDuration
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) GetVideoTitle() *string {
	return s.VideoTitle
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetDAU(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.DAU = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetMediaId(v string) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.MediaId = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetPlayDuration(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.PlayDuration = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetPlayDurationPerUv(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.PlayDurationPerUv = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetPlayPerVv(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.PlayPerVv = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetPlaySuccessVv(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.PlaySuccessVv = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetVideoDuration(v float32) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.VideoDuration = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) SetVideoTitle(v string) *DescribeVodMediaPlayDataResponseBodyQoeInfoList {
	s.VideoTitle = &v
	return s
}

func (s *DescribeVodMediaPlayDataResponseBodyQoeInfoList) Validate() error {
	return dara.Validate(s)
}
