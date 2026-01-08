// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteContactsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteContactsResponseBody
	GetCode() *string
	SetData(v string) *DeleteContactsResponseBody
	GetData() *string
	SetMessage(v string) *DeleteContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContactsResponseBody
	GetSuccess() *bool
}

type DeleteContactsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 11111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactsResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContactsResponseBody) SetAccessDeniedDetail(v string) *DeleteContactsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteContactsResponseBody) SetCode(v string) *DeleteContactsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactsResponseBody) SetData(v string) *DeleteContactsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteContactsResponseBody) SetMessage(v string) *DeleteContactsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactsResponseBody) SetRequestId(v string) *DeleteContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactsResponseBody) SetSuccess(v bool) *DeleteContactsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
