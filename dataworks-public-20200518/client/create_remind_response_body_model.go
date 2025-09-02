// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateRemindResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateRemindResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateRemindResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateRemindResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateRemindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRemindResponseBody
	GetSuccess() *bool
}

type CreateRemindResponseBody struct {
	// The ID of the custom alert rule.
	//
	// example:
	//
	// 1234
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
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
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRemindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRemindResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemindResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateRemindResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateRemindResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateRemindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRemindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRemindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRemindResponseBody) SetData(v int64) *CreateRemindResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRemindResponseBody) SetErrorCode(v string) *CreateRemindResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRemindResponseBody) SetErrorMessage(v string) *CreateRemindResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRemindResponseBody) SetHttpStatusCode(v int32) *CreateRemindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRemindResponseBody) SetRequestId(v string) *CreateRemindResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRemindResponseBody) SetSuccess(v bool) *CreateRemindResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRemindResponseBody) Validate() error {
	return dara.Validate(s)
}
