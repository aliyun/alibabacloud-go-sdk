// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *SubmitFileResponseBody
	GetData() *int64
	SetErrorCode(v string) *SubmitFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SubmitFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SubmitFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SubmitFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitFileResponseBody
	GetSuccess() *bool
}

type SubmitFileResponseBody struct {
	// example:
	//
	// 3000001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitFileResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFileResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SubmitFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitFileResponseBody) SetData(v int64) *SubmitFileResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitFileResponseBody) SetErrorCode(v string) *SubmitFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitFileResponseBody) SetErrorMessage(v string) *SubmitFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitFileResponseBody) SetHttpStatusCode(v int32) *SubmitFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitFileResponseBody) SetRequestId(v string) *SubmitFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFileResponseBody) SetSuccess(v bool) *SubmitFileResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitFileResponseBody) Validate() error {
	return dara.Validate(s)
}
