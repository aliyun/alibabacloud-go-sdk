// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteReceiverRequest
	GetOwnerId() *int64
	SetReceiverId(v string) *DeleteReceiverRequest
	GetReceiverId() *string
	SetResourceOwnerAccount(v string) *DeleteReceiverRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteReceiverRequest
	GetResourceOwnerId() *int64
}

type DeleteReceiverRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Receiver list ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 144adfa772cfe47631de7e86d7da13ae
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiverRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteReceiverRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *DeleteReceiverRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteReceiverRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteReceiverRequest) SetOwnerId(v int64) *DeleteReceiverRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReceiverRequest) SetReceiverId(v string) *DeleteReceiverRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteReceiverRequest) SetResourceOwnerAccount(v string) *DeleteReceiverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReceiverRequest) SetResourceOwnerId(v int64) *DeleteReceiverRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteReceiverRequest) Validate() error {
	return dara.Validate(s)
}
