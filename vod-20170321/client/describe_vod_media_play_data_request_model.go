// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMediaPlayDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DescribeVodMediaPlayDataRequest
	GetMediaId() *string
	SetOrderName(v string) *DescribeVodMediaPlayDataRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribeVodMediaPlayDataRequest
	GetOrderType() *string
	SetOs(v string) *DescribeVodMediaPlayDataRequest
	GetOs() *string
	SetPageNo(v int64) *DescribeVodMediaPlayDataRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeVodMediaPlayDataRequest
	GetPageSize() *int64
	SetPlayDate(v string) *DescribeVodMediaPlayDataRequest
	GetPlayDate() *string
	SetRegion(v string) *DescribeVodMediaPlayDataRequest
	GetRegion() *string
	SetTerminalType(v string) *DescribeVodMediaPlayDataRequest
	GetTerminalType() *string
}

type DescribeVodMediaPlayDataRequest struct {
	// The media ID, which is the audio or video ID (VideoId). Specify this parameter filtered query playback data for a specific media file. Only one media ID can be specified. You can obtain the media ID by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When you upload an audio or video file by calling the [CreateUploadVideo](~~CreateUploadVideo~~) operation, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](~~SearchMedia~~) operation filtered query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// 9ae2af636ca6c10412f44891fc****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The metric name. This parameter is used together with the `OrderType` parameter. Specify this parameter to sort the returned data in ascending or descending order by a specified metric. Valid values:
	//
	// - **PlaySuccessVv**: total plays.
	//
	// - **PlayPerVv**: average plays per user.
	//
	// - **PlayDuration**: total play duration.
	//
	// - **PlayDurationPerUv**: average play duration per user.
	//
	// example:
	//
	// PlaySuccessVv
	OrderName *string `json:"OrderName,omitempty" xml:"OrderName,omitempty"`
	// The sort order. This parameter is used together with the `OrderName` parameter. Specify this parameter to sort the returned data in ascending or descending order by a specified metric. Valid values:
	//
	// - **ASC**: ascending order. The returned data is sorted from smallest to largest.
	//
	// - **DESC**: descending order. The returned data is sorted from largest to smallest.
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The operating system of the playback device. Specify this parameter to perform a filtered query for playback data of all audio and video files by operating system. Valid values:
	//
	// - **Android**
	//
	// - **iOS**
	//
	// - **Windows**
	//
	// - **macOS**
	//
	// - **Linux**
	//
	// example:
	//
	// Android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The page number of the data to return. Specify this parameter to set the page from which data starts to be returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Specify this parameter to set the number of entries displayed on each page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The playback date. Unit: day. Format: yyyyMMdd.
	//
	// > - Only daily queries are supported.
	//
	// > - Only data within the last 30 days can be queried.
	//
	// example:
	//
	// 20240322
	PlayDate *string `json:"PlayDate,omitempty" xml:"PlayDate,omitempty"`
	// The service region. Specify this parameter to perform a filtered query for playback data of all audio and video files by service region. Valid values:
	//
	// - **cn-beijing**: China (Beijing)
	//
	// - **cn-shanghai**: China (Shanghai)
	//
	// - **cn-shenzhen**: China (Shenzhen)
	//
	// - **ap-northeast-1**: Japan (Tokyo)
	//
	// - **ap-southeast-1**: Singapore
	//
	// - **ap-southeast-5**: Indonesia (Jakarta)
	//
	// - **eu-central-1**: Germany (Frankfurt)
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The terminal type of the Player SDK. Specify this parameter to perform a filtered query for playback data of all audio and video files by terminal type. Valid values:
	//
	// - **Native**: Android Player SDK or iOS Player SDK.
	//
	// - **Web**: Web Player SDK.
	//
	// example:
	//
	// Native
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribeVodMediaPlayDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMediaPlayDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodMediaPlayDataRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DescribeVodMediaPlayDataRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribeVodMediaPlayDataRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeVodMediaPlayDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodMediaPlayDataRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeVodMediaPlayDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodMediaPlayDataRequest) GetPlayDate() *string {
	return s.PlayDate
}

func (s *DescribeVodMediaPlayDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodMediaPlayDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodMediaPlayDataRequest) SetMediaId(v string) *DescribeVodMediaPlayDataRequest {
	s.MediaId = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetOrderName(v string) *DescribeVodMediaPlayDataRequest {
	s.OrderName = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetOrderType(v string) *DescribeVodMediaPlayDataRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetOs(v string) *DescribeVodMediaPlayDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetPageNo(v int64) *DescribeVodMediaPlayDataRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetPageSize(v int64) *DescribeVodMediaPlayDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetPlayDate(v string) *DescribeVodMediaPlayDataRequest {
	s.PlayDate = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetRegion(v string) *DescribeVodMediaPlayDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) SetTerminalType(v string) *DescribeVodMediaPlayDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodMediaPlayDataRequest) Validate() error {
	return dara.Validate(s)
}
