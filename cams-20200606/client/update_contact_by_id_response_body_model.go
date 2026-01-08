// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateContactByIdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateContactByIdResponseBody
	GetCode() *string
	SetData(v string) *UpdateContactByIdResponseBody
	GetData() *string
	SetMessage(v string) *UpdateContactByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContactByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContactByIdResponseBody
	GetSuccess() *bool
}

type UpdateContactByIdResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateContactByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactByIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactByIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateContactByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContactByIdResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateContactByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContactByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContactByIdResponseBody) SetAccessDeniedDetail(v string) *UpdateContactByIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateContactByIdResponseBody) SetCode(v string) *UpdateContactByIdResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContactByIdResponseBody) SetData(v string) *UpdateContactByIdResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateContactByIdResponseBody) SetMessage(v string) *UpdateContactByIdResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContactByIdResponseBody) SetRequestId(v string) *UpdateContactByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactByIdResponseBody) SetSuccess(v bool) *UpdateContactByIdResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContactByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
