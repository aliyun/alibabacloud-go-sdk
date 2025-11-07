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
	// Return code: 200 for success, others for failure.
	//
	// Important
	//
	// - This parameter indicates whether the interface was called correctly. For detailed return code descriptions, see the error codes.
	//
	// - Business results should be viewed through the fields in ResultObject.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
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
	// Return result.
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
	// Verification request ID, a unique identifier for the verification service\\"s authentication request.
	//
	// - When querying the authentication result, the authentication request ID must be provided.
	//
	// - The CertifyId field is a billing statistics field. To facilitate subsequent bill reconciliation, please retain this field information locally. The CertifyId returned by the initialization interface is valid for 30 minutes and can only be submitted once for authentication. Please apply it within the validity period to avoid reuse.
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
