// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateUserResponseBody
	GetData() *bool
	SetErrorCode(v string) *CreateUserResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUserResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateUserResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateUserResponseBody
	GetSuccess() *string
}

type CreateUserResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUserResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateUserResponseBody) SetData(v bool) *CreateUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUserResponseBody) SetErrorCode(v string) *CreateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserResponseBody) SetErrorMessage(v string) *CreateUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUserResponseBody) SetHttpStatusCode(v string) *CreateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v string) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
