// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PublishDataServiceApiResponseBody
	GetData() *bool
	SetErrorCode(v string) *PublishDataServiceApiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PublishDataServiceApiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *PublishDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *PublishDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishDataServiceApiResponseBody
	GetSuccess() *bool
}

type PublishDataServiceApiResponseBody struct {
	// Indicates whether the API was published.
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

func (s PublishDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDataServiceApiResponseBody) GetData() *bool {
	return s.Data
}

func (s *PublishDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PublishDataServiceApiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PublishDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishDataServiceApiResponseBody) SetData(v bool) *PublishDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetErrorCode(v string) *PublishDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetErrorMessage(v string) *PublishDataServiceApiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetHttpStatusCode(v int32) *PublishDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetRequestId(v string) *PublishDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) SetSuccess(v bool) *PublishDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *PublishDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
