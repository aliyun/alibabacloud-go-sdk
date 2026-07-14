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
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 399s88-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
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
