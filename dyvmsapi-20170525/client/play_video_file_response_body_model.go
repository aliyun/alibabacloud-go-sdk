// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PlayVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PlayVideoFileResponseBody
	GetCode() *string
	SetMessage(v string) *PlayVideoFileResponseBody
	GetMessage() *string
	SetModel(v bool) *PlayVideoFileResponseBody
	GetModel() *bool
	SetSuccess(v bool) *PlayVideoFileResponseBody
	GetSuccess() *bool
}

type PlayVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Model *bool `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PlayVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *PlayVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PlayVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *PlayVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PlayVideoFileResponseBody) GetModel() *bool {
	return s.Model
}

func (s *PlayVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PlayVideoFileResponseBody) SetAccessDeniedDetail(v string) *PlayVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetCode(v string) *PlayVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetMessage(v string) *PlayVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetModel(v bool) *PlayVideoFileResponseBody {
	s.Model = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetSuccess(v bool) *PlayVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *PlayVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
