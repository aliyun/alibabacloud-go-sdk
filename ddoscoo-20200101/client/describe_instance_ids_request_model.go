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
	// The edition of the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Valid values:
	//
	// - **0**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Insurance Edition.
	//
	// - **1**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Unlimited Edition.
	//
	// - **2**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Acceleration Line.
	//
	// - **3**: Anti-DDoS Pro or Anti-DDoS Premium (outside the Chinese mainland) Secure Acceleration Line.
	//
	// - **9**: Anti-DDoS Pro or Anti-DDoS Premium (the Chinese mainland) Professional Edition.
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The IDs of the Anti-DDoS Pro or Anti-DDoS Premium instances to query. Maximum value of N: 200.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the default resource group is used.
	//
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
