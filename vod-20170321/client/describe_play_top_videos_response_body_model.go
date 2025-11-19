// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayTopVideosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *DescribePlayTopVideosResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayTopVideosResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePlayTopVideosResponseBody
	GetRequestId() *string
	SetTopPlayVideos(v *DescribePlayTopVideosResponseBodyTopPlayVideos) *DescribePlayTopVideosResponseBody
	GetTopPlayVideos() *DescribePlayTopVideosResponseBodyTopPlayVideos
	SetTotalNum(v int64) *DescribePlayTopVideosResponseBody
	GetTotalNum() *int64
}

type DescribePlayTopVideosResponseBody struct {
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
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B0BCF9F-2FD5-4817-****-7BEBBE3AF90B"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The daily playback statistics on each top video.
	TopPlayVideos *DescribePlayTopVideosResponseBodyTopPlayVideos `json:"TopPlayVideos,omitempty" xml:"TopPlayVideos,omitempty" type:"Struct"`
	// The total number of entries that were collected in playback statistics on top videos.
	//
	// example:
	//
	// 2
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlayTopVideosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayTopVideosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayTopVideosResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayTopVideosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayTopVideosResponseBody) GetTopPlayVideos() *DescribePlayTopVideosResponseBodyTopPlayVideos {
	return s.TopPlayVideos
}

func (s *DescribePlayTopVideosResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribePlayTopVideosResponseBody) SetPageNo(v int64) *DescribePlayTopVideosResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetPageSize(v int64) *DescribePlayTopVideosResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetRequestId(v string) *DescribePlayTopVideosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetTopPlayVideos(v *DescribePlayTopVideosResponseBodyTopPlayVideos) *DescribePlayTopVideosResponseBody {
	s.TopPlayVideos = v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetTotalNum(v int64) *DescribePlayTopVideosResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) Validate() error {
	if s.TopPlayVideos != nil {
		if err := s.TopPlayVideos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlayTopVideosResponseBodyTopPlayVideos struct {
	TopPlayVideoStatis []*DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis `json:"TopPlayVideoStatis,omitempty" xml:"TopPlayVideoStatis,omitempty" type:"Repeated"`
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideos) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideos) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideos) GetTopPlayVideoStatis() []*DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	return s.TopPlayVideoStatis
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideos) SetTopPlayVideoStatis(v []*DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) *DescribePlayTopVideosResponseBodyTopPlayVideos {
	s.TopPlayVideoStatis = v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideos) Validate() error {
	if s.TopPlayVideoStatis != nil {
		for _, item := range s.TopPlayVideoStatis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis struct {
	// The total playback duration. Unit: milliseconds.
	//
	// example:
	//
	// 4640369
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	// The title of the video.
	//
	// example:
	//
	// Four streams (two streams encrypted): LD-HLS-encrypted + SD-MP4 + HD-H
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
	// 107
	VV *string `json:"VV,omitempty" xml:"VV,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 2a8d4cb9ecbb487681473a15****8fda
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GetPlayDuration() *string {
	return s.PlayDuration
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GetTitle() *string {
	return s.Title
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GetUV() *string {
	return s.UV
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GetVV() *string {
	return s.VV
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GetVideoId() *string {
	return s.VideoId
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetPlayDuration(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetTitle(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.Title = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetUV(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.UV = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetVV(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.VV = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetVideoId(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.VideoId = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) Validate() error {
	return dara.Validate(s)
}
