// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAxnFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAxnFixedLineResponseBody
	GetCode() *string
	SetData(v *BindAxnFixedLineResponseBodyData) *BindAxnFixedLineResponseBody
	GetData() *BindAxnFixedLineResponseBodyData
	SetMessage(v string) *BindAxnFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxnFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindAxnFixedLineResponseBody
	GetSuccess() *bool
}

type BindAxnFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindAxnFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4D690962-08CE-5D38-A65A-AB247D7DF7A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxnFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAxnFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxnFixedLineResponseBody) GetData() *BindAxnFixedLineResponseBodyData {
	return s.Data
}

func (s *BindAxnFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxnFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxnFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindAxnFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxnFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetCode(v string) *BindAxnFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetData(v *BindAxnFixedLineResponseBodyData) *BindAxnFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetMessage(v string) *BindAxnFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetRequestId(v string) *BindAxnFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) SetSuccess(v bool) *BindAxnFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *BindAxnFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s BindAxnFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *BindAxnFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnFixedLineResponseBodyData) SetSubid(v string) *BindAxnFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxnFixedLineResponseBodyData) SetTelX(v string) *BindAxnFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *BindAxnFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
