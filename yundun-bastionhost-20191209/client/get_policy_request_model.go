// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPolicyRequest
	GetInstanceId() *string
	SetPolicyId(v string) *GetPolicyRequest
	GetPolicyId() *string
	SetProjectId(v int64) *GetPolicyRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetPolicyRequest
	GetRegionId() *string
}

type GetPolicyRequest struct {
	// The instance ID of the bastion host to which the control policy belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp2d3syb0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy to query.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
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

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyRequest) SetInstanceId(v string) *GetPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyId(v string) *GetPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRequest) SetProjectId(v int64) *GetPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPolicyRequest) SetRegionId(v string) *GetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
