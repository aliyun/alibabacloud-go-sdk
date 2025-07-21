// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommerceSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetCommerceSettingRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetCommerceSettingRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetCommerceSettingRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetCommerceSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCommerceSettingRequest
	GetResourceOwnerId() *int64
}

type GetCommerceSettingRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCommerceSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommerceSettingRequest) GoString() string {
	return s.String()
}

func (s *GetCommerceSettingRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetCommerceSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCommerceSettingRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetCommerceSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCommerceSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCommerceSettingRequest) SetCustSpaceId(v string) *GetCommerceSettingRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetCommerceSettingRequest) SetOwnerId(v int64) *GetCommerceSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCommerceSettingRequest) SetPhoneNumber(v string) *GetCommerceSettingRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetCommerceSettingRequest) SetResourceOwnerAccount(v string) *GetCommerceSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCommerceSettingRequest) SetResourceOwnerId(v int64) *GetCommerceSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCommerceSettingRequest) Validate() error {
	return dara.Validate(s)
}
