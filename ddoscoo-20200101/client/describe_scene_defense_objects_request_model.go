// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefenseObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DescribeSceneDefenseObjectsRequest
	GetPolicyId() *string
	SetResourceGroupId(v string) *DescribeSceneDefenseObjectsRequest
	GetResourceGroupId() *string
}

type DescribeSceneDefenseObjectsRequest struct {
	// The ID of the policy that you want to query.
	//
	// > You can call the [DescribeSceneDefensePolicies](https://help.aliyun.com/document_detail/159382.html) operation to query the IDs of all policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47e07ebd-0ba5-4afc-957b-59d15b90****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSceneDefenseObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefenseObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeSceneDefenseObjectsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSceneDefenseObjectsRequest) SetPolicyId(v string) *DescribeSceneDefenseObjectsRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsRequest) SetResourceGroupId(v string) *DescribeSceneDefenseObjectsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsRequest) Validate() error {
	return dara.Validate(s)
}
