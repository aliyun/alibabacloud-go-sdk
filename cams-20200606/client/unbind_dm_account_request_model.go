// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDmAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UnbindDmAccountRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *UnbindDmAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnbindDmAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindDmAccountRequest
	GetResourceOwnerId() *int64
}

type UnbindDmAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-*
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindDmAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDmAccountRequest) GoString() string {
	return s.String()
}

func (s *UnbindDmAccountRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UnbindDmAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDmAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindDmAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindDmAccountRequest) SetCustSpaceId(v string) *UnbindDmAccountRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UnbindDmAccountRequest) SetOwnerId(v int64) *UnbindDmAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDmAccountRequest) SetResourceOwnerAccount(v string) *UnbindDmAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDmAccountRequest) SetResourceOwnerId(v int64) *UnbindDmAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindDmAccountRequest) Validate() error {
	return dara.Validate(s)
}
