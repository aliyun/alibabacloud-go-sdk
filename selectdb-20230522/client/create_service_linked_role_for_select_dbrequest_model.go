// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForSelectDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CreateServiceLinkedRoleForSelectDBRequest
	GetOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateServiceLinkedRoleForSelectDBRequest
	GetResourceOwnerId() *int64
}

type CreateServiceLinkedRoleForSelectDBRequest struct {
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateServiceLinkedRoleForSelectDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) SetOwnerAccount(v string) *CreateServiceLinkedRoleForSelectDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) SetResourceOwnerId(v int64) *CreateServiceLinkedRoleForSelectDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBRequest) Validate() error {
	return dara.Validate(s)
}
