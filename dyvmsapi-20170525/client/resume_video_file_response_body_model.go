// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ResumeVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ResumeVideoFileResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *ResumeVideoFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *ResumeVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *ResumeVideoFileResponseBody
	GetSuccess() *bool
}

type ResumeVideoFileResponseBody struct {
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ResumeVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeVideoFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ResumeVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeVideoFileResponseBody) SetAccessDeniedDetail(v string) *ResumeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetCode(v string) *ResumeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetData(v map[string]interface{}) *ResumeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *ResumeVideoFileResponseBody) SetMessage(v string) *ResumeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetSuccess(v bool) *ResumeVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
