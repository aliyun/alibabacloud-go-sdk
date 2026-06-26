// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApikeyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeApikeyAttributeRequest
	GetApiKey() *string
	SetPageNumber(v int32) *DescribeApikeyAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApikeyAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApikeyAttributeRequest
	GetRegionId() *string
}

type DescribeApikeyAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6c4b1f0317cd4fd7a5b446d3503d**
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApikeyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApikeyAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeApikeyAttributeRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeApikeyAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApikeyAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApikeyAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApikeyAttributeRequest) SetApiKey(v string) *DescribeApikeyAttributeRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeApikeyAttributeRequest) SetPageNumber(v int32) *DescribeApikeyAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApikeyAttributeRequest) SetPageSize(v int32) *DescribeApikeyAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApikeyAttributeRequest) SetRegionId(v string) *DescribeApikeyAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApikeyAttributeRequest) Validate() error {
	return dara.Validate(s)
}
