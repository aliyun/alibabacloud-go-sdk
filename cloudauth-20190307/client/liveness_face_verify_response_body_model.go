// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLivenessFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LivenessFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *LivenessFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *LivenessFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *LivenessFaceVerifyResponseBodyResultObject) *LivenessFaceVerifyResponseBody
	GetResultObject() *LivenessFaceVerifyResponseBodyResultObject
}

type LivenessFaceVerifyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *LivenessFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s LivenessFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LivenessFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *LivenessFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LivenessFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LivenessFaceVerifyResponseBody) GetResultObject() *LivenessFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *LivenessFaceVerifyResponseBody) SetCode(v string) *LivenessFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetMessage(v string) *LivenessFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetRequestId(v string) *LivenessFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBody) SetResultObject(v *LivenessFaceVerifyResponseBodyResultObject) *LivenessFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *LivenessFaceVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type LivenessFaceVerifyResponseBodyResultObject struct {
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// {"faceAttack": "F","facialPictureFront": {"qualityScore": 88.3615493774414,"verifyScore": 50.28594166529785}}
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s LivenessFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s LivenessFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *LivenessFaceVerifyResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *LivenessFaceVerifyResponseBodyResultObject) GetPassed() *string {
	return s.Passed
}

func (s *LivenessFaceVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetPassed(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) SetSubCode(v string) *LivenessFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *LivenessFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
