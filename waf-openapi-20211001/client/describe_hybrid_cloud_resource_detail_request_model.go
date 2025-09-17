// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackend(v string) *DescribeHybridCloudResourceDetailRequest
	GetBackend() *string
	SetCnameEnabled(v bool) *DescribeHybridCloudResourceDetailRequest
	GetCnameEnabled() *bool
	SetDomain(v string) *DescribeHybridCloudResourceDetailRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeHybridCloudResourceDetailRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeHybridCloudResourceDetailRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeHybridCloudResourceDetailRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeHybridCloudResourceDetailRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourceDetailRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudResourceDetailRequest struct {
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.*****.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudResourceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailRequest) GetBackend() *string {
	return s.Backend
}

func (s *DescribeHybridCloudResourceDetailRequest) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *DescribeHybridCloudResourceDetailRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHybridCloudResourceDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudResourceDetailRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeHybridCloudResourceDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeHybridCloudResourceDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudResourceDetailRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudResourceDetailRequest) SetBackend(v string) *DescribeHybridCloudResourceDetailRequest {
	s.Backend = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetCnameEnabled(v bool) *DescribeHybridCloudResourceDetailRequest {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetDomain(v string) *DescribeHybridCloudResourceDetailRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetInstanceId(v string) *DescribeHybridCloudResourceDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetPageNumber(v int64) *DescribeHybridCloudResourceDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetPageSize(v int64) *DescribeHybridCloudResourceDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetRegionId(v string) *DescribeHybridCloudResourceDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourceDetailRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailRequest) Validate() error {
	return dara.Validate(s)
}
