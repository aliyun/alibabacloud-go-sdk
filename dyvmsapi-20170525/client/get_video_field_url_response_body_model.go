// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoFieldUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetVideoFieldUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetVideoFieldUrlResponseBody
	GetCode() *string
	SetMessage(v string) *GetVideoFieldUrlResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *GetVideoFieldUrlResponseBody
	GetModel() map[string]interface{}
	SetSuccess(v bool) *GetVideoFieldUrlResponseBody
	GetSuccess() *bool
}

type GetVideoFieldUrlResponseBody struct {
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

func (s GetVideoFieldUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoFieldUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetVideoFieldUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoFieldUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoFieldUrlResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *GetVideoFieldUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoFieldUrlResponseBody) SetAccessDeniedDetail(v string) *GetVideoFieldUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetCode(v string) *GetVideoFieldUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetMessage(v string) *GetVideoFieldUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetModel(v map[string]interface{}) *GetVideoFieldUrlResponseBody {
	s.Model = v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetSuccess(v bool) *GetVideoFieldUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
