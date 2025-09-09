// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyUserScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPolicyUserScopeRequest
	GetInstanceId() *string
	SetPolicyId(v string) *GetPolicyUserScopeRequest
	GetPolicyId() *string
	SetRegionId(v string) *GetPolicyUserScopeRequest
	GetRegionId() *string
}

type GetPolicyUserScopeRequest struct {
	// The ID of the bastion host to which the control policy to query belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastion-xxxx-xx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to query.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host to which the control policy to query belongs.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolicyUserScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyUserScopeRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyUserScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPolicyUserScopeRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyUserScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyUserScopeRequest) SetInstanceId(v string) *GetPolicyUserScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPolicyUserScopeRequest) SetPolicyId(v string) *GetPolicyUserScopeRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyUserScopeRequest) SetRegionId(v string) *GetPolicyUserScopeRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyUserScopeRequest) Validate() error {
	return dara.Validate(s)
}
