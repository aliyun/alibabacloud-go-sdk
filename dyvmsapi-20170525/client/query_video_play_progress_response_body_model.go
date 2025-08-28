// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoPlayProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryVideoPlayProgressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryVideoPlayProgressResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryVideoPlayProgressResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryVideoPlayProgressResponseBody
	GetMessage() *string
	SetSuccess(v bool) *QueryVideoPlayProgressResponseBody
	GetSuccess() *bool
}

type QueryVideoPlayProgressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVideoPlayProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoPlayProgressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryVideoPlayProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVideoPlayProgressResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryVideoPlayProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryVideoPlayProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVideoPlayProgressResponseBody) SetAccessDeniedDetail(v string) *QueryVideoPlayProgressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetCode(v string) *QueryVideoPlayProgressResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetData(v map[string]interface{}) *QueryVideoPlayProgressResponseBody {
	s.Data = v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetMessage(v string) *QueryVideoPlayProgressResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetSuccess(v bool) *QueryVideoPlayProgressResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
