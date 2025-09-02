// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTableThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTableThemeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateTableThemeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateTableThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTableThemeResponseBody
	GetSuccess() *bool
	SetUpdateResult(v bool) *UpdateTableThemeResponseBody
	GetUpdateResult() *bool
}

type UpdateTableThemeResponseBody struct {
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
	// abcd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether the update result is returned. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	UpdateResult *bool `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty"`
}

func (s UpdateTableThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableThemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTableThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTableThemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTableThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTableThemeResponseBody) GetUpdateResult() *bool {
	return s.UpdateResult
}

func (s *UpdateTableThemeResponseBody) SetErrorCode(v string) *UpdateTableThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTableThemeResponseBody) SetErrorMessage(v string) *UpdateTableThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTableThemeResponseBody) SetHttpStatusCode(v int32) *UpdateTableThemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTableThemeResponseBody) SetRequestId(v string) *UpdateTableThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableThemeResponseBody) SetSuccess(v bool) *UpdateTableThemeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableThemeResponseBody) SetUpdateResult(v bool) *UpdateTableThemeResponseBody {
	s.UpdateResult = &v
	return s
}

func (s *UpdateTableThemeResponseBody) Validate() error {
	return dara.Validate(s)
}
