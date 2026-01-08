// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberVerificationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetPhoneNumberVerificationStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetPhoneNumberVerificationStatusResponseBody
	GetCode() *string
	SetData(v *GetPhoneNumberVerificationStatusResponseBodyData) *GetPhoneNumberVerificationStatusResponseBody
	GetData() *GetPhoneNumberVerificationStatusResponseBodyData
	SetMessage(v string) *GetPhoneNumberVerificationStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhoneNumberVerificationStatusResponseBody
	GetRequestId() *string
}

type GetPhoneNumberVerificationStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetPhoneNumberVerificationStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetPhoneNumberVerificationStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhoneNumberVerificationStatusResponseBody) GetData() *GetPhoneNumberVerificationStatusResponseBodyData {
	return s.Data
}

func (s *GetPhoneNumberVerificationStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhoneNumberVerificationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetAccessDeniedDetail(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetCode(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetData(v *GetPhoneNumberVerificationStatusResponseBodyData) *GetPhoneNumberVerificationStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetMessage(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) SetRequestId(v string) *GetPhoneNumberVerificationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhoneNumberVerificationStatusResponseBodyData struct {
	// The verification status.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// The ID of the number.
	//
	// example:
	//
	// 2224342624
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) GetCodeVerificationStatus() *string {
	return s.CodeVerificationStatus
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetCodeVerificationStatus(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.CodeVerificationStatus = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetId(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) SetPhoneNumber(v string) *GetPhoneNumberVerificationStatusResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
