// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRCSTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateRCSTemplateResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateRCSTemplateResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateRCSTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRCSTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRCSTemplateResponseBody
	GetSuccess() *bool
}

type CreateRCSTemplateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateRCSTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCSTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRCSTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRCSTemplateResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateRCSTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRCSTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCSTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRCSTemplateResponseBody) SetAccessDeniedDetail(v string) *CreateRCSTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRCSTemplateResponseBody) SetCode(v string) *CreateRCSTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRCSTemplateResponseBody) SetData(v map[string]interface{}) *CreateRCSTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateRCSTemplateResponseBody) SetMessage(v string) *CreateRCSTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRCSTemplateResponseBody) SetRequestId(v string) *CreateRCSTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCSTemplateResponseBody) SetSuccess(v bool) *CreateRCSTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRCSTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
