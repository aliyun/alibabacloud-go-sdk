// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteCustomEntityRequest
	GetAlgorithm() *string
	SetCustomEntityId(v string) *DeleteCustomEntityRequest
	GetCustomEntityId() *string
	SetCustomGroupId(v string) *DeleteCustomEntityRequest
	GetCustomGroupId() *string
	SetOwnerAccount(v string) *DeleteCustomEntityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCustomEntityRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCustomEntityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomEntityRequest
	GetResourceOwnerId() *int64
}

type DeleteCustomEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomEntityId *string `json:"CustomEntityId,omitempty" xml:"CustomEntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomGroupId        *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCustomEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomEntityRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteCustomEntityRequest) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *DeleteCustomEntityRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *DeleteCustomEntityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCustomEntityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomEntityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomEntityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomEntityRequest) SetAlgorithm(v string) *DeleteCustomEntityRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetCustomEntityId(v string) *DeleteCustomEntityRequest {
	s.CustomEntityId = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetCustomGroupId(v string) *DeleteCustomEntityRequest {
	s.CustomGroupId = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetOwnerAccount(v string) *DeleteCustomEntityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetOwnerId(v int64) *DeleteCustomEntityRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetResourceOwnerAccount(v string) *DeleteCustomEntityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomEntityRequest) SetResourceOwnerId(v int64) *DeleteCustomEntityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomEntityRequest) Validate() error {
	return dara.Validate(s)
}
