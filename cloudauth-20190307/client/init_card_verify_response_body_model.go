// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitCardVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitCardVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *InitCardVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitCardVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *InitCardVerifyResponseBodyResultObject) *InitCardVerifyResponseBody
	GetResultObject() *InitCardVerifyResponseBodyResultObject
}

type InitCardVerifyResponseBody struct {
	// The response code. A value of 200 indicates success. Other values indicate failure.
	//
	// > **Important*	- This parameter indicates whether the operation is called correctly. For more information about return codes, see error codes. Check the fields in ResultObject for the business result.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
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
	// The returned result.
	ResultObject *InitCardVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s InitCardVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitCardVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *InitCardVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitCardVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitCardVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitCardVerifyResponseBody) GetResultObject() *InitCardVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *InitCardVerifyResponseBody) SetCode(v string) *InitCardVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *InitCardVerifyResponseBody) SetMessage(v string) *InitCardVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *InitCardVerifyResponseBody) SetRequestId(v string) *InitCardVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitCardVerifyResponseBody) SetResultObject(v *InitCardVerifyResponseBodyResultObject) *InitCardVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *InitCardVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitCardVerifyResponseBodyResultObject struct {
	// The verification request ID, which is the unique identifier of the verification service authentication request.
	//
	// - You must specify the authentication request ID when you query the authentication result.
	//
	// - The CertifyId field is used for billing statistics. Save this field locally for future bill reconciliation. The CertifyId returned by the initialization operation is valid for 30 minutes and can be submitted for authentication only once. Use it within the validity period and do not reuse it.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0****
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
}

func (s InitCardVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s InitCardVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *InitCardVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *InitCardVerifyResponseBodyResultObject) SetCertifyId(v string) *InitCardVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *InitCardVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
