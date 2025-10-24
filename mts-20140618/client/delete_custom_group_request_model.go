// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteCustomGroupRequest
	GetAlgorithm() *string
	SetCustomGroupId(v string) *DeleteCustomGroupRequest
	GetCustomGroupId() *string
	SetOwnerAccount(v string) *DeleteCustomGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCustomGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCustomGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteCustomGroupRequest struct {
	// The custom type of the image library. Valid values:
	//
	// 	- landmark: custom landmarks.
	//
	// 	- object: custom objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the custom image library.
	//
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

func (s DeleteCustomGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomGroupRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteCustomGroupRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *DeleteCustomGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCustomGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomGroupRequest) SetAlgorithm(v string) *DeleteCustomGroupRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteCustomGroupRequest) SetCustomGroupId(v string) *DeleteCustomGroupRequest {
	s.CustomGroupId = &v
	return s
}

func (s *DeleteCustomGroupRequest) SetOwnerAccount(v string) *DeleteCustomGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomGroupRequest) SetOwnerId(v int64) *DeleteCustomGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomGroupRequest) SetResourceOwnerAccount(v string) *DeleteCustomGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomGroupRequest) SetResourceOwnerId(v int64) *DeleteCustomGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomGroupRequest) Validate() error {
	return dara.Validate(s)
}
