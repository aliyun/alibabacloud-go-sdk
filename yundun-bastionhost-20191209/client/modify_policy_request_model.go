// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyPolicyRequest
	GetComment() *string
	SetInstanceId(v string) *ModifyPolicyRequest
	GetInstanceId() *string
	SetPolicyId(v string) *ModifyPolicyRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ModifyPolicyRequest
	GetPolicyName() *string
	SetPriority(v string) *ModifyPolicyRequest
	GetPriority() *string
	SetProjectId(v int64) *ModifyPolicyRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyPolicyRequest
	GetRegionId() *string
}

type ModifyPolicyRequest struct {
	// The new description of the control policy.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the Bastionhost instance to which the control policy belongs. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-09k1u8mv501
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The new Policy Name of the control policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The new priority of the control policy. Valid values: 1 to 100. A smaller value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 3
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the control policy belongs. For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyPolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifyPolicyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ModifyPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPolicyRequest) SetComment(v string) *ModifyPolicyRequest {
	s.Comment = &v
	return s
}

func (s *ModifyPolicyRequest) SetInstanceId(v string) *ModifyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyId(v string) *ModifyPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyName(v string) *ModifyPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyPolicyRequest) SetPriority(v string) *ModifyPolicyRequest {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyRequest) SetProjectId(v int64) *ModifyPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyPolicyRequest) SetRegionId(v string) *ModifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
