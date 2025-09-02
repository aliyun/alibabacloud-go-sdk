// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteQualityEntityResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteQualityEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteQualityEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteQualityEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteQualityEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityEntityResponseBody
	GetSuccess() *bool
}

type DeleteQualityEntityResponseBody struct {
	// The returned result.
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
	// You have no permission.
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
	// ecb967ec-c137-48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityEntityResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteQualityEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteQualityEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteQualityEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityEntityResponseBody) SetData(v bool) *DeleteQualityEntityResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) SetErrorCode(v string) *DeleteQualityEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) SetErrorMessage(v string) *DeleteQualityEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) SetHttpStatusCode(v int32) *DeleteQualityEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) SetRequestId(v string) *DeleteQualityEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) SetSuccess(v bool) *DeleteQualityEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
