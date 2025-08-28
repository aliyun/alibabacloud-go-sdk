// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseVideoFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PauseVideoFileResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PauseVideoFileResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *PauseVideoFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *PauseVideoFileResponseBody
	GetMessage() *string
	SetSuccess(v bool) *PauseVideoFileResponseBody
	GetSuccess() *bool
}

type PauseVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseVideoFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *PauseVideoFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PauseVideoFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseVideoFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *PauseVideoFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseVideoFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseVideoFileResponseBody) SetAccessDeniedDetail(v string) *PauseVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetCode(v string) *PauseVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetData(v map[string]interface{}) *PauseVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *PauseVideoFileResponseBody) SetMessage(v string) *PauseVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetSuccess(v bool) *PauseVideoFileResponseBody {
	s.Success = &v
	return s
}

func (s *PauseVideoFileResponseBody) Validate() error {
	return dara.Validate(s)
}
