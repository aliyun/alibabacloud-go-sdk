// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPolicyRequest
	GetOwnerId() *int64
	SetPolicyContent(v string) *ModifyPolicyRequest
	GetPolicyContent() *string
	SetPolicyDesc(v string) *ModifyPolicyRequest
	GetPolicyDesc() *string
	SetPolicyId(v string) *ModifyPolicyRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ModifyPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *ModifyPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPolicyRequest
	GetResourceOwnerAccount() *string
}

type ModifyPolicyRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// 	- true: performs only a dry run.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The document of the tag policy.
	//
	// For more information about the syntax of a tag policy, see [Syntax of a tag policy](https://help.aliyun.com/document_detail/417436.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"tags":{"CostCenter":{"tag_value":{"@@assign":["Beijing","Shanghai"]},"tag_key":{"@@assign":"CostCenter"}}}}
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	//
	// The description must be 0 to 512 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// This is a tag policy example.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The ID of the tag policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-5732750813924f90****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (_).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ModifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPolicyRequest) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *ModifyPolicyRequest) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *ModifyPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPolicyRequest) SetDryRun(v bool) *ModifyPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyPolicyRequest) SetOwnerAccount(v string) *ModifyPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPolicyRequest) SetOwnerId(v int64) *ModifyPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyContent(v string) *ModifyPolicyRequest {
	s.PolicyContent = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyDesc(v string) *ModifyPolicyRequest {
	s.PolicyDesc = &v
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

func (s *ModifyPolicyRequest) SetRegionId(v string) *ModifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyRequest) SetResourceOwnerAccount(v string) *ModifyPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
