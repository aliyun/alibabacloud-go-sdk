// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAXBCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAXBCallResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAXBCallResponseBody
	GetCode() *string
	SetData(v *BindAXBCallResponseBodyData) *BindAXBCallResponseBody
	GetData() *BindAXBCallResponseBodyData
	SetMessage(v string) *BindAXBCallResponseBody
	GetMessage() *string
	SetSuccess(v bool) *BindAXBCallResponseBody
	GetSuccess() *bool
}

type BindAXBCallResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAXBCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s BindAXBCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallResponseBody) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAXBCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAXBCallResponseBody) GetData() *BindAXBCallResponseBodyData {
	return s.Data
}

func (s *BindAXBCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAXBCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAXBCallResponseBody) SetAccessDeniedDetail(v string) *BindAXBCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAXBCallResponseBody) SetCode(v string) *BindAXBCallResponseBody {
	s.Code = &v
	return s
}

func (s *BindAXBCallResponseBody) SetData(v *BindAXBCallResponseBodyData) *BindAXBCallResponseBody {
	s.Data = v
	return s
}

func (s *BindAXBCallResponseBody) SetMessage(v string) *BindAXBCallResponseBody {
	s.Message = &v
	return s
}

func (s *BindAXBCallResponseBody) SetSuccess(v bool) *BindAXBCallResponseBody {
	s.Success = &v
	return s
}

func (s *BindAXBCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAXBCallResponseBodyData struct {
	// 绑定关系ID
	//
	// example:
	//
	// 476567566
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
}

func (s BindAXBCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponseBodyData) GetBindId() *string {
	return s.BindId
}

func (s *BindAXBCallResponseBodyData) SetBindId(v string) *BindAXBCallResponseBodyData {
	s.BindId = &v
	return s
}

func (s *BindAXBCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
