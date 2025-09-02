// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *CreateQualityEntityResponseBody
	GetData() *int32
	SetErrorCode(v string) *CreateQualityEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateQualityEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateQualityEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateQualityEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityEntityResponseBody
	GetSuccess() *bool
}

type CreateQualityEntityResponseBody struct {
	// The partition filter expression ID.
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
	// Your project is not relative with your account.
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
	// 011e1231u3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityEntityResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateQualityEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateQualityEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateQualityEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateQualityEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityEntityResponseBody) SetData(v int32) *CreateQualityEntityResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQualityEntityResponseBody) SetErrorCode(v string) *CreateQualityEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateQualityEntityResponseBody) SetErrorMessage(v string) *CreateQualityEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateQualityEntityResponseBody) SetHttpStatusCode(v int32) *CreateQualityEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateQualityEntityResponseBody) SetRequestId(v string) *CreateQualityEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityEntityResponseBody) SetSuccess(v bool) *CreateQualityEntityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
