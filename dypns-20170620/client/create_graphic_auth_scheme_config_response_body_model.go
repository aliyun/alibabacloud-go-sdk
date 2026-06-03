// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphicAuthSchemeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateGraphicAuthSchemeConfigResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateGraphicAuthSchemeConfigResponseBody
	GetCode() *string
	SetMessage(v string) *CreateGraphicAuthSchemeConfigResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *CreateGraphicAuthSchemeConfigResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *CreateGraphicAuthSchemeConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGraphicAuthSchemeConfigResponseBody
	GetSuccess() *bool
}

type CreateGraphicAuthSchemeConfigResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// "Model": {
	//
	//         "SchemeCode": "FG0000*******04087"
	//
	//     }
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGraphicAuthSchemeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphicAuthSchemeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetAccessDeniedDetail(v string) *CreateGraphicAuthSchemeConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetCode(v string) *CreateGraphicAuthSchemeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetMessage(v string) *CreateGraphicAuthSchemeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetModel(v map[string]interface{}) *CreateGraphicAuthSchemeConfigResponseBody {
	s.Model = v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetRequestId(v string) *CreateGraphicAuthSchemeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) SetSuccess(v bool) *CreateGraphicAuthSchemeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
