// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetMessageCallbackEventListRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetMessageCallbackEventListRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetMessageCallbackEventListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetMessageCallbackEventListRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetMessageCallbackEventListRequest
	GetResourceRealOwnerId() *int64
}

type GetMessageCallbackEventListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s GetMessageCallbackEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackEventListRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackEventListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMessageCallbackEventListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMessageCallbackEventListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMessageCallbackEventListRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetMessageCallbackEventListRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetMessageCallbackEventListRequest) SetOwnerAccount(v string) *GetMessageCallbackEventListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMessageCallbackEventListRequest) SetOwnerId(v string) *GetMessageCallbackEventListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMessageCallbackEventListRequest) SetResourceOwnerAccount(v string) *GetMessageCallbackEventListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMessageCallbackEventListRequest) SetResourceOwnerId(v string) *GetMessageCallbackEventListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMessageCallbackEventListRequest) SetResourceRealOwnerId(v int64) *GetMessageCallbackEventListRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetMessageCallbackEventListRequest) Validate() error {
	return dara.Validate(s)
}
