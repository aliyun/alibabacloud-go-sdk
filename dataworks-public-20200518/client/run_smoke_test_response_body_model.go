// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSmokeTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *RunSmokeTestResponseBody
	GetData() *int64
	SetErrorCode(v string) *RunSmokeTestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RunSmokeTestResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *RunSmokeTestResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RunSmokeTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunSmokeTestResponseBody
	GetSuccess() *bool
}

type RunSmokeTestResponseBody struct {
	// The workflow ID.
	//
	// example:
	//
	// 3333333
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunSmokeTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSmokeTestResponseBody) GoString() string {
	return s.String()
}

func (s *RunSmokeTestResponseBody) GetData() *int64 {
	return s.Data
}

func (s *RunSmokeTestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunSmokeTestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunSmokeTestResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RunSmokeTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSmokeTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunSmokeTestResponseBody) SetData(v int64) *RunSmokeTestResponseBody {
	s.Data = &v
	return s
}

func (s *RunSmokeTestResponseBody) SetErrorCode(v string) *RunSmokeTestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RunSmokeTestResponseBody) SetErrorMessage(v string) *RunSmokeTestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RunSmokeTestResponseBody) SetHttpStatusCode(v int32) *RunSmokeTestResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunSmokeTestResponseBody) SetRequestId(v string) *RunSmokeTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSmokeTestResponseBody) SetSuccess(v bool) *RunSmokeTestResponseBody {
	s.Success = &v
	return s
}

func (s *RunSmokeTestResponseBody) Validate() error {
	return dara.Validate(s)
}
