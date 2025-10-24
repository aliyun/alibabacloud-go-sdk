// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateCustomGroupRequest
	GetAlgorithm() *string
	SetCustomGroupDescription(v string) *CreateCustomGroupRequest
	GetCustomGroupDescription() *string
	SetCustomGroupName(v string) *CreateCustomGroupRequest
	GetCustomGroupName() *string
	SetOwnerAccount(v string) *CreateCustomGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCustomGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCustomGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCustomGroupRequest
	GetResourceOwnerId() *int64
}

type CreateCustomGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm              *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CustomGroupDescription *string `json:"CustomGroupDescription,omitempty" xml:"CustomGroupDescription,omitempty"`
	// This parameter is required.
	CustomGroupName      *string `json:"CustomGroupName,omitempty" xml:"CustomGroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateCustomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomGroupRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateCustomGroupRequest) GetCustomGroupDescription() *string {
	return s.CustomGroupDescription
}

func (s *CreateCustomGroupRequest) GetCustomGroupName() *string {
	return s.CustomGroupName
}

func (s *CreateCustomGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCustomGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCustomGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCustomGroupRequest) SetAlgorithm(v string) *CreateCustomGroupRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCustomGroupRequest) SetCustomGroupDescription(v string) *CreateCustomGroupRequest {
	s.CustomGroupDescription = &v
	return s
}

func (s *CreateCustomGroupRequest) SetCustomGroupName(v string) *CreateCustomGroupRequest {
	s.CustomGroupName = &v
	return s
}

func (s *CreateCustomGroupRequest) SetOwnerAccount(v string) *CreateCustomGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCustomGroupRequest) SetOwnerId(v int64) *CreateCustomGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomGroupRequest) SetResourceOwnerAccount(v string) *CreateCustomGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCustomGroupRequest) SetResourceOwnerId(v int64) *CreateCustomGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCustomGroupRequest) Validate() error {
	return dara.Validate(s)
}
