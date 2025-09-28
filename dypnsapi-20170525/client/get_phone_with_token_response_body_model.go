// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneWithTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhoneWithTokenResponseBody
	GetCode() *string
	SetData(v *GetPhoneWithTokenResponseBodyData) *GetPhoneWithTokenResponseBody
	GetData() *GetPhoneWithTokenResponseBodyData
	SetMessage(v string) *GetPhoneWithTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhoneWithTokenResponseBody
	GetRequestId() *string
}

type GetPhoneWithTokenResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetPhoneWithTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F335F48-****-****-****-CA7914FE5D77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneWithTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhoneWithTokenResponseBody) GetData() *GetPhoneWithTokenResponseBodyData {
	return s.Data
}

func (s *GetPhoneWithTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhoneWithTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhoneWithTokenResponseBody) SetCode(v string) *GetPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetData(v *GetPhoneWithTokenResponseBodyData) *GetPhoneWithTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetMessage(v string) *GetPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetRequestId(v string) *GetPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhoneWithTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPhoneWithTokenResponseBodyData struct {
	// The phone number.
	//
	// example:
	//
	// 13900001234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetPhoneWithTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneWithTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *GetPhoneWithTokenResponseBodyData) SetMobile(v string) *GetPhoneWithTokenResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetPhoneWithTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
