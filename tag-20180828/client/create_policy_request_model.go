// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreatePolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreatePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePolicyRequest
	GetOwnerId() *int64
	SetPolicyContent(v string) *CreatePolicyRequest
	GetPolicyContent() *string
	SetPolicyDesc(v string) *CreatePolicyRequest
	GetPolicyDesc() *string
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *CreatePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePolicyRequest
	GetResourceOwnerAccount() *string
	SetUserType(v string) *CreatePolicyRequest
	GetUserType() *string
}

type CreatePolicyRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// {"tags":{"CostCenter":{"tag_value":{"@@assign":["Beijing","Shanghai"]},"tag_key":{"@@assign":"CostCenter"}}}}
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	//
	// The description must be 0 to 512 characters in length.
	//
	// example:
	//
	// This is a tag policy example.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The name of the tag policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
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
	// The mode of the Tag Policy feature. Valid values:
	//
	// 	- USER: single-account mode. Set the value to USER if you use an Alibaba Cloud account or a member of a resource directory to call this API operation to create a tag policy for the Alibaba Cloud account or member.
	//
	// 	- RD: multi-account mode. Set the value to RD if you use the management account of a resource directory to call this API operation to create a tag policy for the resource directory.
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// example:
	//
	// RD
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreatePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePolicyRequest) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *CreatePolicyRequest) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePolicyRequest) GetUserType() *string {
	return s.UserType
}

func (s *CreatePolicyRequest) SetDryRun(v bool) *CreatePolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreatePolicyRequest) SetOwnerAccount(v string) *CreatePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePolicyRequest) SetOwnerId(v int64) *CreatePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyContent(v string) *CreatePolicyRequest {
	s.PolicyContent = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDesc(v string) *CreatePolicyRequest {
	s.PolicyDesc = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetRegionId(v string) *CreatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyRequest) SetResourceOwnerAccount(v string) *CreatePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePolicyRequest) SetUserType(v string) *CreatePolicyRequest {
	s.UserType = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
