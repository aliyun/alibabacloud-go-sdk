// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAxnExtensionFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *BindAxnExtensionFixedLineResponseBody
	GetCode() *string
	SetData(v *BindAxnExtensionFixedLineResponseBodyData) *BindAxnExtensionFixedLineResponseBody
	GetData() *BindAxnExtensionFixedLineResponseBodyData
	SetMessage(v string) *BindAxnExtensionFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindAxnExtensionFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v string) *BindAxnExtensionFixedLineResponseBody
	GetSuccess() *string
}

type BindAxnExtensionFixedLineResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码 0：成功，其它失败，具体见文档
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应信息
	Data *BindAxnExtensionFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AE2D6997-643A-59CB-9B3C-918572E5CEAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAxnExtensionFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAxnExtensionFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindAxnExtensionFixedLineResponseBody) GetData() *BindAxnExtensionFixedLineResponseBodyData {
	return s.Data
}

func (s *BindAxnExtensionFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindAxnExtensionFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAxnExtensionFixedLineResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *BindAxnExtensionFixedLineResponseBody) SetAccessDeniedDetail(v string) *BindAxnExtensionFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetCode(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetData(v *BindAxnExtensionFixedLineResponseBodyData) *BindAxnExtensionFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetMessage(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetRequestId(v string) *BindAxnExtensionFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) SetSuccess(v string) *BindAxnExtensionFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnExtensionFixedLineResponseBodyData struct {
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号，只有4位
	//
	// example:
	//
	// 1001
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
}

func (s BindAxnExtensionFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *BindAxnExtensionFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnExtensionFixedLineResponseBodyData) GetTelXext() *string {
	return s.TelXext
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetSubid(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetTelX(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBodyData) SetTelXext(v string) *BindAxnExtensionFixedLineResponseBodyData {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
