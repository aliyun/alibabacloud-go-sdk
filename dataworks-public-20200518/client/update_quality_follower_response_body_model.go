// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityFollowerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateQualityFollowerResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateQualityFollowerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateQualityFollowerResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateQualityFollowerResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateQualityFollowerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityFollowerResponseBody
	GetSuccess() *bool
}

type UpdateQualityFollowerResponseBody struct {
	// Indicates whether the update is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 401
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You have no permission
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
	// 576b9457-2cf5-4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityFollowerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityFollowerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityFollowerResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateQualityFollowerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateQualityFollowerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateQualityFollowerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateQualityFollowerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityFollowerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityFollowerResponseBody) SetData(v bool) *UpdateQualityFollowerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) SetErrorCode(v string) *UpdateQualityFollowerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) SetErrorMessage(v string) *UpdateQualityFollowerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) SetHttpStatusCode(v int32) *UpdateQualityFollowerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) SetRequestId(v string) *UpdateQualityFollowerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) SetSuccess(v bool) *UpdateQualityFollowerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityFollowerResponseBody) Validate() error {
	return dara.Validate(s)
}
