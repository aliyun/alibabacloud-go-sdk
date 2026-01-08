// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappOpenStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetChatappOpenStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatappOpenStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappOpenStatusRequest
	GetResourceOwnerId() *int64
}

type GetChatappOpenStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetChatappOpenStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetChatappOpenStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappOpenStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappOpenStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappOpenStatusRequest) SetOwnerId(v int64) *GetChatappOpenStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappOpenStatusRequest) SetResourceOwnerAccount(v string) *GetChatappOpenStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappOpenStatusRequest) SetResourceOwnerId(v int64) *GetChatappOpenStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappOpenStatusRequest) Validate() error {
	return dara.Validate(s)
}
