// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplatePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyChatappTemplatePropertiesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyChatappTemplatePropertiesResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyChatappTemplatePropertiesResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *ModifyChatappTemplatePropertiesResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *ModifyChatappTemplatePropertiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyChatappTemplatePropertiesResponseBody
	GetSuccess() *bool
}

type ModifyChatappTemplatePropertiesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyChatappTemplatePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplatePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyChatappTemplatePropertiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetAccessDeniedDetail(v string) *ModifyChatappTemplatePropertiesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetCode(v string) *ModifyChatappTemplatePropertiesResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetMessage(v string) *ModifyChatappTemplatePropertiesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetModel(v map[string]interface{}) *ModifyChatappTemplatePropertiesResponseBody {
	s.Model = v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetRequestId(v string) *ModifyChatappTemplatePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) SetSuccess(v bool) *ModifyChatappTemplatePropertiesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponseBody) Validate() error {
	return dara.Validate(s)
}
