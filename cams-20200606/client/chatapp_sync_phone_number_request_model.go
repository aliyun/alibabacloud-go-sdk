// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappSyncPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappSyncPhoneNumberRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappSyncPhoneNumberRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ChatappSyncPhoneNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappSyncPhoneNumberRequest
	GetResourceOwnerId() *int64
}

type ChatappSyncPhoneNumberRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChatappSyncPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappSyncPhoneNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappSyncPhoneNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappSyncPhoneNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappSyncPhoneNumberRequest) SetCustSpaceId(v string) *ChatappSyncPhoneNumberRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappSyncPhoneNumberRequest) SetOwnerId(v int64) *ChatappSyncPhoneNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappSyncPhoneNumberRequest) SetResourceOwnerAccount(v string) *ChatappSyncPhoneNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappSyncPhoneNumberRequest) SetResourceOwnerId(v int64) *ChatappSyncPhoneNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappSyncPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
