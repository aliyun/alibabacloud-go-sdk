// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTableLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTableLevelResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateTableLevelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateTableLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTableLevelResponseBody
	GetSuccess() *bool
	SetUpdateResult(v bool) *UpdateTableLevelResponseBody
	GetUpdateResult() *bool
}

type UpdateTableLevelResponseBody struct {
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
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether the table level is updated.
	//
	// example:
	//
	// true
	UpdateResult *bool `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty"`
}

func (s UpdateTableLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableLevelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTableLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTableLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTableLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTableLevelResponseBody) GetUpdateResult() *bool {
	return s.UpdateResult
}

func (s *UpdateTableLevelResponseBody) SetErrorCode(v string) *UpdateTableLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTableLevelResponseBody) SetErrorMessage(v string) *UpdateTableLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTableLevelResponseBody) SetHttpStatusCode(v int32) *UpdateTableLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTableLevelResponseBody) SetRequestId(v string) *UpdateTableLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableLevelResponseBody) SetSuccess(v bool) *UpdateTableLevelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableLevelResponseBody) SetUpdateResult(v bool) *UpdateTableLevelResponseBody {
	s.UpdateResult = &v
	return s
}

func (s *UpdateTableLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
