// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxbBindFixedLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAxbBindFixedLineResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAxbBindFixedLineResponseBody
	GetCode() *string
	SetData(v *QueryAxbBindFixedLineResponseBodyData) *QueryAxbBindFixedLineResponseBody
	GetData() *QueryAxbBindFixedLineResponseBodyData
	SetMessage(v string) *QueryAxbBindFixedLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAxbBindFixedLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAxbBindFixedLineResponseBody
	GetSuccess() *bool
}

type QueryAxbBindFixedLineResponseBody struct {
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
	// 绑定信息
	Data *QueryAxbBindFixedLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3179F199-C6C5-5963-85A6-21CBA2F47592
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 处理是否成功 true-成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAxbBindFixedLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAxbBindFixedLineResponseBody) GetData() *QueryAxbBindFixedLineResponseBodyData {
	return s.Data
}

func (s *QueryAxbBindFixedLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAxbBindFixedLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAxbBindFixedLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAxbBindFixedLineResponseBody) SetAccessDeniedDetail(v string) *QueryAxbBindFixedLineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetCode(v string) *QueryAxbBindFixedLineResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetData(v *QueryAxbBindFixedLineResponseBodyData) *QueryAxbBindFixedLineResponseBody {
	s.Data = v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetMessage(v string) *QueryAxbBindFixedLineResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetRequestId(v string) *QueryAxbBindFixedLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) SetSuccess(v bool) *QueryAxbBindFixedLineResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAxbBindFixedLineResponseBodyData struct {
	// 接通前放音编码，放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 接通后被叫侧放音编码
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
	// 绑定过期时间
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	Extra *QueryAxbBindFixedLineResponseBodyDataExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// 接入商自有字段，最大250字符长度
	//
	// example:
	//
	// 19394
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 绑定id
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxb
	Subid *string `json:"Subid,omitempty" xml:"Subid,omitempty"`
	// 绑定时间，格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码
	//
	// example:
	//
	// 示例值示例值
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// B号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 小号号码
	//
	// example:
	//
	// 19700002222
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetAnucode() *string {
	return s.Anucode
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetAreacode() *string {
	return s.Areacode
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetExtra() *QueryAxbBindFixedLineResponseBodyDataExtra {
	return s.Extra
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetSubid() *string {
	return s.Subid
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetSubts() *string {
	return s.Subts
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetTelA() *string {
	return s.TelA
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetTelB() *string {
	return s.TelB
}

func (s *QueryAxbBindFixedLineResponseBodyData) GetTelX() *string {
	return s.TelX
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAnucode(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Anucode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAnucodecalled(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Anucodecalled = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetAreacode(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Areacode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetExpiration(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetExtra(v *QueryAxbBindFixedLineResponseBodyDataExtra) *QueryAxbBindFixedLineResponseBodyData {
	s.Extra = v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetRemark(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetSubid(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Subid = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetSubts(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.Subts = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTAnucodeConnect(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TAnucodeConnect = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelA(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelA = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelB(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelB = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) SetTelX(v string) *QueryAxbBindFixedLineResponseBodyData {
	s.TelX = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyData) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAxbBindFixedLineResponseBodyDataExtra struct {
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s QueryAxbBindFixedLineResponseBodyDataExtra) String() string {
	return dara.Prettify(s)
}

func (s QueryAxbBindFixedLineResponseBodyDataExtra) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponseBodyDataExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *QueryAxbBindFixedLineResponseBodyDataExtra) SetCallrecording(v string) *QueryAxbBindFixedLineResponseBodyDataExtra {
	s.Callrecording = &v
	return s
}

func (s *QueryAxbBindFixedLineResponseBodyDataExtra) Validate() error {
	return dara.Validate(s)
}
