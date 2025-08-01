// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RetryClusterResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RetryClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryClusterResponseBody
	GetSuccess() *bool
}

type RetryClusterResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
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

func (s RetryClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RetryClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryClusterResponseBody) SetErrorCode(v string) *RetryClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryClusterResponseBody) SetMessage(v string) *RetryClusterResponseBody {
	s.Message = &v
	return s
}

func (s *RetryClusterResponseBody) SetRequestId(v string) *RetryClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryClusterResponseBody) SetSuccess(v bool) *RetryClusterResponseBody {
	s.Success = &v
	return s
}

func (s *RetryClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
