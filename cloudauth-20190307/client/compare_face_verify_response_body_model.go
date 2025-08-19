// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CompareFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *CompareFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CompareFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *CompareFaceVerifyResponseBodyResultObject) *CompareFaceVerifyResponseBody
	GetResultObject() *CompareFaceVerifyResponseBodyResultObject
}

type CompareFaceVerifyResponseBody struct {
	// Return code: 200 for success, other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Face comparison result information.
	ResultObject *CompareFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CompareFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompareFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompareFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareFaceVerifyResponseBody) GetResultObject() *CompareFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *CompareFaceVerifyResponseBody) SetCode(v string) *CompareFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetMessage(v string) *CompareFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetRequestId(v string) *CompareFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceVerifyResponseBody) SetResultObject(v *CompareFaceVerifyResponseBodyResultObject) *CompareFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *CompareFaceVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CompareFaceVerifyResponseBodyResultObject struct {
	// Unique identifier for the real-person authentication request.
	//
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Whether the verification passed, T for pass, F for fail.
	//
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Face comparison score.
	//
	// example:
	//
	// 99.60875
	VerifyScore *float32 `json:"VerifyScore,omitempty" xml:"VerifyScore,omitempty"`
}

func (s CompareFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *CompareFaceVerifyResponseBodyResultObject) GetPassed() *string {
	return s.Passed
}

func (s *CompareFaceVerifyResponseBodyResultObject) GetVerifyScore() *float32 {
	return s.VerifyScore
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetPassed(v string) *CompareFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) SetVerifyScore(v float32) *CompareFaceVerifyResponseBodyResultObject {
	s.VerifyScore = &v
	return s
}

func (s *CompareFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
