// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhoneNumberIdentificationResultResponseBody
	GetCode() *string
	SetData(v *GetPhoneNumberIdentificationResultResponseBodyData) *GetPhoneNumberIdentificationResultResponseBody
	GetData() *GetPhoneNumberIdentificationResultResponseBodyData
	SetMessage(v string) *GetPhoneNumberIdentificationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhoneNumberIdentificationResultResponseBody
	GetRequestId() *string
}

type GetPhoneNumberIdentificationResultResponseBody struct {
	// The return code. Valid values:
	//
	// 	- OK: The request is successful.
	//
	// 	- NoIdentificationResult: No verification result is available or the verification failed.
	//
	// 	- SessionNotValid: The session is invalid or expired.
	//
	// 	- MobileNumberIllegal: The format of the phone number is invalid.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPhoneNumberIdentificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the return code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneNumberIdentificationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhoneNumberIdentificationResultResponseBody) GetData() *GetPhoneNumberIdentificationResultResponseBodyData {
	return s.Data
}

func (s *GetPhoneNumberIdentificationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhoneNumberIdentificationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetCode(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetData(v *GetPhoneNumberIdentificationResultResponseBodyData) *GetPhoneNumberIdentificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetMessage(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetRequestId(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhoneNumberIdentificationResultResponseBodyData struct {
	// Indicates whether the phone number passed the verification.
	//
	// example:
	//
	// true
	IsIdentified *string `json:"IsIdentified,omitempty" xml:"IsIdentified,omitempty"`
}

func (s GetPhoneNumberIdentificationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponseBodyData) GetIsIdentified() *string {
	return s.IsIdentified
}

func (s *GetPhoneNumberIdentificationResultResponseBodyData) SetIsIdentified(v string) *GetPhoneNumberIdentificationResultResponseBodyData {
	s.IsIdentified = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
