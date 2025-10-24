// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteCustomViewRequest
	GetAlgorithm() *string
	SetCustomEntityId(v string) *DeleteCustomViewRequest
	GetCustomEntityId() *string
	SetCustomGroupId(v string) *DeleteCustomViewRequest
	GetCustomGroupId() *string
	SetCustomViewId(v string) *DeleteCustomViewRequest
	GetCustomViewId() *string
	SetOwnerAccount(v string) *DeleteCustomViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCustomViewRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCustomViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomViewRequest
	GetResourceOwnerId() *int64
}

type DeleteCustomViewRequest struct {
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
	CustomGroupId *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomViewId         *string `json:"CustomViewId,omitempty" xml:"CustomViewId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCustomViewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomViewRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomViewRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteCustomViewRequest) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *DeleteCustomViewRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *DeleteCustomViewRequest) GetCustomViewId() *string {
	return s.CustomViewId
}

func (s *DeleteCustomViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCustomViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomViewRequest) SetAlgorithm(v string) *DeleteCustomViewRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteCustomViewRequest) SetCustomEntityId(v string) *DeleteCustomViewRequest {
	s.CustomEntityId = &v
	return s
}

func (s *DeleteCustomViewRequest) SetCustomGroupId(v string) *DeleteCustomViewRequest {
	s.CustomGroupId = &v
	return s
}

func (s *DeleteCustomViewRequest) SetCustomViewId(v string) *DeleteCustomViewRequest {
	s.CustomViewId = &v
	return s
}

func (s *DeleteCustomViewRequest) SetOwnerAccount(v string) *DeleteCustomViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomViewRequest) SetOwnerId(v int64) *DeleteCustomViewRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomViewRequest) SetResourceOwnerAccount(v string) *DeleteCustomViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomViewRequest) SetResourceOwnerId(v int64) *DeleteCustomViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomViewRequest) Validate() error {
	return dara.Validate(s)
}
