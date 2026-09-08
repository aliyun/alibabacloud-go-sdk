// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCommonContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadCommonContactResponseBody
	GetCode() *string
	SetData(v *ReadCommonContactResponseBodyData) *ReadCommonContactResponseBody
	GetData() *ReadCommonContactResponseBodyData
	SetMessage(v string) *ReadCommonContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadCommonContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadCommonContactResponseBody
	GetSuccess() *bool
}

type ReadCommonContactResponseBody struct {
	// The error code returned if the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *ReadCommonContactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
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

func (s ReadCommonContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadCommonContactResponseBody) GoString() string {
	return s.String()
}

func (s *ReadCommonContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadCommonContactResponseBody) GetData() *ReadCommonContactResponseBodyData {
	return s.Data
}

func (s *ReadCommonContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadCommonContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadCommonContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadCommonContactResponseBody) SetCode(v string) *ReadCommonContactResponseBody {
	s.Code = &v
	return s
}

func (s *ReadCommonContactResponseBody) SetData(v *ReadCommonContactResponseBodyData) *ReadCommonContactResponseBody {
	s.Data = v
	return s
}

func (s *ReadCommonContactResponseBody) SetMessage(v string) *ReadCommonContactResponseBody {
	s.Message = &v
	return s
}

func (s *ReadCommonContactResponseBody) SetRequestId(v string) *ReadCommonContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadCommonContactResponseBody) SetSuccess(v bool) *ReadCommonContactResponseBody {
	s.Success = &v
	return s
}

func (s *ReadCommonContactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadCommonContactResponseBodyData struct {
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
	// The contact ID in Account Center. A value of 0 indicates the account contact.
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The mobile phone number of the contact in Account Center (masked).
	//
	// example:
	//
	// 130**123
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The contact name in Account Center.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Indicates whether the email address of the contact is verified.
	//
	// example:
	//
	// true
	EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	// Indicates whether the mobile phone number of the contact in Account Center is verified.
	//
	// example:
	//
	// true
	MobileConfirmed *bool `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// The position of the contact in Account Center.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ReadCommonContactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadCommonContactResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadCommonContactResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReadCommonContactResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ReadCommonContactResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadCommonContactResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *ReadCommonContactResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadCommonContactResponseBodyData) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *ReadCommonContactResponseBodyData) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *ReadCommonContactResponseBodyData) GetPosition() *string {
	return s.Position
}

func (s *ReadCommonContactResponseBodyData) SetAliUid(v int64) *ReadCommonContactResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetContactEmail(v string) *ReadCommonContactResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetContactId(v int64) *ReadCommonContactResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetContactMobile(v string) *ReadCommonContactResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetContactName(v string) *ReadCommonContactResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetEmailConfirmed(v bool) *ReadCommonContactResponseBodyData {
	s.EmailConfirmed = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetMobileConfirmed(v bool) *ReadCommonContactResponseBodyData {
	s.MobileConfirmed = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) SetPosition(v string) *ReadCommonContactResponseBodyData {
	s.Position = &v
	return s
}

func (s *ReadCommonContactResponseBodyData) Validate() error {
	return dara.Validate(s)
}
