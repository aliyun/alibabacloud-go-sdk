// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateModelApplicationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateModelApplicationResponseBody
	GetCode() *string
	SetData(v string) *UpdateModelApplicationResponseBody
	GetData() *string
	SetMessage(v string) *UpdateModelApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelApplicationResponseBody
	GetSuccess() *bool
}

type UpdateModelApplicationResponseBody struct {
	// The detailed reason for the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned by the request.
	//
	// example:
	//
	// DKSDLSA
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModelApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateModelApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateModelApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelApplicationResponseBody) SetAccessDeniedDetail(v string) *UpdateModelApplicationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) SetCode(v string) *UpdateModelApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) SetData(v string) *UpdateModelApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) SetMessage(v string) *UpdateModelApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) SetRequestId(v string) *UpdateModelApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) SetSuccess(v bool) *UpdateModelApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
