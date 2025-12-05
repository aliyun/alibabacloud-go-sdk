// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ListVpcConfigsRequest
	GetConfigId() *string
	SetConfigName(v string) *ListVpcConfigsRequest
	GetConfigName() *string
	SetPageNumber(v int32) *ListVpcConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVpcConfigsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVpcConfigsRequest
	GetRegionId() *string
	SetVpcId(v string) *ListVpcConfigsRequest
	GetVpcId() *string
}

type ListVpcConfigsRequest struct {
	// example:
	//
	// 05AHW
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// my-vpc-config-info
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vpc-bp11w3qellkjwnhqxzmt2
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcConfigsRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListVpcConfigsRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListVpcConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVpcConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVpcConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcConfigsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcConfigsRequest) SetConfigId(v string) *ListVpcConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListVpcConfigsRequest) SetConfigName(v string) *ListVpcConfigsRequest {
	s.ConfigName = &v
	return s
}

func (s *ListVpcConfigsRequest) SetPageNumber(v int32) *ListVpcConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpcConfigsRequest) SetPageSize(v int32) *ListVpcConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcConfigsRequest) SetRegionId(v string) *ListVpcConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcConfigsRequest) SetVpcId(v string) *ListVpcConfigsRequest {
	s.VpcId = &v
	return s
}

func (s *ListVpcConfigsRequest) Validate() error {
	return dara.Validate(s)
}
