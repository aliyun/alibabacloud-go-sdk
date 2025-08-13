// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DetachPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachPolicyRequest
	GetOwnerId() *int64
	SetPolicyId(v string) *DetachPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DetachPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachPolicyRequest
	GetResourceOwnerAccount() *string
	SetTargetId(v string) *DetachPolicyRequest
	GetTargetId() *string
	SetTargetType(v string) *DetachPolicyRequest
	GetTargetType() *string
}

type DetachPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-a3381efe2fe34a75****
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
	// > If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this parameter is required.
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
	// > If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this parameter is required. The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DetachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DetachPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachPolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DetachPolicyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DetachPolicyRequest) SetOwnerAccount(v string) *DetachPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachPolicyRequest) SetOwnerId(v int64) *DetachPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyId(v string) *DetachPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachPolicyRequest) SetRegionId(v string) *DetachPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DetachPolicyRequest) SetResourceOwnerAccount(v string) *DetachPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachPolicyRequest) SetTargetId(v string) *DetachPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *DetachPolicyRequest) SetTargetType(v string) *DetachPolicyRequest {
	s.TargetType = &v
	return s
}

func (s *DetachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
