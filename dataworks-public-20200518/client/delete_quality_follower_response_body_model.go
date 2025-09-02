// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityFollowerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteQualityFollowerResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteQualityFollowerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteQualityFollowerResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteQualityFollowerResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteQualityFollowerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityFollowerResponseBody
	GetSuccess() *bool
}

type DeleteQualityFollowerResponseBody struct {
	// Indicates whether the subscriber was successfully deleted. Valid values:
	//
	// 	- true: The subscriber was successfully deleted.
	//
	// 	- false: The subscriber failed to be deleted. You can troubleshoot errors based on the error message returned.
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
	// The error message returned when the subscriber failed to be deleted.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can troubleshoot errors based on the ID.
	//
	// example:
	//
	// 6d739ef6-098a-47****
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
}

func (s DeleteQualityFollowerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityFollowerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityFollowerResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteQualityFollowerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteQualityFollowerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteQualityFollowerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityFollowerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityFollowerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityFollowerResponseBody) SetData(v bool) *DeleteQualityFollowerResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) SetErrorCode(v string) *DeleteQualityFollowerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) SetErrorMessage(v string) *DeleteQualityFollowerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) SetHttpStatusCode(v int32) *DeleteQualityFollowerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) SetRequestId(v string) *DeleteQualityFollowerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) SetSuccess(v bool) *DeleteQualityFollowerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityFollowerResponseBody) Validate() error {
	return dara.Validate(s)
}
