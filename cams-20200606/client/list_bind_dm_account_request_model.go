// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindDmAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListBindDmAccountRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListBindDmAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListBindDmAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListBindDmAccountRequest
	GetResourceOwnerId() *int64
}

type ListBindDmAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListBindDmAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindDmAccountRequest) GoString() string {
	return s.String()
}

func (s *ListBindDmAccountRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListBindDmAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListBindDmAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListBindDmAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListBindDmAccountRequest) SetCustSpaceId(v string) *ListBindDmAccountRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListBindDmAccountRequest) SetOwnerId(v int64) *ListBindDmAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ListBindDmAccountRequest) SetResourceOwnerAccount(v string) *ListBindDmAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListBindDmAccountRequest) SetResourceOwnerId(v int64) *ListBindDmAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListBindDmAccountRequest) Validate() error {
	return dara.Validate(s)
}
