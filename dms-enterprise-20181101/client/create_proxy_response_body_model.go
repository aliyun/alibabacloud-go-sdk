// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateProxyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateProxyResponseBody
	GetErrorMessage() *string
	SetProxyId(v int64) *CreateProxyResponseBody
	GetProxyId() *int64
	SetRequestId(v string) *CreateProxyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProxyResponseBody
	GetSuccess() *bool
}

type CreateProxyResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// the instance proxy already exists.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the secure access proxy.
	//
	// example:
	//
	// 4**
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4FFD154E-F57F-5374-B568-D6276F15****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProxyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProxyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateProxyResponseBody) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *CreateProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProxyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProxyResponseBody) SetErrorCode(v string) *CreateProxyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProxyResponseBody) SetErrorMessage(v string) *CreateProxyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProxyResponseBody) SetProxyId(v int64) *CreateProxyResponseBody {
	s.ProxyId = &v
	return s
}

func (s *CreateProxyResponseBody) SetRequestId(v string) *CreateProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProxyResponseBody) SetSuccess(v bool) *CreateProxyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
