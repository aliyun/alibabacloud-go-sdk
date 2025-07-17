// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteProxyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteProxyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteProxyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProxyResponseBody
	GetSuccess() *bool
}

type DeleteProxyResponseBody struct {
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

func (s DeleteProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProxyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteProxyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProxyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProxyResponseBody) SetErrorCode(v string) *DeleteProxyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProxyResponseBody) SetErrorMessage(v string) *DeleteProxyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProxyResponseBody) SetRequestId(v string) *DeleteProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProxyResponseBody) SetSuccess(v bool) *DeleteProxyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
