// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeInstancesRequest(v *DescribeInstancesRequestDescribeInstancesRequest) *DescribeInstancesRequest
	GetDescribeInstancesRequest() *DescribeInstancesRequestDescribeInstancesRequest
}

type DescribeInstancesRequest struct {
	// This parameter is required.
	DescribeInstancesRequest *DescribeInstancesRequestDescribeInstancesRequest `json:"DescribeInstancesRequest,omitempty" xml:"DescribeInstancesRequest,omitempty" type:"Struct"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetDescribeInstancesRequest() *DescribeInstancesRequestDescribeInstancesRequest {
	return s.DescribeInstancesRequest
}

func (s *DescribeInstancesRequest) SetDescribeInstancesRequest(v *DescribeInstancesRequestDescribeInstancesRequest) *DescribeInstancesRequest {
	s.DescribeInstancesRequest = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	if s.DescribeInstancesRequest != nil {
		if err := s.DescribeInstancesRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesRequestDescribeInstancesRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 223493C7-FCA9-13D4-B75B-AF8B32F4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// rg-***
	ResourceGroupId *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*DescribeInstancesRequestDescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequestDescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestDescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) GetTags() []*DescribeInstancesRequestDescribeInstancesRequestTags {
	return s.Tags
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetPageIndex(v int32) *DescribeInstancesRequestDescribeInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequestDescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetRegion(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetTags(v []*DescribeInstancesRequestDescribeInstancesRequestTags) *DescribeInstancesRequestDescribeInstancesRequest {
	s.Tags = v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) Validate() error {
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

type DescribeInstancesRequestDescribeInstancesRequestTags struct {
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestDescribeInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestDescribeInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) SetKey(v string) *DescribeInstancesRequestDescribeInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) SetValue(v string) *DescribeInstancesRequestDescribeInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
