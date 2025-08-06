// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStatisTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListAIStatisTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIStatisTypeRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIStatisTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIStatisTypeRequest
	GetResourceOwnerId() *string
}

type ListAIStatisTypeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIStatisTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIStatisTypeRequest) GoString() string {
	return s.String()
}

func (s *ListAIStatisTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIStatisTypeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIStatisTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIStatisTypeRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIStatisTypeRequest) SetOwnerAccount(v string) *ListAIStatisTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIStatisTypeRequest) SetOwnerId(v string) *ListAIStatisTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIStatisTypeRequest) SetResourceOwnerAccount(v string) *ListAIStatisTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIStatisTypeRequest) SetResourceOwnerId(v string) *ListAIStatisTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIStatisTypeRequest) Validate() error {
	return dara.Validate(s)
}
