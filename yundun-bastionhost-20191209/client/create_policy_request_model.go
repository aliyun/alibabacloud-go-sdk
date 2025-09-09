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
	SetRegionId(v string) *CreatePolicyRequest
	GetRegionId() *string
}

type CreatePolicyRequest struct {
	// The remarks of the control policy. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host for which you want to create a control policy.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-20p364c1w0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the control policy. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// policytest
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The priority of the control policy.
	//
	// 	- Valid values: 1 to 100. The default value is 1, which indicates the highest priority.
	//
	// 	- You can configure the same priority for different control policies. If multiple control policies have the same priority, the control policy that is created at the latest point in time has the highest priority. If a command control policy and a command approval policy contain the same commands, the commands are prioritized in descending order: reject, allow, and approve. In access control policies, a blacklist has a higher priority than a whitelist.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the bastion host for which you want to create a control policy.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *CreatePolicyRequest) SetRegionId(v string) *CreatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
