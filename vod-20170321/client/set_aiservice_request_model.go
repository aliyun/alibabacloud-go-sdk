// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v string) *SetAIServiceRequest
	GetOperation() *string
	SetOwnerAccount(v string) *SetAIServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetAIServiceRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SetAIServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetAIServiceRequest
	GetResourceOwnerId() *string
	SetTypes(v string) *SetAIServiceRequest
	GetTypes() *string
}

type SetAIServiceRequest struct {
	// This parameter is required.
	Operation            *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s SetAIServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAIServiceRequest) GoString() string {
	return s.String()
}

func (s *SetAIServiceRequest) GetOperation() *string {
	return s.Operation
}

func (s *SetAIServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetAIServiceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetAIServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetAIServiceRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetAIServiceRequest) GetTypes() *string {
	return s.Types
}

func (s *SetAIServiceRequest) SetOperation(v string) *SetAIServiceRequest {
	s.Operation = &v
	return s
}

func (s *SetAIServiceRequest) SetOwnerAccount(v string) *SetAIServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetAIServiceRequest) SetOwnerId(v string) *SetAIServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAIServiceRequest) SetResourceOwnerAccount(v string) *SetAIServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetAIServiceRequest) SetResourceOwnerId(v string) *SetAIServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetAIServiceRequest) SetTypes(v string) *SetAIServiceRequest {
	s.Types = &v
	return s
}

func (s *SetAIServiceRequest) Validate() error {
	return dara.Validate(s)
}
