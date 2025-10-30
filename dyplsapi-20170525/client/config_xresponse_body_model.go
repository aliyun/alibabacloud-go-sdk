// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigXResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ConfigXResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ConfigXResponseBody
	GetCode() *string
	SetData(v *ConfigXResponseBodyData) *ConfigXResponseBody
	GetData() *ConfigXResponseBodyData
	SetMessage(v string) *ConfigXResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigXResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigXResponseBody
	GetSuccess() *bool
}

type ConfigXResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ConfigXResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigXResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ConfigXResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigXResponseBody) GetData() *ConfigXResponseBodyData {
	return s.Data
}

func (s *ConfigXResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigXResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigXResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigXResponseBody) SetAccessDeniedDetail(v string) *ConfigXResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ConfigXResponseBody) SetCode(v string) *ConfigXResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBody) SetData(v *ConfigXResponseBodyData) *ConfigXResponseBody {
	s.Data = v
	return s
}

func (s *ConfigXResponseBody) SetMessage(v string) *ConfigXResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBody) SetRequestId(v string) *ConfigXResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigXResponseBody) SetSuccess(v bool) *ConfigXResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigXResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigXResponseBodyData struct {
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回是否成功 true  表示成功 false表示失败
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigXResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigXResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigXResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ConfigXResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ConfigXResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigXResponseBodyData) SetCode(v string) *ConfigXResponseBodyData {
	s.Code = &v
	return s
}

func (s *ConfigXResponseBodyData) SetMessage(v string) *ConfigXResponseBodyData {
	s.Message = &v
	return s
}

func (s *ConfigXResponseBodyData) SetSuccess(v bool) *ConfigXResponseBodyData {
	s.Success = &v
	return s
}

func (s *ConfigXResponseBodyData) Validate() error {
	return dara.Validate(s)
}
