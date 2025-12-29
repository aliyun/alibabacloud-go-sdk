// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactEmail(v string) *UpdateContactsRequest
	GetContactEmail() *string
	SetContactId(v int64) *UpdateContactsRequest
	GetContactId() *int64
	SetContactName(v string) *UpdateContactsRequest
	GetContactName() *string
	SetContactPhone(v string) *UpdateContactsRequest
	GetContactPhone() *string
	SetMailStatus(v int32) *UpdateContactsRequest
	GetMailStatus() *int32
	SetOpenStatusWarning(v bool) *UpdateContactsRequest
	GetOpenStatusWarning() *bool
	SetOpentAttributionWarning(v bool) *UpdateContactsRequest
	GetOpentAttributionWarning() *bool
	SetOwnerId(v int64) *UpdateContactsRequest
	GetOwnerId() *int64
	SetPhoneStatus(v int32) *UpdateContactsRequest
	GetPhoneStatus() *int32
	SetResourceOwnerAccount(v string) *UpdateContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateContactsRequest
	GetResourceOwnerId() *int64
}

type UpdateContactsRequest struct {
	// example:
	//
	// XXXX@alibaba-inc.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1194432
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// XXX
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 192XXXXXXXX
	ContactPhone *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	// example:
	//
	// 1
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
	// 1
	PhoneStatus          *int32  `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactsRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactsRequest) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *UpdateContactsRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateContactsRequest) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *UpdateContactsRequest) GetMailStatus() *int32 {
	return s.MailStatus
}

func (s *UpdateContactsRequest) GetOpenStatusWarning() *bool {
	return s.OpenStatusWarning
}

func (s *UpdateContactsRequest) GetOpentAttributionWarning() *bool {
	return s.OpentAttributionWarning
}

func (s *UpdateContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateContactsRequest) GetPhoneStatus() *int32 {
	return s.PhoneStatus
}

func (s *UpdateContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateContactsRequest) SetContactEmail(v string) *UpdateContactsRequest {
	s.ContactEmail = &v
	return s
}

func (s *UpdateContactsRequest) SetContactId(v int64) *UpdateContactsRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateContactsRequest) SetContactName(v string) *UpdateContactsRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateContactsRequest) SetContactPhone(v string) *UpdateContactsRequest {
	s.ContactPhone = &v
	return s
}

func (s *UpdateContactsRequest) SetMailStatus(v int32) *UpdateContactsRequest {
	s.MailStatus = &v
	return s
}

func (s *UpdateContactsRequest) SetOpenStatusWarning(v bool) *UpdateContactsRequest {
	s.OpenStatusWarning = &v
	return s
}

func (s *UpdateContactsRequest) SetOpentAttributionWarning(v bool) *UpdateContactsRequest {
	s.OpentAttributionWarning = &v
	return s
}

func (s *UpdateContactsRequest) SetOwnerId(v int64) *UpdateContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContactsRequest) SetPhoneStatus(v int32) *UpdateContactsRequest {
	s.PhoneStatus = &v
	return s
}

func (s *UpdateContactsRequest) SetResourceOwnerAccount(v string) *UpdateContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContactsRequest) SetResourceOwnerId(v int64) *UpdateContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContactsRequest) Validate() error {
	return dara.Validate(s)
}
