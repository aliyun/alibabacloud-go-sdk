// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPhoneWithTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyPhoneWithTokenResponseBody
	GetCode() *string
	SetGateVerify(v *VerifyPhoneWithTokenResponseBodyGateVerify) *VerifyPhoneWithTokenResponseBody
	GetGateVerify() *VerifyPhoneWithTokenResponseBodyGateVerify
	SetMessage(v string) *VerifyPhoneWithTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyPhoneWithTokenResponseBody
	GetRequestId() *string
}

type VerifyPhoneWithTokenResponseBody struct {
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
	GateVerify *VerifyPhoneWithTokenResponseBodyGateVerify `json:"GateVerify,omitempty" xml:"GateVerify,omitempty" type:"Struct"`
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
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyPhoneWithTokenResponseBody) GetGateVerify() *VerifyPhoneWithTokenResponseBodyGateVerify {
	return s.GateVerify
}

func (s *VerifyPhoneWithTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyPhoneWithTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyPhoneWithTokenResponseBody) SetCode(v string) *VerifyPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetGateVerify(v *VerifyPhoneWithTokenResponseBodyGateVerify) *VerifyPhoneWithTokenResponseBody {
	s.GateVerify = v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetMessage(v string) *VerifyPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetRequestId(v string) *VerifyPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type VerifyPhoneWithTokenResponseBodyGateVerify struct {
	// The external ID.
	//
	// example:
	//
	// 12134****
	VerifyId *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- PASS: The input phone number is consistent with the phone number used in HTML5 pages.
	//
	// 	- REJECT: The input phone number is different from the phone number used in HTML5 pages.
	//
	// 	- UNKNOWN: The system cannot judge whether the input phone number is consistent with the phone number used in HTML5 pages.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) String() string {
	return dara.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) GetVerifyId() *string {
	return s.VerifyId
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyId(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyId = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyResult(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyResult = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) Validate() error {
	return dara.Validate(s)
}
