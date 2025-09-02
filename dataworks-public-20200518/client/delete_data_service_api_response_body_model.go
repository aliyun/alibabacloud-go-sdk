// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteDataServiceApiResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteDataServiceApiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataServiceApiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataServiceApiResponseBody
	GetSuccess() *bool
}

type DeleteDataServiceApiResponseBody struct {
	// Indicates whether the API was deleted.
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

func (s DeleteDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataServiceApiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataServiceApiResponseBody) SetData(v bool) *DeleteDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) SetErrorCode(v string) *DeleteDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) SetErrorMessage(v string) *DeleteDataServiceApiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) SetHttpStatusCode(v int32) *DeleteDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) SetRequestId(v string) *DeleteDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) SetSuccess(v bool) *DeleteDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
