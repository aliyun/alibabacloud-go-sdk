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
	SetRegionId(v string) *ModifyPolicyRequest
	GetRegionId() *string
}

type ModifyPolicyRequest struct {
	// The new remarks of the control policy.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host to which the control policy to modify belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-09k1u8mv501
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The new name of the control policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The priority of the modified control policy. Valid values: 1 to 100. The smaller the value, the higher the priority. Default value: 1.
	//
	// example:
	//
	// 3
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the bastion host to which the control policy to modify belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ModifyPolicyRequest) SetRegionId(v string) *ModifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
