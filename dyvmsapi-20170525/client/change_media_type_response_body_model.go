// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMediaTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChangeMediaTypeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChangeMediaTypeResponseBody
	GetCode() *string
	SetMessage(v string) *ChangeMediaTypeResponseBody
	GetMessage() *string
	SetModel(v bool) *ChangeMediaTypeResponseBody
	GetModel() *bool
	SetSuccess(v bool) *ChangeMediaTypeResponseBody
	GetSuccess() *bool
}

type ChangeMediaTypeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMediaTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMediaTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChangeMediaTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeMediaTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeMediaTypeResponseBody) GetModel() *bool {
	return s.Model
}

func (s *ChangeMediaTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMediaTypeResponseBody) SetAccessDeniedDetail(v string) *ChangeMediaTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetCode(v string) *ChangeMediaTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetMessage(v string) *ChangeMediaTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetModel(v bool) *ChangeMediaTypeResponseBody {
	s.Model = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetSuccess(v bool) *ChangeMediaTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
