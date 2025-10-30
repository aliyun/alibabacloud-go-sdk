// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnExtensionBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAxnExtensionBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAxnExtensionBindFixedLineResponseBody
	GetCode() *string
	SetData(v []*QueryAxnExtensionBindFixedLineResponseBodyData) *QueryAxnExtensionBindFixedLineResponseBody
	GetData() []*QueryAxnExtensionBindFixedLineResponseBodyData
	SetMessage(v string) *QueryAxnExtensionBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAxnExtensionBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAxnExtensionBindFixedLineResponseBody
	GetSuccess() *bool
}

type QueryAxnExtensionBindFixedLineResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 响应码
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 查询绑定对象集合，具体对象字段见绑定请求
	Data []*QueryAxnExtensionBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 74EFA0E8-CFCA-54D9-BFE5-904F9FA88DBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetData() []*QueryAxnExtensionBindFixedLineResponseBodyData {
	return s.Data
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetCode(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetData(v []*QueryAxnExtensionBindFixedLineResponseBodyData) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetMessage(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetRequestId(v string) *QueryAxnExtensionBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxnExtensionBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAxnExtensionBindFixedLineResponseBodyData struct {
	// 放音编码
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 隐私号码区号
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extraaxx *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
	// 接入商自有字段，最大250字符长度
	//
	// example:
	//
	// 12444
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定ID
	//
	// example:
	//
	// 可参考绑定响应
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 绑定时间
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 小号号码
	//
	// example:
	//
	// 19700002222
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetAnucode() *string {
	return s.Anucode
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetAreacode() *string {
	return s.Areacode
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetExtraaxx() *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx {
	return s.Extraaxx
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetSubts() *string {
	return s.Subts
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetTelA() *string {
	return s.TelA
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) GetTelXext() *string {
	return s.TelXext
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetExtraaxx(v *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Extraaxx = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) SetTelXext(v string) *QueryAxnExtensionBindFixedLineResponseBodyData {
	s.TelXext = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyData) Validate() error {
	if s.Extraaxx != nil {
		if err := s.Extraaxx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx struct {
	// 回拨控制
	//
	// example:
	//
	// 示例值示例值
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制
	//
	// example:
	//
	// 示例值示例值
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) GetCallback() *string {
	return s.Callback
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) GetCallrecording() *string {
	return s.Callrecording
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) SetCallback(v string) *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx {
	s.Callback = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) SetCallrecording(v string) *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx {
	s.Callrecording = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponseBodyDataExtraaxx) Validate() error {
	return dara.Validate(s)
}
