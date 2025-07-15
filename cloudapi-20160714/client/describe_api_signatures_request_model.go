// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiSignaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *DescribeApiSignaturesRequest
	GetApiIds() *string
	SetGroupId(v string) *DescribeApiSignaturesRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeApiSignaturesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiSignaturesRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiSignaturesRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiSignaturesRequest
	GetStageName() *string
}

type DescribeApiSignaturesRequest struct {
	// The IDs of the APIs that you want to query. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 123
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
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiSignaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiSignaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *DescribeApiSignaturesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiSignaturesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiSignaturesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiSignaturesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiSignaturesRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiSignaturesRequest) SetApiIds(v string) *DescribeApiSignaturesRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetGroupId(v string) *DescribeApiSignaturesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetPageNumber(v int32) *DescribeApiSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetPageSize(v int32) *DescribeApiSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetSecurityToken(v string) *DescribeApiSignaturesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetStageName(v string) *DescribeApiSignaturesRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiSignaturesRequest) Validate() error {
	return dara.Validate(s)
}
