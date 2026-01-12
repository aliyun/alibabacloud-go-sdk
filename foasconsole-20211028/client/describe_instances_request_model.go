// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeInstancesRequest
	GetArchitectureType() *string
	SetChargeType(v string) *DescribeInstancesRequest
	GetChargeType() *string
	SetElastic(v bool) *DescribeInstancesRequest
	GetElastic() *bool
	SetInstanceId(v string) *DescribeInstancesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeInstancesRequest
	GetInstanceName() *string
	SetNamespaceName(v string) *DescribeInstancesRequest
	GetNamespaceName() *string
	SetPageIndex(v int32) *DescribeInstancesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeInstancesRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetTags(v []*DescribeInstancesRequestTags) *DescribeInstancesRequest
	GetTags() []*DescribeInstancesRequestTags
}

type DescribeInstancesRequest struct {
	// example:
	//
	// X86
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// true
	Elastic *bool `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// e2e-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// e2e-test-default
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region          *string                         `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*DescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesRequest) GetElastic() *bool {
	return s.Elastic
}

func (s *DescribeInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DescribeInstancesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequest) GetTags() []*DescribeInstancesRequestTags {
	return s.Tags
}

func (s *DescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetElastic(v bool) *DescribeInstancesRequest {
	s.Elastic = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetNamespaceName(v string) *DescribeInstancesRequest {
	s.NamespaceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageIndex(v int32) *DescribeInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegion(v string) *DescribeInstancesRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTags(v []*DescribeInstancesRequestTags) *DescribeInstancesRequest {
	s.Tags = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesRequestTags struct {
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// ys
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTags) SetKey(v string) *DescribeInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTags) SetValue(v string) *DescribeInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
