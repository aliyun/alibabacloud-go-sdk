// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAxnBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAxnBindFixedLineResponseBody
	GetCode() *string
	SetData(v []*QueryAxnBindFixedLineResponseBodyData) *QueryAxnBindFixedLineResponseBody
	GetData() []*QueryAxnBindFixedLineResponseBodyData
	SetMessage(v string) *QueryAxnBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAxnBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAxnBindFixedLineResponseBody
	GetSuccess() *bool
}

type QueryAxnBindFixedLineResponseBody struct {
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
	// 绑定对象
	Data []*QueryAxnBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E8B9C3ED-D9BD-5E27-9588-6D84D3070160
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAxnBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAxnBindFixedLineResponseBody) GetData() []*QueryAxnBindFixedLineResponseBodyData {
	return s.Data
}

func (s *QueryAxnBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAxnBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAxnBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAxnBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxnBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetCode(v string) *QueryAxnBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetData(v []*QueryAxnBindFixedLineResponseBodyData) *QueryAxnBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetMessage(v string) *QueryAxnBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetRequestId(v string) *QueryAxnBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxnBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBody) Validate() error {
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

type QueryAxnBindFixedLineResponseBodyData struct {
	// 接通前放音啊编码
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 隐私号码区号
	//
	// example:
	//
	// 010
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extra *QueryAxnBindFixedLineResponseBodyDataExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 接入商自有字段，最大250字符长度
	//
	// example:
	//
	// 12444
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
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
	// N号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 057112345678
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetAnucode() *string {
	return s.Anucode
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetAreacode() *string {
	return s.Areacode
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetExtra() *QueryAxnBindFixedLineResponseBodyDataExtra {
	return s.Extra
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetSubts() *string {
	return s.Subts
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetTelA() *string {
	return s.TelA
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetTelB() *string {
	return s.TelB
}

func (s *QueryAxnBindFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetExtra(v *QueryAxnBindFixedLineResponseBodyDataExtra) *QueryAxnBindFixedLineResponseBodyData {
	s.Extra = v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelB(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelB = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxnBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyData) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAxnBindFixedLineResponseBodyDataExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 1
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxnBindFixedLineResponseBodyDataExtra) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnBindFixedLineResponseBodyDataExtra) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) GetCallback() *string {
	return s.Callback
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) SetCallback(v string) *QueryAxnBindFixedLineResponseBodyDataExtra {
	s.Callback = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) SetCallrecording(v string) *QueryAxnBindFixedLineResponseBodyDataExtra {
	s.Callrecording = &v
	return s
}

func (s *QueryAxnBindFixedLineResponseBodyDataExtra) Validate() error {
	return dara.Validate(s)
}
