// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteClusterResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *DeleteClusterResponseBody
	GetHttpCode() *string
	SetMessage(v string) *DeleteClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteClusterResponseBody
	GetSuccess() *bool
}

type DeleteClusterResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3369AD10-F1A6-4E6F-B99E-20F51826****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteClusterResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DeleteClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteClusterResponseBody) SetErrorCode(v string) *DeleteClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteClusterResponseBody) SetHttpCode(v string) *DeleteClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
