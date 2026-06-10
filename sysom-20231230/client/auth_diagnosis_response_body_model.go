// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthDiagnosisResponseBody
	GetCode() *string
	SetData(v interface{}) *AuthDiagnosisResponseBody
	GetData() interface{}
	SetMessage(v string) *AuthDiagnosisResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthDiagnosisResponseBody
	GetRequestId() *string
}

type AuthDiagnosisResponseBody struct {
	// Status code
	//
	// - If `code == Success`, the authorization succeeded.
	//
	// - Any other status code indicates that the authorization failed. In this case, check the `message` field for detailed error information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This API returns no data.
	//
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request RequestId
	//
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s AuthDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthDiagnosisResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AuthDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthDiagnosisResponseBody) SetCode(v string) *AuthDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetData(v interface{}) *AuthDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *AuthDiagnosisResponseBody) SetMessage(v string) *AuthDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetRequestId(v string) *AuthDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
