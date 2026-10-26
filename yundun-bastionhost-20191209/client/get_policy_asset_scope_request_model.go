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
	SetProjectId(v int64) *GetPolicyAssetScopeRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetPolicyAssetScopeRequest
	GetRegionId() *string
}

type GetPolicyAssetScopeRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetPolicyAssetScopeRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *GetPolicyAssetScopeRequest) SetProjectId(v int64) *GetPolicyAssetScopeRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPolicyAssetScopeRequest) SetRegionId(v string) *GetPolicyAssetScopeRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyAssetScopeRequest) Validate() error {
	return dara.Validate(s)
}
