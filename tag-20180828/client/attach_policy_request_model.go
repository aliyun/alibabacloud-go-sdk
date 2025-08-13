// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *AttachPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachPolicyRequest
	GetOwnerId() *int64
	SetPolicyId(v string) *AttachPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *AttachPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachPolicyRequest
	GetResourceOwnerAccount() *string
	SetTargetId(v string) *AttachPolicyRequest
	GetTargetId() *string
	SetTargetType(v string) *AttachPolicyRequest
	GetTargetType() *string
}

type AttachPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-de62a0bf400e4b69****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required.
	//
	// example:
	//
	// 151266687691****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required. The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AttachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachPolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *AttachPolicyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *AttachPolicyRequest) SetOwnerAccount(v string) *AttachPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachPolicyRequest) SetOwnerId(v int64) *AttachPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyId(v string) *AttachPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachPolicyRequest) SetRegionId(v string) *AttachPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *AttachPolicyRequest) SetResourceOwnerAccount(v string) *AttachPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachPolicyRequest) SetTargetId(v string) *AttachPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *AttachPolicyRequest) SetTargetType(v string) *AttachPolicyRequest {
	s.TargetType = &v
	return s
}

func (s *AttachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
