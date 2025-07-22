// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelList(v *DescribeRtcChannelListResponseBodyChannelList) *DescribeRtcChannelListResponseBody
	GetChannelList() *DescribeRtcChannelListResponseBodyChannelList
	SetPageNo(v int64) *DescribeRtcChannelListResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeRtcChannelListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeRtcChannelListResponseBody
	GetRequestId() *string
	SetTotalCnt(v int64) *DescribeRtcChannelListResponseBody
	GetTotalCnt() *int64
}

type DescribeRtcChannelListResponseBody struct {
	ChannelList *DescribeRtcChannelListResponseBodyChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	TotalCnt *int64 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeRtcChannelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBody) GetChannelList() *DescribeRtcChannelListResponseBodyChannelList {
	return s.ChannelList
}

func (s *DescribeRtcChannelListResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeRtcChannelListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRtcChannelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcChannelListResponseBody) GetTotalCnt() *int64 {
	return s.TotalCnt
}

func (s *DescribeRtcChannelListResponseBody) SetChannelList(v *DescribeRtcChannelListResponseBodyChannelList) *DescribeRtcChannelListResponseBody {
	s.ChannelList = v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageNo(v int64) *DescribeRtcChannelListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageSize(v int64) *DescribeRtcChannelListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetRequestId(v string) *DescribeRtcChannelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelListResponseBodyChannelList struct {
	ChannelList []*DescribeRtcChannelListResponseBodyChannelListChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelListResponseBodyChannelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelList) GetChannelList() []*DescribeRtcChannelListResponseBodyChannelListChannelList {
	return s.ChannelList
}

func (s *DescribeRtcChannelListResponseBodyChannelList) SetChannelList(v []*DescribeRtcChannelListResponseBodyChannelListChannelList) *DescribeRtcChannelListResponseBodyChannelList {
	s.ChannelList = v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelList) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelListResponseBodyChannelListChannelList struct {
	CallArea *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea `json:"CallArea,omitempty" xml:"CallArea,omitempty" type:"Struct"`
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2018-01-29T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2018-01-29T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	TotalUserCnt *int64 `json:"TotalUserCnt,omitempty" xml:"TotalUserCnt,omitempty"`
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) GetCallArea() *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea {
	return s.CallArea
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) GetTotalUserCnt() *int64 {
	return s.TotalUserCnt
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetCallArea(v *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.CallArea = v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetChannelId(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetEndTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetStartTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetTotalUserCnt(v int64) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.TotalUserCnt = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelListResponseBodyChannelListChannelListCallArea struct {
	CallArea []*string `json:"CallArea,omitempty" xml:"CallArea,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) GetCallArea() []*string {
	return s.CallArea
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) SetCallArea(v []*string) *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea {
	s.CallArea = v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) Validate() error {
	return dara.Validate(s)
}
