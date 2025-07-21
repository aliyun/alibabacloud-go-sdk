// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMailAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMailAddressId(v int32) *DeleteMailAddressRequest
	GetMailAddressId() *int32
	SetOwnerId(v int64) *DeleteMailAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteMailAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMailAddressRequest
	GetResourceOwnerId() *int64
}

type DeleteMailAddressRequest struct {
	// Mail Address ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 23457
	MailAddressId        *int32  `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMailAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMailAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteMailAddressRequest) GetMailAddressId() *int32 {
	return s.MailAddressId
}

func (s *DeleteMailAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMailAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMailAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMailAddressRequest) SetMailAddressId(v int32) *DeleteMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *DeleteMailAddressRequest) SetOwnerId(v int64) *DeleteMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMailAddressRequest) SetResourceOwnerAccount(v string) *DeleteMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMailAddressRequest) SetResourceOwnerId(v int64) *DeleteMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMailAddressRequest) Validate() error {
	return dara.Validate(s)
}
