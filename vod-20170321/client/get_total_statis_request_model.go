// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetTotalStatisRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetTotalStatisRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetTotalStatisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetTotalStatisRequest
	GetResourceOwnerId() *string
}

type GetTotalStatisRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetTotalStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTotalStatisRequest) GoString() string {
	return s.String()
}

func (s *GetTotalStatisRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetTotalStatisRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetTotalStatisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTotalStatisRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetTotalStatisRequest) SetOwnerAccount(v string) *GetTotalStatisRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetTotalStatisRequest) SetOwnerId(v string) *GetTotalStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTotalStatisRequest) SetResourceOwnerAccount(v string) *GetTotalStatisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTotalStatisRequest) SetResourceOwnerId(v string) *GetTotalStatisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTotalStatisRequest) Validate() error {
	return dara.Validate(s)
}
