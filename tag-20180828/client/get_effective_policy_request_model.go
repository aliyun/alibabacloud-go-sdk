// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectivePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetEffectivePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetEffectivePolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEffectivePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetEffectivePolicyRequest
	GetResourceOwnerAccount() *string
	SetTagKeys(v []*string) *GetEffectivePolicyRequest
	GetTagKeys() []*string
	SetTargetId(v string) *GetEffectivePolicyRequest
	GetTargetId() *string
	SetTargetType(v string) *GetEffectivePolicyRequest
	GetTargetType() *string
}

type GetEffectivePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TagKeys              []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The ID of the object.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required.
	//
	// example:
	//
	// 154950938137****
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

func (s GetEffectivePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetEffectivePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEffectivePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEffectivePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEffectivePolicyRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *GetEffectivePolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GetEffectivePolicyRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *GetEffectivePolicyRequest) SetOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetOwnerId(v int64) *GetEffectivePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetRegionId(v string) *GetEffectivePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetResourceOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTagKeys(v []*string) *GetEffectivePolicyRequest {
	s.TagKeys = v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetId(v string) *GetEffectivePolicyRequest {
	s.TargetId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetType(v string) *GetEffectivePolicyRequest {
	s.TargetType = &v
	return s
}

func (s *GetEffectivePolicyRequest) Validate() error {
	return dara.Validate(s)
}
