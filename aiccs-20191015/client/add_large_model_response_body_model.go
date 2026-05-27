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
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
