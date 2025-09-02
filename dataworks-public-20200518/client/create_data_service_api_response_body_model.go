// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateDataServiceApiResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateDataServiceApiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataServiceApiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataServiceApiResponseBody
	GetSuccess() *bool
}

type CreateDataServiceApiResponseBody struct {
	// The ID of the API.
	//
	// example:
	//
	// 100003
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
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

func (s CreateDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataServiceApiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataServiceApiResponseBody) SetData(v int64) *CreateDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetErrorCode(v string) *CreateDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetErrorMessage(v string) *CreateDataServiceApiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetHttpStatusCode(v int32) *CreateDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetRequestId(v string) *CreateDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetSuccess(v bool) *CreateDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
