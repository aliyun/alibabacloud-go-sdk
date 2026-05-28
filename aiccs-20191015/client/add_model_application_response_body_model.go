// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddModelApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddModelApplicationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddModelApplicationResponseBody
	GetCode() *string
	SetData(v string) *AddModelApplicationResponseBody
	GetData() *string
	SetMessage(v string) *AddModelApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddModelApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddModelApplicationResponseBody
	GetSuccess() *bool
}

type AddModelApplicationResponseBody struct {
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
	// 示例值示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddModelApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *AddModelApplicationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddModelApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddModelApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *AddModelApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddModelApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddModelApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddModelApplicationResponseBody) SetAccessDeniedDetail(v string) *AddModelApplicationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetCode(v string) *AddModelApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetData(v string) *AddModelApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetMessage(v string) *AddModelApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetRequestId(v string) *AddModelApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetSuccess(v bool) *AddModelApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *AddModelApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
