// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityFollowerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *CreateQualityFollowerResponseBody
	GetData() *int32
	SetErrorCode(v string) *CreateQualityFollowerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateQualityFollowerResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateQualityFollowerResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateQualityFollowerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityFollowerResponseBody
	GetSuccess() *bool
}

type CreateQualityFollowerResponseBody struct {
	// The ID of the subscription relationship.
	//
	// example:
	//
	// 12345
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateQualityFollowerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityFollowerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityFollowerResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateQualityFollowerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateQualityFollowerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateQualityFollowerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateQualityFollowerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityFollowerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityFollowerResponseBody) SetData(v int32) *CreateQualityFollowerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) SetErrorCode(v string) *CreateQualityFollowerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) SetErrorMessage(v string) *CreateQualityFollowerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) SetHttpStatusCode(v int32) *CreateQualityFollowerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) SetRequestId(v string) *CreateQualityFollowerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) SetSuccess(v bool) *CreateQualityFollowerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityFollowerResponseBody) Validate() error {
	return dara.Validate(s)
}
