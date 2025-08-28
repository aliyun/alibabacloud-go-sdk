// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDegradeVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DegradeVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DegradeVideoFileResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *DegradeVideoFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *DegradeVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *DegradeVideoFileResponseBody
	GetSuccess() *bool
}

type DegradeVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DegradeVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DegradeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DegradeVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DegradeVideoFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *DegradeVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DegradeVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DegradeVideoFileResponseBody) SetAccessDeniedDetail(v string) *DegradeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetCode(v string) *DegradeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetData(v map[string]interface{}) *DegradeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *DegradeVideoFileResponseBody) SetMessage(v string) *DegradeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetSuccess(v bool) *DegradeVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *DegradeVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
