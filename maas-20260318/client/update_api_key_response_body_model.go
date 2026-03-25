// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApiKeyResponseBody
	GetSuccess() *bool
}

type UpdateApiKeyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 53EAEBC0-4DEC-5AF4-AA21-3923D5A819C3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApiKeyResponseBody) SetCode(v string) *UpdateApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetHttpStatusCode(v int32) *UpdateApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetMessage(v string) *UpdateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetRequestId(v string) *UpdateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetSuccess(v bool) *UpdateApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
