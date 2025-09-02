// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SuspendInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *SuspendInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SuspendInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SuspendInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SuspendInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendInstanceResponseBody
	GetSuccess() *bool
}

type SuspendInstanceResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ProjectNotExists
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
	// 0baf87f0159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *SuspendInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SuspendInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SuspendInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendInstanceResponseBody) SetData(v bool) *SuspendInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendInstanceResponseBody) SetErrorCode(v string) *SuspendInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SuspendInstanceResponseBody) SetErrorMessage(v string) *SuspendInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SuspendInstanceResponseBody) SetHttpStatusCode(v int32) *SuspendInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendInstanceResponseBody) SetRequestId(v string) *SuspendInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendInstanceResponseBody) SetSuccess(v bool) *SuspendInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
