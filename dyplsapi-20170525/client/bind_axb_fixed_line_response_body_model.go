// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAxbFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAxbFixedLineResponseBody
	GetCode() *string
	SetData(v *BindAxbFixedLineResponseBodyData) *BindAxbFixedLineResponseBody
	GetData() *BindAxbFixedLineResponseBodyData
	SetMessage(v string) *BindAxbFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxbFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindAxbFixedLineResponseBody
	GetSuccess() *bool
}

type BindAxbFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0-成功
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应内容
	Data *BindAxbFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 响应消息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 649E9EB5-9436-53CF-B41A-C4F0433212E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否处理成功  true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxbFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAxbFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxbFixedLineResponseBody) GetData() *BindAxbFixedLineResponseBodyData {
	return s.Data
}

func (s *BindAxbFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxbFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxbFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAxbFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxbFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetCode(v string) *BindAxbFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetData(v *BindAxbFixedLineResponseBodyData) *BindAxbFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetMessage(v string) *BindAxbFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetRequestId(v string) *BindAxbFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) SetSuccess(v bool) *BindAxbFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *BindAxbFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxbFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// 示例值
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// X号码
	//
	// example:
	//
	// 示例值
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindAxbFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *BindAxbFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *BindAxbFixedLineResponseBodyData) SetSubid(v string) *BindAxbFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxbFixedLineResponseBodyData) SetTelX(v string) *BindAxbFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
