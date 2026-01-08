// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessengerPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountIds(v []*string) *CreateMessengerPageRequest
	GetAdAccountIds() []*string
	SetAuthenticationCode(v string) *CreateMessengerPageRequest
	GetAuthenticationCode() *string
	SetBusinessId(v string) *CreateMessengerPageRequest
	GetBusinessId() *string
	SetCustSpaceId(v string) *CreateMessengerPageRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *CreateMessengerPageRequest
	GetOwnerId() *int64
	SetPageId(v string) *CreateMessengerPageRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *CreateMessengerPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMessengerPageRequest
	GetResourceOwnerId() *int64
}

type CreateMessengerPageRequest struct {
	// This parameter is required.
	AdAccountIds []*string `json:"AdAccountIds,omitempty" xml:"AdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// EAA****
	AuthenticationCode *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 293***
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-s***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19283***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMessengerPageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessengerPageRequest) GoString() string {
	return s.String()
}

func (s *CreateMessengerPageRequest) GetAdAccountIds() []*string {
	return s.AdAccountIds
}

func (s *CreateMessengerPageRequest) GetAuthenticationCode() *string {
	return s.AuthenticationCode
}

func (s *CreateMessengerPageRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateMessengerPageRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateMessengerPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMessengerPageRequest) GetPageId() *string {
	return s.PageId
}

func (s *CreateMessengerPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMessengerPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMessengerPageRequest) SetAdAccountIds(v []*string) *CreateMessengerPageRequest {
	s.AdAccountIds = v
	return s
}

func (s *CreateMessengerPageRequest) SetAuthenticationCode(v string) *CreateMessengerPageRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *CreateMessengerPageRequest) SetBusinessId(v string) *CreateMessengerPageRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateMessengerPageRequest) SetCustSpaceId(v string) *CreateMessengerPageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateMessengerPageRequest) SetOwnerId(v int64) *CreateMessengerPageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMessengerPageRequest) SetPageId(v string) *CreateMessengerPageRequest {
	s.PageId = &v
	return s
}

func (s *CreateMessengerPageRequest) SetResourceOwnerAccount(v string) *CreateMessengerPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMessengerPageRequest) SetResourceOwnerId(v int64) *CreateMessengerPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMessengerPageRequest) Validate() error {
	return dara.Validate(s)
}
