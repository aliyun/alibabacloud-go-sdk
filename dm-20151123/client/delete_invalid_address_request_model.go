// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvalidAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteInvalidAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteInvalidAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInvalidAddressRequest
	GetResourceOwnerId() *int64
	SetToAddress(v string) *DeleteInvalidAddressRequest
	GetToAddress() *string
}

type DeleteInvalidAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Target address
	//
	// example:
	//
	// test1***@example.net
	ToAddress *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s DeleteInvalidAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvalidAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInvalidAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInvalidAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInvalidAddressRequest) GetToAddress() *string {
	return s.ToAddress
}

func (s *DeleteInvalidAddressRequest) SetOwnerId(v int64) *DeleteInvalidAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetResourceOwnerAccount(v string) *DeleteInvalidAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetResourceOwnerId(v int64) *DeleteInvalidAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInvalidAddressRequest) SetToAddress(v string) *DeleteInvalidAddressRequest {
	s.ToAddress = &v
	return s
}

func (s *DeleteInvalidAddressRequest) Validate() error {
	return dara.Validate(s)
}
