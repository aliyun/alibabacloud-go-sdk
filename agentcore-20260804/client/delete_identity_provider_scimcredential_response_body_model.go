// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderSCIMCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIdentityProviderSCIMCredentialResponseBody
	GetSuccess() *bool
}

type DeleteIdentityProviderSCIMCredentialResponseBody struct {
	// The result code. Handle failures based on the HTTP status code and message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code of the request. This does not indicate the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message returned for the request. For asynchronous tasks, query the task result by calling the corresponding API operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. You can use this ID to troubleshoot issues. Keep the ID if the request fails.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates that the request was successful. For asynchronous write operations, you must still query the task result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteIdentityProviderSCIMCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderSCIMCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetCode(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetHttpStatusCode(v int32) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetMessage(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetRequestId(v string) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) SetSuccess(v bool) *DeleteIdentityProviderSCIMCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIdentityProviderSCIMCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
