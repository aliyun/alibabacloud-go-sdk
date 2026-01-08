// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViberServiceMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListViberServiceMessageRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListViberServiceMessageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListViberServiceMessageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListViberServiceMessageRequest
	GetResourceOwnerId() *int64
}

type ListViberServiceMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-8e4**96uhvk
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListViberServiceMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListViberServiceMessageRequest) GoString() string {
	return s.String()
}

func (s *ListViberServiceMessageRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListViberServiceMessageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListViberServiceMessageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListViberServiceMessageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListViberServiceMessageRequest) SetCustSpaceId(v string) *ListViberServiceMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListViberServiceMessageRequest) SetOwnerId(v int64) *ListViberServiceMessageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListViberServiceMessageRequest) SetResourceOwnerAccount(v string) *ListViberServiceMessageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListViberServiceMessageRequest) SetResourceOwnerId(v int64) *ListViberServiceMessageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListViberServiceMessageRequest) Validate() error {
	return dara.Validate(s)
}
