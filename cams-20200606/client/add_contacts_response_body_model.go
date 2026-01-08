// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddContactsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddContactsResponseBody
	GetCode() *string
	SetData(v string) *AddContactsResponseBody
	GetData() *string
	SetMessage(v string) *AddContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddContactsResponseBody
	GetSuccess() *bool
}

type AddContactsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddContactsResponseBody) GoString() string {
	return s.String()
}

func (s *AddContactsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddContactsResponseBody) GetData() *string {
	return s.Data
}

func (s *AddContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddContactsResponseBody) SetAccessDeniedDetail(v string) *AddContactsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddContactsResponseBody) SetCode(v string) *AddContactsResponseBody {
	s.Code = &v
	return s
}

func (s *AddContactsResponseBody) SetData(v string) *AddContactsResponseBody {
	s.Data = &v
	return s
}

func (s *AddContactsResponseBody) SetMessage(v string) *AddContactsResponseBody {
	s.Message = &v
	return s
}

func (s *AddContactsResponseBody) SetRequestId(v string) *AddContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContactsResponseBody) SetSuccess(v bool) *AddContactsResponseBody {
	s.Success = &v
	return s
}

func (s *AddContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
