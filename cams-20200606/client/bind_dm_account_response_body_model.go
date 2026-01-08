// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDmAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindDmAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindDmAccountResponseBody
	GetCode() *string
	SetData(v string) *BindDmAccountResponseBody
	GetData() *string
	SetMessage(v string) *BindDmAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindDmAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindDmAccountResponseBody
	GetSuccess() *bool
}

type BindDmAccountResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
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

func (s BindDmAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindDmAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BindDmAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindDmAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindDmAccountResponseBody) GetData() *string {
	return s.Data
}

func (s *BindDmAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindDmAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindDmAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindDmAccountResponseBody) SetAccessDeniedDetail(v string) *BindDmAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindDmAccountResponseBody) SetCode(v string) *BindDmAccountResponseBody {
	s.Code = &v
	return s
}

func (s *BindDmAccountResponseBody) SetData(v string) *BindDmAccountResponseBody {
	s.Data = &v
	return s
}

func (s *BindDmAccountResponseBody) SetMessage(v string) *BindDmAccountResponseBody {
	s.Message = &v
	return s
}

func (s *BindDmAccountResponseBody) SetRequestId(v string) *BindDmAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDmAccountResponseBody) SetSuccess(v bool) *BindDmAccountResponseBody {
	s.Success = &v
	return s
}

func (s *BindDmAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
