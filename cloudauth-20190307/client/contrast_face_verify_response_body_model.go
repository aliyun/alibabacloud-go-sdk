// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContrastFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ContrastFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *ContrastFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ContrastFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *ContrastFaceVerifyResponseBodyResultObject) *ContrastFaceVerifyResponseBody
	GetResultObject() *ContrastFaceVerifyResponseBodyResultObject
}

type ContrastFaceVerifyResponseBody struct {
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *ContrastFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s ContrastFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContrastFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ContrastFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ContrastFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContrastFaceVerifyResponseBody) GetResultObject() *ContrastFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *ContrastFaceVerifyResponseBody) SetCode(v string) *ContrastFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetMessage(v string) *ContrastFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetRequestId(v string) *ContrastFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBody) SetResultObject(v *ContrastFaceVerifyResponseBodyResultObject) *ContrastFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *ContrastFaceVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ContrastFaceVerifyResponseBodyResultObject struct {
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// null
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
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

func (s ContrastFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ContrastFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *ContrastFaceVerifyResponseBodyResultObject) GetIdentityInfo() *string {
	return s.IdentityInfo
}

func (s *ContrastFaceVerifyResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *ContrastFaceVerifyResponseBodyResultObject) GetPassed() *string {
	return s.Passed
}

func (s *ContrastFaceVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetPassed(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) SetSubCode(v string) *ContrastFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *ContrastFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
