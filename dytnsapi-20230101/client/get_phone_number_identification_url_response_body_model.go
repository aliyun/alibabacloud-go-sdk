// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhoneNumberIdentificationUrlResponseBody
	GetCode() *string
	SetData(v *GetPhoneNumberIdentificationUrlResponseBodyData) *GetPhoneNumberIdentificationUrlResponseBody
	GetData() *GetPhoneNumberIdentificationUrlResponseBodyData
	SetMessage(v string) *GetPhoneNumberIdentificationUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhoneNumberIdentificationUrlResponseBody
	GetRequestId() *string
}

type GetPhoneNumberIdentificationUrlResponseBody struct {
	// The return code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **IdentificationNotAvailable**: The verification system does not support the phone number that corresponds to the IP address.
	//
	// 	- **MobileNumberIllegal**: The format of the phone number is invalid.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPhoneNumberIdentificationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetPhoneNumberIdentificationUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) GetData() *GetPhoneNumberIdentificationUrlResponseBodyData {
	return s.Data
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetCode(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetData(v *GetPhoneNumberIdentificationUrlResponseBodyData) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetMessage(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetRequestId(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhoneNumberIdentificationUrlResponseBodyData struct {
	// The verification URL.
	//
	// example:
	//
	// https://global-ip-auth.dycpaas.com/global/biz/ip_auth/start?ipa_s_c_c=IPF0000000000000******&ipa_s_i=8636b75e2fcb40c53ffecc2b59******
	IdentificationUrl *string `json:"IdentificationUrl,omitempty" xml:"IdentificationUrl,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 8636b75e2fcb40c53ffecc2b5947115c.149b03d2a7494e6e8f5b34c915245815.707c7f0d93f4409db0761aa5da94ce01.1686******041
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) GetIdentificationUrl() *string {
	return s.IdentificationUrl
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) SetIdentificationUrl(v string) *GetPhoneNumberIdentificationUrlResponseBodyData {
	s.IdentificationUrl = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) SetSessionId(v string) *GetPhoneNumberIdentificationUrlResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
