// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AbolishDataServiceApiResponseBody
	GetData() *bool
	SetErrorCode(v string) *AbolishDataServiceApiResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AbolishDataServiceApiResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *AbolishDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AbolishDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbolishDataServiceApiResponseBody
	GetSuccess() *bool
}

type AbolishDataServiceApiResponseBody struct {
	// Indicates whether the API is unpublished.
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

func (s AbolishDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbolishDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishDataServiceApiResponseBody) GetData() *bool {
	return s.Data
}

func (s *AbolishDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AbolishDataServiceApiResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AbolishDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AbolishDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbolishDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbolishDataServiceApiResponseBody) SetData(v bool) *AbolishDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) SetErrorCode(v string) *AbolishDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) SetErrorMessage(v string) *AbolishDataServiceApiResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) SetHttpStatusCode(v int32) *AbolishDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) SetRequestId(v string) *AbolishDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) SetSuccess(v bool) *AbolishDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *AbolishDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
