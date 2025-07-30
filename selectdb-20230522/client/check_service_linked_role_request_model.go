// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest
	GetOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest
	GetResourceOwnerId() *int64
}

type CheckServiceLinkedRoleRequest struct {
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckServiceLinkedRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
