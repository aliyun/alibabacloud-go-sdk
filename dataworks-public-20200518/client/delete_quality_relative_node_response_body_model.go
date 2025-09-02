// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRelativeNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteQualityRelativeNodeResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteQualityRelativeNodeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteQualityRelativeNodeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteQualityRelativeNodeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteQualityRelativeNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityRelativeNodeResponseBody
	GetSuccess() *bool
}

type DeleteQualityRelativeNodeResponseBody struct {
	// Indicates whether the node is disassociated from the partition filter expression.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// 401
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6d739ef6-098a-47****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityRelativeNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRelativeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityRelativeNodeResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteQualityRelativeNodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteQualityRelativeNodeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteQualityRelativeNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityRelativeNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityRelativeNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityRelativeNodeResponseBody) SetData(v bool) *DeleteQualityRelativeNodeResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) SetErrorCode(v string) *DeleteQualityRelativeNodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) SetErrorMessage(v string) *DeleteQualityRelativeNodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) SetHttpStatusCode(v int32) *DeleteQualityRelativeNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) SetRequestId(v string) *DeleteQualityRelativeNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) SetSuccess(v bool) *DeleteQualityRelativeNodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
