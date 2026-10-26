// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreatePolicyRequest
	GetComment() *string
	SetInstanceId(v string) *CreatePolicyRequest
	GetInstanceId() *string
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
	SetPriority(v string) *CreatePolicyRequest
	GetPriority() *string
	SetProjectId(v int64) *CreatePolicyRequest
	GetProjectId() *int64
	SetRegionId(v string) *CreatePolicyRequest
	GetRegionId() *string
}

type CreatePolicyRequest struct {
	// The description of the control policy. Maximum length: 500 characters.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host instance for which you want to create a control policy.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-20p364c1w0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the control policy. Maximum length: 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// policytest
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The priority of the control policy.
	//
	// - Valid values: 1 to 100. Default value: 1, which indicates the policy priority.
	//
	// - Different control policies can have the same priority. If multiple control policies have the same priority, the most recently created policy has the policy priority. Within a single policy, if the same command is configured in both command control and command approval, the priority from highest to lowest is: reject, allow, and approval.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host for which you want to create a control policy.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetComment() *string {
	return s.Comment
}

func (s *CreatePolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreatePolicyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolicyRequest) SetComment(v string) *CreatePolicyRequest {
	s.Comment = &v
	return s
}

func (s *CreatePolicyRequest) SetInstanceId(v string) *CreatePolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetPriority(v string) *CreatePolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreatePolicyRequest) SetProjectId(v int64) *CreatePolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *CreatePolicyRequest) SetRegionId(v string) *CreatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
