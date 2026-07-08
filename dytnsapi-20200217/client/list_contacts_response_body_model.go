// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContactsResponseBody
	GetCode() *string
	SetData(v []*ListContactsResponseBodyData) *ListContactsResponseBody
	GetData() []*ListContactsResponseBodyData
	SetMessage(v string) *ListContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContactsResponseBody
	GetRequestId() *string
}

type ListContactsResponseBody struct {
	// The status code. `OK` indicates a successful request.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array of contact information objects.
	Data []*ListContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContactsResponseBody) GetData() []*ListContactsResponseBodyData {
	return s.Data
}

func (s *ListContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactsResponseBody) SetCode(v string) *ListContactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactsResponseBody) SetData(v []*ListContactsResponseBodyData) *ListContactsResponseBody {
	s.Data = v
	return s
}

func (s *ListContactsResponseBody) SetMessage(v string) *ListContactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactsResponseBody) SetRequestId(v string) *ListContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContactsResponseBodyData struct {
	// The contact email.
	//
	// example:
	//
	// xxxx @alibaba-inc.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// The contact ID.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The contact name.
	//
	// example:
	//
	// 黄勇
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The contact mobile number.
	//
	// example:
	//
	// 19211111111
	ContactPhone *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	// The email status.
	//
	// - 1: Normal
	//
	// - 0: Abnormal
	//
	// example:
	//
	// 1
	MailStatus *int32 `json:"MailStatus,omitempty" xml:"MailStatus,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 0
	Main *int32 `json:"Main,omitempty" xml:"Main,omitempty"`
	// Specifies whether the number status warning is enabled.
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	OpenStatusWarning *bool `json:"OpenStatusWarning,omitempty" xml:"OpenStatusWarning,omitempty"`
	// Specifies whether the number attribution query warning is enabled.
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	OpentAttributionWarning *bool `json:"OpentAttributionWarning,omitempty" xml:"OpentAttributionWarning,omitempty"`
	// The number status.
	//
	// - 1: Normal
	//
	// - 0: Abnormal
	//
	// example:
	//
	// 1
	PhoneStatus *int32 `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty"`
}

func (s ListContactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ListContactsResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ListContactsResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ListContactsResponseBodyData) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *ListContactsResponseBodyData) GetMailStatus() *int32 {
	return s.MailStatus
}

func (s *ListContactsResponseBodyData) GetMain() *int32 {
	return s.Main
}

func (s *ListContactsResponseBodyData) GetOpenStatusWarning() *bool {
	return s.OpenStatusWarning
}

func (s *ListContactsResponseBodyData) GetOpentAttributionWarning() *bool {
	return s.OpentAttributionWarning
}

func (s *ListContactsResponseBodyData) GetPhoneStatus() *int32 {
	return s.PhoneStatus
}

func (s *ListContactsResponseBodyData) SetContactEmail(v string) *ListContactsResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *ListContactsResponseBodyData) SetContactId(v int64) *ListContactsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListContactsResponseBodyData) SetContactName(v string) *ListContactsResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ListContactsResponseBodyData) SetContactPhone(v string) *ListContactsResponseBodyData {
	s.ContactPhone = &v
	return s
}

func (s *ListContactsResponseBodyData) SetMailStatus(v int32) *ListContactsResponseBodyData {
	s.MailStatus = &v
	return s
}

func (s *ListContactsResponseBodyData) SetMain(v int32) *ListContactsResponseBodyData {
	s.Main = &v
	return s
}

func (s *ListContactsResponseBodyData) SetOpenStatusWarning(v bool) *ListContactsResponseBodyData {
	s.OpenStatusWarning = &v
	return s
}

func (s *ListContactsResponseBodyData) SetOpentAttributionWarning(v bool) *ListContactsResponseBodyData {
	s.OpentAttributionWarning = &v
	return s
}

func (s *ListContactsResponseBodyData) SetPhoneStatus(v int32) *ListContactsResponseBodyData {
	s.PhoneStatus = &v
	return s
}

func (s *ListContactsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
