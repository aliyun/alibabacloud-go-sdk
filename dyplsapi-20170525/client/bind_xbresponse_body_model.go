// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindXBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindXBResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindXBResponseBody
	GetCode() *string
	SetData(v *BindXBResponseBodyData) *BindXBResponseBody
	GetData() *BindXBResponseBodyData
	SetMessage(v string) *BindXBResponseBody
	GetMessage() *string
	SetSuccess(v bool) *BindXBResponseBody
	GetSuccess() *bool
}

type BindXBResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 返回状态码 0000表示成功 其他表示失败
	//
	// example:
	//
	// 0000
	Code *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindXBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s BindXBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindXBResponseBody) GoString() string {
	return s.String()
}

func (s *BindXBResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindXBResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindXBResponseBody) GetData() *BindXBResponseBodyData {
	return s.Data
}

func (s *BindXBResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindXBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindXBResponseBody) SetAccessDeniedDetail(v string) *BindXBResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindXBResponseBody) SetCode(v string) *BindXBResponseBody {
	s.Code = &v
	return s
}

func (s *BindXBResponseBody) SetData(v *BindXBResponseBodyData) *BindXBResponseBody {
	s.Data = v
	return s
}

func (s *BindXBResponseBody) SetMessage(v string) *BindXBResponseBody {
	s.Message = &v
	return s
}

func (s *BindXBResponseBody) SetSuccess(v bool) *BindXBResponseBody {
	s.Success = &v
	return s
}

func (s *BindXBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindXBResponseBodyData struct {
	// 工作号关系绑定的唯一标识
	//
	// example:
	//
	// 4353453456
	AuthId *string `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// X号码
	//
	// example:
	//
	// 18640577897
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindXBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindXBResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindXBResponseBodyData) GetAuthId() *string {
	return s.AuthId
}

func (s *BindXBResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *BindXBResponseBodyData) SetAuthId(v string) *BindXBResponseBodyData {
	s.AuthId = &v
	return s
}

func (s *BindXBResponseBodyData) SetTelX(v string) *BindXBResponseBodyData {
	s.TelX = &v
	return s
}

func (s *BindXBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
