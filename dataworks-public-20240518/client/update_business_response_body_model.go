// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBusinessResponseBody
	GetSuccess() *bool
}

type UpdateBusinessResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBusinessResponseBody) SetErrorCode(v string) *UpdateBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateBusinessResponseBody) SetErrorMessage(v string) *UpdateBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateBusinessResponseBody) SetHttpStatusCode(v int32) *UpdateBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBusinessResponseBody) SetRequestId(v string) *UpdateBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBusinessResponseBody) SetSuccess(v bool) *UpdateBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}
