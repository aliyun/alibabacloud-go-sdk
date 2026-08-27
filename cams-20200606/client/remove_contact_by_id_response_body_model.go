// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContactByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RemoveContactByIdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RemoveContactByIdResponseBody
	GetCode() *string
	SetData(v string) *RemoveContactByIdResponseBody
	GetData() *string
	SetMessage(v string) *RemoveContactByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveContactByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveContactByIdResponseBody
	GetSuccess() *bool
}

type RemoveContactByIdResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// SampleValueSampleValueSampleValue.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// SampleValueSampleValue.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. A success message is returned if the request succeeds. A failure reason is returned if the request fails.
	//
	// example:
	//
	// SampleValueSampleValueSampleValue.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// SampleValueSampleValueSampleValue.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveContactByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveContactByIdResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveContactByIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RemoveContactByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveContactByIdResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveContactByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveContactByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveContactByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveContactByIdResponseBody) SetAccessDeniedDetail(v string) *RemoveContactByIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetCode(v string) *RemoveContactByIdResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetData(v string) *RemoveContactByIdResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetMessage(v string) *RemoveContactByIdResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetRequestId(v string) *RemoveContactByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetSuccess(v bool) *RemoveContactByIdResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveContactByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
