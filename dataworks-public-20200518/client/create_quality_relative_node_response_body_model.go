// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRelativeNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateQualityRelativeNodeResponseBody
	GetData() *bool
	SetErrorCode(v string) *CreateQualityRelativeNodeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateQualityRelativeNodeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateQualityRelativeNodeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateQualityRelativeNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityRelativeNodeResponseBody
	GetSuccess() *bool
}

type CreateQualityRelativeNodeResponseBody struct {
	// Indicates whether the node is associated with the partition filter expression.
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

func (s CreateQualityRelativeNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRelativeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityRelativeNodeResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateQualityRelativeNodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateQualityRelativeNodeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateQualityRelativeNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateQualityRelativeNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityRelativeNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityRelativeNodeResponseBody) SetData(v bool) *CreateQualityRelativeNodeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) SetErrorCode(v string) *CreateQualityRelativeNodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) SetErrorMessage(v string) *CreateQualityRelativeNodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) SetHttpStatusCode(v int32) *CreateQualityRelativeNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) SetRequestId(v string) *CreateQualityRelativeNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) SetSuccess(v bool) *CreateQualityRelativeNodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityRelativeNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
