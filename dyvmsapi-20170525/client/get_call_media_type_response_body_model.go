// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallMediaTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCallMediaTypeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCallMediaTypeResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetCallMediaTypeResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetCallMediaTypeResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetCallMediaTypeResponseBody
	GetSuccess() *bool
}

type GetCallMediaTypeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallMediaTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallMediaTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCallMediaTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallMediaTypeResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetCallMediaTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallMediaTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallMediaTypeResponseBody) SetAccessDeniedDetail(v string) *GetCallMediaTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetCode(v string) *GetCallMediaTypeResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetData(v map[string]interface{}) *GetCallMediaTypeResponseBody {
	s.Data = v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetMessage(v string) *GetCallMediaTypeResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetSuccess(v bool) *GetCallMediaTypeResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
