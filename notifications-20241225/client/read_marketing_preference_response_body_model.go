// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMarketingPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMarketingPreferenceResponseBody
	GetCode() *string
	SetData(v *ReadMarketingPreferenceResponseBodyData) *ReadMarketingPreferenceResponseBody
	GetData() *ReadMarketingPreferenceResponseBodyData
	SetMessage(v string) *ReadMarketingPreferenceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMarketingPreferenceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMarketingPreferenceResponseBody
	GetSuccess() *bool
}

type ReadMarketingPreferenceResponseBody struct {
	// The error code returned by the system. For more information about error codes, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data *ReadMarketingPreferenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// Succeeded
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMarketingPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMarketingPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMarketingPreferenceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMarketingPreferenceResponseBody) GetData() *ReadMarketingPreferenceResponseBodyData {
	return s.Data
}

func (s *ReadMarketingPreferenceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMarketingPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMarketingPreferenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMarketingPreferenceResponseBody) SetCode(v string) *ReadMarketingPreferenceResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBody) SetData(v *ReadMarketingPreferenceResponseBodyData) *ReadMarketingPreferenceResponseBody {
	s.Data = v
	return s
}

func (s *ReadMarketingPreferenceResponseBody) SetMessage(v string) *ReadMarketingPreferenceResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBody) SetRequestId(v string) *ReadMarketingPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBody) SetSuccess(v bool) *ReadMarketingPreferenceResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadMarketingPreferenceResponseBodyData struct {
	// Indicates whether notifications are allowed.
	//
	// example:
	//
	// true
	AllowMarketing *bool `json:"AllowMarketing,omitempty" xml:"AllowMarketing,omitempty"`
	// The email address of the contact in Account Center.
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
	// 130*123
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The name of the contact in Account Center.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The position of the contact in Account Center.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ReadMarketingPreferenceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadMarketingPreferenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadMarketingPreferenceResponseBodyData) GetAllowMarketing() *bool {
	return s.AllowMarketing
}

func (s *ReadMarketingPreferenceResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ReadMarketingPreferenceResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadMarketingPreferenceResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *ReadMarketingPreferenceResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadMarketingPreferenceResponseBodyData) GetPosition() *string {
	return s.Position
}

func (s *ReadMarketingPreferenceResponseBodyData) SetAllowMarketing(v bool) *ReadMarketingPreferenceResponseBodyData {
	s.AllowMarketing = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) SetContactEmail(v string) *ReadMarketingPreferenceResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) SetContactId(v int64) *ReadMarketingPreferenceResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) SetContactMobile(v string) *ReadMarketingPreferenceResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) SetContactName(v string) *ReadMarketingPreferenceResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) SetPosition(v string) *ReadMarketingPreferenceResponseBodyData {
	s.Position = &v
	return s
}

func (s *ReadMarketingPreferenceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
