// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DescribeCertsRequest
	GetAlgorithm() *string
	SetDomain(v string) *DescribeCertsRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeCertsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeCertsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCertsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeCertsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCertsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCertsRequest struct {
	// The type of the encryption algorithm. Valid values:
	//
	// 	- **NotSM2**: The encryption algorithm is not the SM2 algorithm. This is the default value.
	//
	// 	- **SM2**: The encryption algorithm is the SM2 algorithm.
	//
	// example:
	//
	// SM2
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-5yd****tb02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmoiy****p2oq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertsRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeCertsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCertsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCertsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCertsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCertsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCertsRequest) SetAlgorithm(v string) *DescribeCertsRequest {
	s.Algorithm = &v
	return s
}

func (s *DescribeCertsRequest) SetDomain(v string) *DescribeCertsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCertsRequest) SetInstanceId(v string) *DescribeCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCertsRequest) SetPageNumber(v int64) *DescribeCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCertsRequest) SetPageSize(v int64) *DescribeCertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCertsRequest) SetRegionId(v string) *DescribeCertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCertsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCertsRequest) Validate() error {
	return dara.Validate(s)
}
