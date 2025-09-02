// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SubmitDataServiceApiResponseBody
	GetData() *bool
	SetErrorCode(v string) *SubmitDataServiceApiResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *SubmitDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SubmitDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDataServiceApiResponseBody
	GetSuccess() *bool
}

type SubmitDataServiceApiResponseBody struct {
	// Indicates whether the API was submitted.
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

func (s SubmitDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDataServiceApiResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitDataServiceApiResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDataServiceApiResponseBody) SetData(v bool) *SubmitDataServiceApiResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitDataServiceApiResponseBody) SetErrorCode(v string) *SubmitDataServiceApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitDataServiceApiResponseBody) SetHttpStatusCode(v int32) *SubmitDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDataServiceApiResponseBody) SetRequestId(v string) *SubmitDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDataServiceApiResponseBody) SetSuccess(v bool) *SubmitDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}
