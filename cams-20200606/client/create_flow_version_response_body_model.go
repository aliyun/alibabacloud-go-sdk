// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *CreateFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *CreateFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *CreateFlowVersionResponseBody
	GetSuccess() *bool
}

type CreateFlowVersionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Response  map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *CreateFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowVersionResponseBody) SetAccessDeniedDetail(v string) *CreateFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetCode(v string) *CreateFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetMessage(v string) *CreateFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetRequestId(v string) *CreateFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetResponse(v map[string]interface{}) *CreateFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *CreateFlowVersionResponseBody) SetSuccess(v bool) *CreateFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
