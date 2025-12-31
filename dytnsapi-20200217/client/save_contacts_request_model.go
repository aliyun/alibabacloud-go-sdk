// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *SaveContactsRequest
	GetBizType() *string
	SetContactEmail(v string) *SaveContactsRequest
	GetContactEmail() *string
	SetContactName(v string) *SaveContactsRequest
	GetContactName() *string
	SetContactPhone(v string) *SaveContactsRequest
	GetContactPhone() *string
	SetMailStatus(v int32) *SaveContactsRequest
	GetMailStatus() *int32
	SetOpenStatusWarning(v bool) *SaveContactsRequest
	GetOpenStatusWarning() *bool
	SetOpentAttributionWarning(v bool) *SaveContactsRequest
	GetOpentAttributionWarning() *bool
	SetOwnerId(v int64) *SaveContactsRequest
	GetOwnerId() *int64
	SetPhoneStatus(v int32) *SaveContactsRequest
	GetPhoneStatus() *int32
	SetResourceOwnerAccount(v string) *SaveContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SaveContactsRequest
	GetResourceOwnerId() *int64
}

type SaveContactsRequest struct {
	// example:
	//
	// dytns
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1234@alibaba-inc.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// This parameter is required.
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 122354532434
	ContactPhone *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	// example:
	//
	// 0
	MailStatus *int32 `json:"MailStatus,omitempty" xml:"MailStatus,omitempty"`
	// example:
	//
	// false
	OpenStatusWarning *bool `json:"OpenStatusWarning,omitempty" xml:"OpenStatusWarning,omitempty"`
	// example:
	//
	// true
	OpentAttributionWarning *bool  `json:"OpentAttributionWarning,omitempty" xml:"OpentAttributionWarning,omitempty"`
	OwnerId                 *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 0
	PhoneStatus          *int32  `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SaveContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContactsRequest) GoString() string {
	return s.String()
}

func (s *SaveContactsRequest) GetBizType() *string {
	return s.BizType
}

func (s *SaveContactsRequest) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *SaveContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *SaveContactsRequest) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *SaveContactsRequest) GetMailStatus() *int32 {
	return s.MailStatus
}

func (s *SaveContactsRequest) GetOpenStatusWarning() *bool {
	return s.OpenStatusWarning
}

func (s *SaveContactsRequest) GetOpentAttributionWarning() *bool {
	return s.OpentAttributionWarning
}

func (s *SaveContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SaveContactsRequest) GetPhoneStatus() *int32 {
	return s.PhoneStatus
}

func (s *SaveContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SaveContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SaveContactsRequest) SetBizType(v string) *SaveContactsRequest {
	s.BizType = &v
	return s
}

func (s *SaveContactsRequest) SetContactEmail(v string) *SaveContactsRequest {
	s.ContactEmail = &v
	return s
}

func (s *SaveContactsRequest) SetContactName(v string) *SaveContactsRequest {
	s.ContactName = &v
	return s
}

func (s *SaveContactsRequest) SetContactPhone(v string) *SaveContactsRequest {
	s.ContactPhone = &v
	return s
}

func (s *SaveContactsRequest) SetMailStatus(v int32) *SaveContactsRequest {
	s.MailStatus = &v
	return s
}

func (s *SaveContactsRequest) SetOpenStatusWarning(v bool) *SaveContactsRequest {
	s.OpenStatusWarning = &v
	return s
}

func (s *SaveContactsRequest) SetOpentAttributionWarning(v bool) *SaveContactsRequest {
	s.OpentAttributionWarning = &v
	return s
}

func (s *SaveContactsRequest) SetOwnerId(v int64) *SaveContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *SaveContactsRequest) SetPhoneStatus(v int32) *SaveContactsRequest {
	s.PhoneStatus = &v
	return s
}

func (s *SaveContactsRequest) SetResourceOwnerAccount(v string) *SaveContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SaveContactsRequest) SetResourceOwnerId(v int64) *SaveContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SaveContactsRequest) Validate() error {
	return dara.Validate(s)
}
