// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteProxyAccessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteProxyAccessResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteProxyAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProxyAccessResponseBody
	GetSuccess() *bool
}

type DeleteProxyAccessResponseBody struct {
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

func (s DeleteProxyAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProxyAccessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteProxyAccessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteProxyAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProxyAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProxyAccessResponseBody) SetErrorCode(v string) *DeleteProxyAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProxyAccessResponseBody) SetErrorMessage(v string) *DeleteProxyAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProxyAccessResponseBody) SetRequestId(v string) *DeleteProxyAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProxyAccessResponseBody) SetSuccess(v bool) *DeleteProxyAccessResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProxyAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
