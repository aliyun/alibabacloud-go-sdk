// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCallProgressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCallProgressResponseBody
	GetCode() *string
	SetMessage(v string) *GetCallProgressResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *GetCallProgressResponseBody
	GetModel() map[string]interface{}
	SetSuccess(v bool) *GetCallProgressResponseBody
	GetSuccess() *bool
}

type GetCallProgressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallProgressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCallProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallProgressResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *GetCallProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallProgressResponseBody) SetAccessDeniedDetail(v string) *GetCallProgressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCallProgressResponseBody) SetCode(v string) *GetCallProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallProgressResponseBody) SetMessage(v string) *GetCallProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallProgressResponseBody) SetModel(v map[string]interface{}) *GetCallProgressResponseBody {
	s.Model = v
	return s
}

func (s *GetCallProgressResponseBody) SetSuccess(v bool) *GetCallProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
