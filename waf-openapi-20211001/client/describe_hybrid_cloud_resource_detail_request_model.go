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
	// The back-to-origin address.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// Specifies whether to enable public cloud disaster recovery. Valid values:
	//
	// - **true**: Public cloud disaster recovery is enabled.
	//
	// - **false**: Public cloud disaster recovery is disabled.
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.*****.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
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
