// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiIpControlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *DescribeApiIpControlsRequest
	GetApiIds() *string
	SetGroupId(v string) *DescribeApiIpControlsRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeApiIpControlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiIpControlsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiIpControlsRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiIpControlsRequest
	GetStageName() *string
}

type DescribeApiIpControlsRequest struct {
	// The IDs of APIs. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 123,234
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group.
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
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiIpControlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiIpControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *DescribeApiIpControlsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiIpControlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiIpControlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiIpControlsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiIpControlsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiIpControlsRequest) SetApiIds(v string) *DescribeApiIpControlsRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetGroupId(v string) *DescribeApiIpControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetPageNumber(v int32) *DescribeApiIpControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetPageSize(v int32) *DescribeApiIpControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetSecurityToken(v string) *DescribeApiIpControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetStageName(v string) *DescribeApiIpControlsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiIpControlsRequest) Validate() error {
	return dara.Validate(s)
}
