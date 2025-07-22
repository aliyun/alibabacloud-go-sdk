// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallList(v []*DescribeCallListResponseBodyCallList) *DescribeCallListResponseBody
	GetCallList() []*DescribeCallListResponseBodyCallList
	SetPageNo(v int32) *DescribeCallListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeCallListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCallListResponseBody
	GetRequestId() *string
	SetTotalCnt(v int32) *DescribeCallListResponseBody
	GetTotalCnt() *int32
}

type DescribeCallListResponseBody struct {
	CallList []*DescribeCallListResponseBodyCallList `json:"CallList,omitempty" xml:"CallList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeCallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBody) GetCallList() []*DescribeCallListResponseBodyCallList {
	return s.CallList
}

func (s *DescribeCallListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCallListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCallListResponseBody) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeCallListResponseBody) SetCallList(v []*DescribeCallListResponseBodyCallList) *DescribeCallListResponseBody {
	s.CallList = v
	return s
}

func (s *DescribeCallListResponseBody) SetPageNo(v int32) *DescribeCallListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListResponseBody) SetPageSize(v int32) *DescribeCallListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListResponseBody) SetRequestId(v string) *DescribeCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallListResponseBody) SetTotalCnt(v int32) *DescribeCallListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeCallListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCallListResponseBodyCallList struct {
	// App ID。
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BadExpUserCnt *int32 `json:"BadExpUserCnt,omitempty" xml:"BadExpUserCnt,omitempty"`
	// example:
	//
	// OUT
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 904
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 5
	UserCnt *int32 `json:"UserCnt,omitempty" xml:"UserCnt,omitempty"`
}

func (s DescribeCallListResponseBodyCallList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallListResponseBodyCallList) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBodyCallList) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCallListResponseBodyCallList) GetBadExpUserCnt() *int32 {
	return s.BadExpUserCnt
}

func (s *DescribeCallListResponseBodyCallList) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeCallListResponseBodyCallList) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCallListResponseBodyCallList) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeCallListResponseBodyCallList) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeCallListResponseBodyCallList) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeCallListResponseBodyCallList) GetUserCnt() *int32 {
	return s.UserCnt
}

func (s *DescribeCallListResponseBodyCallList) SetAppId(v string) *DescribeCallListResponseBodyCallList {
	s.AppId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetBadExpUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.BadExpUserCnt = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCallStatus(v string) *DescribeCallListResponseBodyCallList {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetChannelId(v string) *DescribeCallListResponseBodyCallList {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCreatedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDestroyedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDuration(v int64) *DescribeCallListResponseBodyCallList {
	s.Duration = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.UserCnt = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) Validate() error {
	return dara.Validate(s)
}
