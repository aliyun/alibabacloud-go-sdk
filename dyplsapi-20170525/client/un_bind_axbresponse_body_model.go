// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindAXBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnBindAXBResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UnBindAXBResponseBody
	GetCode() *string
	SetData(v *UnBindAXBResponseBodyData) *UnBindAXBResponseBody
	GetData() *UnBindAXBResponseBodyData
	SetMessage(v string) *UnBindAXBResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnBindAXBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnBindAXBResponseBody
	GetSuccess() *bool
}

type UnBindAXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0000
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UnBindAXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnBindAXBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBResponseBody) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnBindAXBResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnBindAXBResponseBody) GetData() *UnBindAXBResponseBodyData {
	return s.Data
}

func (s *UnBindAXBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnBindAXBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnBindAXBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindAXBResponseBody) SetAccessDeniedDetail(v string) *UnBindAXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnBindAXBResponseBody) SetCode(v string) *UnBindAXBResponseBody {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBody) SetData(v *UnBindAXBResponseBodyData) *UnBindAXBResponseBody {
	s.Data = v
	return s
}

func (s *UnBindAXBResponseBody) SetMessage(v string) *UnBindAXBResponseBody {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBody) SetRequestId(v string) *UnBindAXBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnBindAXBResponseBody) SetSuccess(v bool) *UnBindAXBResponseBody {
	s.Success = &v
	return s
}

func (s *UnBindAXBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnBindAXBResponseBodyData struct {
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

func (s UnBindAXBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UnBindAXBResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnBindAXBResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnBindAXBResponseBodyData) SetCode(v string) *UnBindAXBResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetMessage(v string) *UnBindAXBResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnBindAXBResponseBodyData) SetSuccess(v bool) *UnBindAXBResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnBindAXBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
