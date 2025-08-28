// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SkipVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SkipVideoFileResponseBody
	GetCode() *string
	SetData(v bool) *SkipVideoFileResponseBody
	GetData() *bool
	SetMessage(v string) *SkipVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *SkipVideoFileResponseBody
	GetSuccess() *bool
}

type SkipVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *SkipVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SkipVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *SkipVideoFileResponseBody) GetData() *bool {
	return s.Data
}

func (s *SkipVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SkipVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipVideoFileResponseBody) SetAccessDeniedDetail(v string) *SkipVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetCode(v string) *SkipVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetData(v bool) *SkipVideoFileResponseBody {
	s.Data = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetMessage(v string) *SkipVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetSuccess(v bool) *SkipVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *SkipVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
