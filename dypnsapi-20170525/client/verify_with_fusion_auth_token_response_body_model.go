// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWithFusionAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyWithFusionAuthTokenResponseBody
	GetCode() *string
	SetMessage(v string) *VerifyWithFusionAuthTokenResponseBody
	GetMessage() *string
	SetModel(v *VerifyWithFusionAuthTokenResponseBodyModel) *VerifyWithFusionAuthTokenResponseBody
	GetModel() *VerifyWithFusionAuthTokenResponseBodyModel
	SetRequestId(v string) *VerifyWithFusionAuthTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyWithFusionAuthTokenResponseBody
	GetSuccess() *bool
}

type VerifyWithFusionAuthTokenResponseBody struct {
	// The response code. If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *VerifyWithFusionAuthTokenResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID, which is used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyWithFusionAuthTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyWithFusionAuthTokenResponseBody) GetModel() *VerifyWithFusionAuthTokenResponseBodyModel {
	return s.Model
}

func (s *VerifyWithFusionAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyWithFusionAuthTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetCode(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetMessage(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetModel(v *VerifyWithFusionAuthTokenResponseBodyModel) *VerifyWithFusionAuthTokenResponseBody {
	s.Model = v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetRequestId(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetSuccess(v bool) *VerifyWithFusionAuthTokenResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type VerifyWithFusionAuthTokenResponseBodyModel struct {
	// The phone number, which is returned when the verification is successful.
	//
	// example:
	//
	// 180********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The phone number score, which is generated only after the phone number scoring node is enabled and the verification is successful. The higher the score, the more risky the phone number. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	PhoneScore *int64 `json:"PhoneScore,omitempty" xml:"PhoneScore,omitempty"`
	// The verification result. Valid values: PASS and UNKNOWN.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponseBodyModel) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) GetPhoneScore() *int64 {
	return s.PhoneScore
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetPhoneNumber(v string) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetPhoneScore(v int64) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.PhoneScore = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetVerifyResult(v string) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.VerifyResult = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
