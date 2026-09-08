// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllCommonContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadAllCommonContactsResponseBody
	GetCode() *string
	SetData(v []*ReadAllCommonContactsResponseBodyData) *ReadAllCommonContactsResponseBody
	GetData() []*ReadAllCommonContactsResponseBodyData
	SetMessage(v string) *ReadAllCommonContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadAllCommonContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadAllCommonContactsResponseBody
	GetSuccess() *bool
}

type ReadAllCommonContactsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data []*ReadAllCommonContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The result message.
	//
	// example:
	//
	// /
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73FD6AE8-898F-5D09-9763-69B8A875488A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadAllCommonContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadAllCommonContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ReadAllCommonContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadAllCommonContactsResponseBody) GetData() []*ReadAllCommonContactsResponseBodyData {
	return s.Data
}

func (s *ReadAllCommonContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadAllCommonContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadAllCommonContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadAllCommonContactsResponseBody) SetCode(v string) *ReadAllCommonContactsResponseBody {
	s.Code = &v
	return s
}

func (s *ReadAllCommonContactsResponseBody) SetData(v []*ReadAllCommonContactsResponseBodyData) *ReadAllCommonContactsResponseBody {
	s.Data = v
	return s
}

func (s *ReadAllCommonContactsResponseBody) SetMessage(v string) *ReadAllCommonContactsResponseBody {
	s.Message = &v
	return s
}

func (s *ReadAllCommonContactsResponseBody) SetRequestId(v string) *ReadAllCommonContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadAllCommonContactsResponseBody) SetSuccess(v bool) *ReadAllCommonContactsResponseBody {
	s.Success = &v
	return s
}

func (s *ReadAllCommonContactsResponseBody) Validate() error {
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

type ReadAllCommonContactsResponseBodyData struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// /
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The email address of the contact.
	//
	// example:
	//
	// t*@qq.*
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// The contact ID in the Account Center. A value of 0 indicates the account contact.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The mobile phone number of the contact in the Account Center (masked).
	//
	// example:
	//
	// 130**123
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The contact name in the Account Center.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Indicates whether the email address is verified.
	//
	// example:
	//
	// true
	EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	// Indicates whether the mobile phone number of the contact in the Account Center is verified.
	//
	// example:
	//
	// true
	MobileConfirmed *bool `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// The position of the contact in the Account Center.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ReadAllCommonContactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadAllCommonContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadAllCommonContactsResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReadAllCommonContactsResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ReadAllCommonContactsResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadAllCommonContactsResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *ReadAllCommonContactsResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadAllCommonContactsResponseBodyData) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *ReadAllCommonContactsResponseBodyData) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *ReadAllCommonContactsResponseBodyData) GetPosition() *string {
	return s.Position
}

func (s *ReadAllCommonContactsResponseBodyData) SetAliUid(v int64) *ReadAllCommonContactsResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetContactEmail(v string) *ReadAllCommonContactsResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetContactId(v int64) *ReadAllCommonContactsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetContactMobile(v string) *ReadAllCommonContactsResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetContactName(v string) *ReadAllCommonContactsResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetEmailConfirmed(v bool) *ReadAllCommonContactsResponseBodyData {
	s.EmailConfirmed = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetMobileConfirmed(v bool) *ReadAllCommonContactsResponseBodyData {
	s.MobileConfirmed = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) SetPosition(v string) *ReadAllCommonContactsResponseBodyData {
	s.Position = &v
	return s
}

func (s *ReadAllCommonContactsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
