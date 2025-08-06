// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetAIServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAIServiceRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetAIServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAIServiceRequest
	GetResourceOwnerId() *string
	SetTypes(v string) *GetAIServiceRequest
	GetTypes() *string
}

type GetAIServiceRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s GetAIServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIServiceRequest) GoString() string {
	return s.String()
}

func (s *GetAIServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAIServiceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAIServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAIServiceRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAIServiceRequest) GetTypes() *string {
	return s.Types
}

func (s *GetAIServiceRequest) SetOwnerAccount(v string) *GetAIServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIServiceRequest) SetOwnerId(v string) *GetAIServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIServiceRequest) SetResourceOwnerAccount(v string) *GetAIServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIServiceRequest) SetResourceOwnerId(v string) *GetAIServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAIServiceRequest) SetTypes(v string) *GetAIServiceRequest {
	s.Types = &v
	return s
}

func (s *GetAIServiceRequest) Validate() error {
	return dara.Validate(s)
}
