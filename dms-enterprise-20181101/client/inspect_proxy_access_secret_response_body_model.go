// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInspectProxyAccessSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessSecret(v string) *InspectProxyAccessSecretResponseBody
	GetAccessSecret() *string
	SetErrorCode(v string) *InspectProxyAccessSecretResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *InspectProxyAccessSecretResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *InspectProxyAccessSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InspectProxyAccessSecretResponseBody
	GetSuccess() *bool
}

type InspectProxyAccessSecretResponseBody struct {
	// The authorization password of the security protection agent.
	//
	// example:
	//
	// xxx
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
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

func (s InspectProxyAccessSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InspectProxyAccessSecretResponseBody) GoString() string {
	return s.String()
}

func (s *InspectProxyAccessSecretResponseBody) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *InspectProxyAccessSecretResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InspectProxyAccessSecretResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *InspectProxyAccessSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InspectProxyAccessSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InspectProxyAccessSecretResponseBody) SetAccessSecret(v string) *InspectProxyAccessSecretResponseBody {
	s.AccessSecret = &v
	return s
}

func (s *InspectProxyAccessSecretResponseBody) SetErrorCode(v string) *InspectProxyAccessSecretResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InspectProxyAccessSecretResponseBody) SetErrorMessage(v string) *InspectProxyAccessSecretResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InspectProxyAccessSecretResponseBody) SetRequestId(v string) *InspectProxyAccessSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *InspectProxyAccessSecretResponseBody) SetSuccess(v bool) *InspectProxyAccessSecretResponseBody {
	s.Success = &v
	return s
}

func (s *InspectProxyAccessSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
