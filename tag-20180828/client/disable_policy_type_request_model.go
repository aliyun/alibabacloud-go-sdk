// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolicyTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenType(v string) *DisablePolicyTypeRequest
	GetOpenType() *string
	SetOwnerAccount(v string) *DisablePolicyTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisablePolicyTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisablePolicyTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisablePolicyTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DisablePolicyTypeRequest
	GetResourceOwnerId() *string
	SetUserType(v string) *DisablePolicyTypeRequest
	GetUserType() *string
}

type DisablePolicyTypeRequest struct {
	OpenType     *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DisablePolicyTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DisablePolicyTypeRequest) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeRequest) GetOpenType() *string {
	return s.OpenType
}

func (s *DisablePolicyTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisablePolicyTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisablePolicyTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisablePolicyTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisablePolicyTypeRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DisablePolicyTypeRequest) GetUserType() *string {
	return s.UserType
}

func (s *DisablePolicyTypeRequest) SetOpenType(v string) *DisablePolicyTypeRequest {
	s.OpenType = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetOwnerAccount(v string) *DisablePolicyTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetOwnerId(v int64) *DisablePolicyTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetRegionId(v string) *DisablePolicyTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetResourceOwnerAccount(v string) *DisablePolicyTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetResourceOwnerId(v string) *DisablePolicyTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetUserType(v string) *DisablePolicyTypeRequest {
	s.UserType = &v
	return s
}

func (s *DisablePolicyTypeRequest) Validate() error {
	return dara.Validate(s)
}
