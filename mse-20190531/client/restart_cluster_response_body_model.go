// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RestartClusterResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RestartClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartClusterResponseBody
	GetSuccess() *bool
}

type RestartClusterResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
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
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
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

func (s RestartClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartClusterResponseBody) SetErrorCode(v string) *RestartClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartClusterResponseBody) SetMessage(v string) *RestartClusterResponseBody {
	s.Message = &v
	return s
}

func (s *RestartClusterResponseBody) SetRequestId(v string) *RestartClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartClusterResponseBody) SetSuccess(v bool) *RestartClusterResponseBody {
	s.Success = &v
	return s
}

func (s *RestartClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
