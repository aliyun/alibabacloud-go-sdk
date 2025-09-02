// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateDataServiceApiResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateDataServiceApiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataServiceApiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataServiceApiResponseBody
	GetSuccess() *bool
}

type UpdateDataServiceApiResponseBody struct {
	// Indicates whether the information about the API was updated.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceApiResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataServiceApiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataServiceApiResponseBody) SetData(v bool) *UpdateDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) SetErrorCode(v string) *UpdateDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) SetErrorMessage(v string) *UpdateDataServiceApiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) SetHttpStatusCode(v int32) *UpdateDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) SetRequestId(v string) *UpdateDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) SetSuccess(v bool) *UpdateDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
