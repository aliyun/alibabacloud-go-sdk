// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLargeModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddLargeModelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddLargeModelResponseBody
	GetCode() *string
	SetData(v bool) *AddLargeModelResponseBody
	GetData() *bool
	SetMessage(v string) *AddLargeModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddLargeModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddLargeModelResponseBody
	GetSuccess() *bool
}

type AddLargeModelResponseBody struct {
	// The reason for the authentication failure.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. A value of `OK` indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1B8E483-372F-5AA8-A4B2-CA82EC967B0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddLargeModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLargeModelResponseBody) GoString() string {
	return s.String()
}

func (s *AddLargeModelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddLargeModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddLargeModelResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddLargeModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddLargeModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLargeModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddLargeModelResponseBody) SetAccessDeniedDetail(v string) *AddLargeModelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddLargeModelResponseBody) SetCode(v string) *AddLargeModelResponseBody {
	s.Code = &v
	return s
}

func (s *AddLargeModelResponseBody) SetData(v bool) *AddLargeModelResponseBody {
	s.Data = &v
	return s
}

func (s *AddLargeModelResponseBody) SetMessage(v string) *AddLargeModelResponseBody {
	s.Message = &v
	return s
}

func (s *AddLargeModelResponseBody) SetRequestId(v string) *AddLargeModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLargeModelResponseBody) SetSuccess(v bool) *AddLargeModelResponseBody {
	s.Success = &v
	return s
}

func (s *AddLargeModelResponseBody) Validate() error {
	return dara.Validate(s)
}
