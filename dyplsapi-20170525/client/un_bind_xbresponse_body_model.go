// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindXBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnBindXBResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UnBindXBResponseBody
	GetCode() *string
	SetData(v *UnBindXBResponseBodyData) *UnBindXBResponseBody
	GetData() *UnBindXBResponseBodyData
	SetMessage(v string) *UnBindXBResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnBindXBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnBindXBResponseBody
	GetSuccess() *bool
}

type UnBindXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UnBindXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindXBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnBindXBResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindXBResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnBindXBResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnBindXBResponseBody) GetData() *UnBindXBResponseBodyData {
	return s.Data
}

func (s *UnBindXBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnBindXBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnBindXBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindXBResponseBody) SetAccessDeniedDetail(v string) *UnBindXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnBindXBResponseBody) SetCode(v string) *UnBindXBResponseBody {
	s.Code = &v
	return s
}

func (s *UnBindXBResponseBody) SetData(v *UnBindXBResponseBodyData) *UnBindXBResponseBody {
	s.Data = v
	return s
}

func (s *UnBindXBResponseBody) SetMessage(v string) *UnBindXBResponseBody {
	s.Message = &v
	return s
}

func (s *UnBindXBResponseBody) SetRequestId(v string) *UnBindXBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindXBResponseBody) SetSuccess(v bool) *UnBindXBResponseBody {
	s.Success = &v
	return s
}

func (s *UnBindXBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnBindXBResponseBodyData struct {
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
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindXBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnBindXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnBindXBResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UnBindXBResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnBindXBResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindXBResponseBodyData) SetCode(v string) *UnBindXBResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnBindXBResponseBodyData) SetMessage(v string) *UnBindXBResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnBindXBResponseBodyData) SetSuccess(v bool) *UnBindXBResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnBindXBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
