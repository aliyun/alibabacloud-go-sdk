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
	SetProjectId(v int64) *GetPolicyUserScopeRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetPolicyUserScopeRequest
	GetRegionId() *string
}

type GetPolicyUserScopeRequest struct {
	// The instance ID of the bastion host to which the control policy belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastion-xxxx-xx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy to query.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host instance to which the control policy belongs.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetPolicyUserScopeRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *GetPolicyUserScopeRequest) SetProjectId(v int64) *GetPolicyUserScopeRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPolicyUserScopeRequest) SetRegionId(v string) *GetPolicyUserScopeRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyUserScopeRequest) Validate() error {
	return dara.Validate(s)
}
