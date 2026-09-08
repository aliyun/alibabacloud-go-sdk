// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMarketingPreferencesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadAllMarketingPreferencesResponseBody
	GetCode() *string
	SetData(v []*ReadAllMarketingPreferencesResponseBodyData) *ReadAllMarketingPreferencesResponseBody
	GetData() []*ReadAllMarketingPreferencesResponseBodyData
	SetMessage(v string) *ReadAllMarketingPreferencesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadAllMarketingPreferencesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadAllMarketingPreferencesResponseBody
	GetSuccess() *bool
}

type ReadAllMarketingPreferencesResponseBody struct {
	// The error code returned by the system. For more information about error codes, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data []*ReadAllMarketingPreferencesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message.
	//
	// example:
	//
	// /
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
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

func (s ReadAllMarketingPreferencesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMarketingPreferencesResponseBody) GoString() string {
	return s.String()
}

func (s *ReadAllMarketingPreferencesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadAllMarketingPreferencesResponseBody) GetData() []*ReadAllMarketingPreferencesResponseBodyData {
	return s.Data
}

func (s *ReadAllMarketingPreferencesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadAllMarketingPreferencesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadAllMarketingPreferencesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadAllMarketingPreferencesResponseBody) SetCode(v string) *ReadAllMarketingPreferencesResponseBody {
	s.Code = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBody) SetData(v []*ReadAllMarketingPreferencesResponseBodyData) *ReadAllMarketingPreferencesResponseBody {
	s.Data = v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBody) SetMessage(v string) *ReadAllMarketingPreferencesResponseBody {
	s.Message = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBody) SetRequestId(v string) *ReadAllMarketingPreferencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBody) SetSuccess(v bool) *ReadAllMarketingPreferencesResponseBody {
	s.Success = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBody) Validate() error {
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

type ReadAllMarketingPreferencesResponseBodyData struct {
	// Indicates whether notifications are allowed.
	//
	// example:
	//
	// true
	AllowMarketing *bool `json:"AllowMarketing,omitempty" xml:"AllowMarketing,omitempty"`
	// The email address of the contact in the Account Center (masked).
	//
	// example:
	//
	// test@aliyun.com
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
	// 130*123
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// The contact name in the Account Center.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The position of the contact in the Account Center.
	//
	// example:
	//
	// CEO
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ReadAllMarketingPreferencesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMarketingPreferencesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetAllowMarketing() *bool {
	return s.AllowMarketing
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadAllMarketingPreferencesResponseBodyData) GetPosition() *string {
	return s.Position
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetAllowMarketing(v bool) *ReadAllMarketingPreferencesResponseBodyData {
	s.AllowMarketing = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetContactEmail(v string) *ReadAllMarketingPreferencesResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetContactId(v int64) *ReadAllMarketingPreferencesResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetContactMobile(v string) *ReadAllMarketingPreferencesResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetContactName(v string) *ReadAllMarketingPreferencesResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) SetPosition(v string) *ReadAllMarketingPreferencesResponseBodyData {
	s.Position = &v
	return s
}

func (s *ReadAllMarketingPreferencesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
