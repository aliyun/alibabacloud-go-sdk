// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficControlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *DescribeApiTrafficControlsRequest
	GetApiIds() *string
	SetGroupId(v string) *DescribeApiTrafficControlsRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeApiTrafficControlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiTrafficControlsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiTrafficControlsRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiTrafficControlsRequest
	GetStageName() *string
}

type DescribeApiTrafficControlsRequest struct {
	// The IDs of APIs that you want to query. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 123,234
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The runtime environment of the API. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiTrafficControlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *DescribeApiTrafficControlsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiTrafficControlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiTrafficControlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiTrafficControlsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiTrafficControlsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiTrafficControlsRequest) SetApiIds(v string) *DescribeApiTrafficControlsRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetGroupId(v string) *DescribeApiTrafficControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetPageNumber(v int32) *DescribeApiTrafficControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetPageSize(v int32) *DescribeApiTrafficControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetSecurityToken(v string) *DescribeApiTrafficControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetStageName(v string) *DescribeApiTrafficControlsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) Validate() error {
	return dara.Validate(s)
}
