// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v int32) *DescribeInstanceIdsRequest
	GetEdition() *int32
	SetInstanceIds(v []*string) *DescribeInstanceIdsRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribeInstanceIdsRequest
	GetResourceGroupId() *string
}

type DescribeInstanceIdsRequest struct {
	// The type of the instance to query. Valid values:
	//
	// 	- **0**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Insurance mitigation plan
	//
	// 	- **1**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Unlimited mitigation plan
	//
	// 	- **2**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Chinese Mainland Acceleration (CMA) mitigation plan
	//
	// 	- **3**: Anti-DDoS Proxy (Outside Chinese Mainland) instance of the Secure Chinese Mainland Acceleration (Sec-CMA) mitigation plan
	//
	// 	- **9**: Anti-DDoS Proxy (Chinese Mainland) instance of the Profession mitigation plan
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The IDs of instances to query.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIdsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsRequest) GetEdition() *int32 {
	return s.Edition
}

func (s *DescribeInstanceIdsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstanceIdsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceIdsRequest) SetEdition(v int32) *DescribeInstanceIdsRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceIdsRequest) SetInstanceIds(v []*string) *DescribeInstanceIdsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstanceIdsRequest) SetResourceGroupId(v string) *DescribeInstanceIdsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceIdsRequest) Validate() error {
	return dara.Validate(s)
}
