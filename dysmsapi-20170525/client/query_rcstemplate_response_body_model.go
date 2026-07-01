// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRCSTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryRCSTemplateResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryRCSTemplateResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryRCSTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRCSTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRCSTemplateResponseBody
	GetSuccess() *bool
}

type QueryRCSTemplateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
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

func (s QueryRCSTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRCSTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRCSTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRCSTemplateResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryRCSTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRCSTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRCSTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRCSTemplateResponseBody) SetAccessDeniedDetail(v string) *QueryRCSTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRCSTemplateResponseBody) SetCode(v string) *QueryRCSTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRCSTemplateResponseBody) SetData(v map[string]interface{}) *QueryRCSTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryRCSTemplateResponseBody) SetMessage(v string) *QueryRCSTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRCSTemplateResponseBody) SetRequestId(v string) *QueryRCSTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRCSTemplateResponseBody) SetSuccess(v bool) *QueryRCSTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRCSTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
