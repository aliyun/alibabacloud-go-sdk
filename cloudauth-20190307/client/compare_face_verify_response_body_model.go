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
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The face comparison result.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareFaceVerifyResponseBodyResultObject struct {
	// The unique identifier of the ID Verification request.
	//
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Indicates whether the verification passed. A value of T indicates passed. A value of F indicates not passed.
	//
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The face comparison score.
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
