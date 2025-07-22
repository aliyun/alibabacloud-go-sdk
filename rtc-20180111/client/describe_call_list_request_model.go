// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCallListRequest
	GetAppId() *string
	SetCallStatus(v string) *DescribeCallListRequest
	GetCallStatus() *string
	SetChannelId(v string) *DescribeCallListRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeCallListRequest
	GetEndTs() *int64
	SetOrderBy(v string) *DescribeCallListRequest
	GetOrderBy() *string
	SetPageNo(v int32) *DescribeCallListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeCallListRequest
	GetPageSize() *int32
	SetQueryMode(v string) *DescribeCallListRequest
	GetQueryMode() *string
	SetStartTs(v int64) *DescribeCallListRequest
	GetStartTs() *int64
	SetUserId(v string) *DescribeCallListRequest
	GetUserId() *string
}

type DescribeCallListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// OUT
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// example:
	//
	// BAD_EXP_USER_COUNT_DESC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ALL
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615806196
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCallListRequest) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeCallListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCallListRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeCallListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeCallListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCallListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCallListRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *DescribeCallListRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeCallListRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCallListRequest) SetAppId(v string) *DescribeCallListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallListRequest) SetCallStatus(v string) *DescribeCallListRequest {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListRequest) SetChannelId(v string) *DescribeCallListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListRequest) SetEndTs(v int64) *DescribeCallListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeCallListRequest) SetOrderBy(v string) *DescribeCallListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeCallListRequest) SetPageNo(v int32) *DescribeCallListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListRequest) SetPageSize(v int32) *DescribeCallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListRequest) SetQueryMode(v string) *DescribeCallListRequest {
	s.QueryMode = &v
	return s
}

func (s *DescribeCallListRequest) SetStartTs(v int64) *DescribeCallListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeCallListRequest) SetUserId(v string) *DescribeCallListRequest {
	s.UserId = &v
	return s
}

func (s *DescribeCallListRequest) Validate() error {
	return dara.Validate(s)
}
