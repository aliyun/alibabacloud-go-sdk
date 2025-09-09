// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyAssetScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPolicyAssetScopeRequest
	GetInstanceId() *string
	SetPolicyId(v string) *GetPolicyAssetScopeRequest
	GetPolicyId() *string
	SetRegionId(v string) *GetPolicyAssetScopeRequest
	GetRegionId() *string
}

type GetPolicyAssetScopeRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The control policy ID.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolicyAssetScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPolicyAssetScopeRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyAssetScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyAssetScopeRequest) SetInstanceId(v string) *GetPolicyAssetScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPolicyAssetScopeRequest) SetPolicyId(v string) *GetPolicyAssetScopeRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyAssetScopeRequest) SetRegionId(v string) *GetPolicyAssetScopeRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyAssetScopeRequest) Validate() error {
	return dara.Validate(s)
}
