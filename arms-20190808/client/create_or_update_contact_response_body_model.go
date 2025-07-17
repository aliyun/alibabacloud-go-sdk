// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertContact(v *CreateOrUpdateContactResponseBodyAlertContact) *CreateOrUpdateContactResponseBody
	GetAlertContact() *CreateOrUpdateContactResponseBodyAlertContact
	SetRequestId(v string) *CreateOrUpdateContactResponseBody
	GetRequestId() *string
}

type CreateOrUpdateContactResponseBody struct {
	// The object of the alert contact.
	AlertContact *CreateOrUpdateContactResponseBodyAlertContact `json:"AlertContact,omitempty" xml:"AlertContact,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E9C9DA3D-10FE-472E-9EEF-2D0A3E41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponseBody) GetAlertContact() *CreateOrUpdateContactResponseBodyAlertContact {
	return s.AlertContact
}

func (s *CreateOrUpdateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateContactResponseBody) SetAlertContact(v *CreateOrUpdateContactResponseBodyAlertContact) *CreateOrUpdateContactResponseBody {
	s.AlertContact = v
	return s
}

func (s *CreateOrUpdateContactResponseBody) SetRequestId(v string) *CreateOrUpdateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateContactResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateContactResponseBodyAlertContact struct {
	// The ID of the alert contact.
	//
	// example:
	//
	// 123
	ContactId *float32 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// JohnDoe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://test1.com
	DingRobotUrl *string `json:"DingRobotUrl,omitempty" xml:"DingRobotUrl,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the mobile number was verified. Valid values:
	//
	// 	- `false` (default value): No
	//
	// 	- `true`: Yes
	//
	// You can call the **SendTTSVerifyLink*	- operation to verify the mobile number of an alert contact. Only verified mobile numbers can be specified in a notification policy to receive phone calls.
	//
	// example:
	//
	// false
	IsVerify *bool `json:"IsVerify,omitempty" xml:"IsVerify,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The operation that you want to perform if phone calls fail to be answered. Valid values: 0: No operation is performed. 1: A phone call is made again. 2: A text message is sent. 3 (default value): The global default value is used.
	//
	// example:
	//
	// 3
	ReissueSendNotice *int64 `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
	// Indicates whether the email address was verified.
	//
	// example:
	//
	// true
	IsEmailVerify *bool `json:"isEmailVerify,omitempty" xml:"isEmailVerify,omitempty"`
}

func (s CreateOrUpdateContactResponseBodyAlertContact) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactResponseBodyAlertContact) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetContactId() *float32 {
	return s.ContactId
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetContactName() *string {
	return s.ContactName
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetDingRobotUrl() *string {
	return s.DingRobotUrl
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetEmail() *string {
	return s.Email
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetIsVerify() *bool {
	return s.IsVerify
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetPhone() *string {
	return s.Phone
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetReissueSendNotice() *int64 {
	return s.ReissueSendNotice
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) GetIsEmailVerify() *bool {
	return s.IsEmailVerify
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetContactId(v float32) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ContactId = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetContactName(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ContactName = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetDingRobotUrl(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.DingRobotUrl = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetEmail(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.Email = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetIsVerify(v bool) *CreateOrUpdateContactResponseBodyAlertContact {
	s.IsVerify = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetPhone(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.Phone = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetReissueSendNotice(v int64) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ReissueSendNotice = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetIsEmailVerify(v bool) *CreateOrUpdateContactResponseBodyAlertContact {
	s.IsEmailVerify = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) Validate() error {
	return dara.Validate(s)
}
