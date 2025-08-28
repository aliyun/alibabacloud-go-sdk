// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeekVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SeekVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SeekVideoFileResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *SeekVideoFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *SeekVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *SeekVideoFileResponseBody
	GetSuccess() *bool
}

type SeekVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SeekVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SeekVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *SeekVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SeekVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *SeekVideoFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SeekVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SeekVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SeekVideoFileResponseBody) SetAccessDeniedDetail(v string) *SeekVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetCode(v string) *SeekVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetData(v map[string]interface{}) *SeekVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *SeekVideoFileResponseBody) SetMessage(v string) *SeekVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetSuccess(v bool) *SeekVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *SeekVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
