// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateProxyAccessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateProxyAccessResponseBody
	GetErrorMessage() *string
	SetProxyAccessId(v int64) *CreateProxyAccessResponseBody
	GetProxyAccessId() *int64
	SetRequestId(v string) *CreateProxyAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProxyAccessResponseBody
	GetSuccess() *bool
}

type CreateProxyAccessResponseBody struct {
	// The error code returned to the query task.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the security protection authorization. After the security protection agent authorizes the target user, the system automatically generates a security protection authorization ID. The ID is globally unique.
	//
	// example:
	//
	// 1
	ProxyAccessId *int64 `json:"ProxyAccessId,omitempty" xml:"ProxyAccessId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProxyAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyAccessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProxyAccessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProxyAccessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateProxyAccessResponseBody) GetProxyAccessId() *int64 {
	return s.ProxyAccessId
}

func (s *CreateProxyAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProxyAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProxyAccessResponseBody) SetErrorCode(v string) *CreateProxyAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProxyAccessResponseBody) SetErrorMessage(v string) *CreateProxyAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProxyAccessResponseBody) SetProxyAccessId(v int64) *CreateProxyAccessResponseBody {
	s.ProxyAccessId = &v
	return s
}

func (s *CreateProxyAccessResponseBody) SetRequestId(v string) *CreateProxyAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProxyAccessResponseBody) SetSuccess(v bool) *CreateProxyAccessResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProxyAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
