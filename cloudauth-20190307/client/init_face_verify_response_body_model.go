// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *InitFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *InitFaceVerifyResponseBodyResultObject) *InitFaceVerifyResponseBody
	GetResultObject() *InitFaceVerifyResponseBodyResultObject
}

type InitFaceVerifyResponseBody struct {
	// The response code. 200 indicates success. Other values indicate failure.
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
	// The result object.
	ResultObject *InitFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitFaceVerifyResponseBody) GetResultObject() *InitFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *InitFaceVerifyResponseBody) SetCode(v string) *InitFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetMessage(v string) *InitFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetRequestId(v string) *InitFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitFaceVerifyResponseBody) SetResultObject(v *InitFaceVerifyResponseBodyResultObject) *InitFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *InitFaceVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitFaceVerifyResponseBodyResultObject struct {
	// The unique identifier for ID Verification.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The URL for performing ID Verification in a web browser. After authentication is complete, the page redirects based on the ReturnUrl parameter.
	//
	// 	Notice:
	//
	// - The CertifyUrl returned by the initialization operation is **valid for 30 minutes and can be submitted for authentication only once**. Use it within the validity period and do not reuse it.
	//
	// - This parameter requires the **MetaInfo*	- parameter to be correctly passed in to return a CertifyUrl that matches the client. If the URL cannot be obtained, check whether **MetaInfo*	- and other input parameters are correct.
	//
	// - The domain name of this URL may change with service updates. To ensure normal service availability, do not apply access control to this domain name.
	//
	// - Do not use incognito mode or modify the URL during browser redirection. Otherwise, a **signature exception*	- error may occur.
	//
	// .
	//
	// example:
	//
	// https://t.aliyun.com/****
	CertifyUrl *string `json:"CertifyUrl,omitempty" xml:"CertifyUrl,omitempty"`
}

func (s InitFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s InitFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitFaceVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InitFaceVerifyResponseBodyResultObject) GetCertifyUrl() *string {
	return s.CertifyUrl
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitFaceVerifyResponseBodyResultObject) SetCertifyUrl(v string) *InitFaceVerifyResponseBodyResultObject {
	s.CertifyUrl = &v
	return s
}

func (s *InitFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
