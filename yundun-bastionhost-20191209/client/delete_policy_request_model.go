// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeletePolicyRequest
	GetInstanceId() *string
	SetPolicyId(v string) *DeletePolicyRequest
	GetPolicyId() *string
	SetProjectId(v int64) *DeletePolicyRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeletePolicyRequest
	GetRegionId() *string
}

type DeletePolicyRequest struct {
	// The instance ID of the bastion host to which the control policy to be deleted belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy to delete.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query this parameter.
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
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeletePolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePolicyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeletePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolicyRequest) SetInstanceId(v string) *DeletePolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicyId(v string) *DeletePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyRequest) SetProjectId(v int64) *DeletePolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *DeletePolicyRequest) SetRegionId(v string) *DeletePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
