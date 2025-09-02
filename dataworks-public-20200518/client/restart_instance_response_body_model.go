// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RestartInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *RestartInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RestartInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *RestartInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RestartInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartInstanceResponseBody
	GetSuccess() *bool
}

type RestartInstanceResponseBody struct {
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
	// The request ID.
	//
	// example:
	//
	// q32ety****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RestartInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartInstanceResponseBody) SetData(v bool) *RestartInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorCode(v string) *RestartInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorMessage(v string) *RestartInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RestartInstanceResponseBody) SetHttpStatusCode(v int32) *RestartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetSuccess(v bool) *RestartInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RestartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
