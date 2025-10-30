// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxbFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxbFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxbFixedLineShrinkRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxbFixedLineShrinkRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxbFixedLineShrinkRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxbFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraShrink(v string) *BindAxbFixedLineShrinkRequest
	GetExtraShrink() *string
	SetOrderId(v string) *BindAxbFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxbFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxbFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxbFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxbFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxbFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxbFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxbFixedLineShrinkRequest
	GetTelA() *string
	SetTelB(v string) *BindAxbFixedLineShrinkRequest
	GetTelB() *string
	SetTelX(v string) *BindAxbFixedLineShrinkRequest
	GetTelX() *string
	SetTs(v string) *BindAxbFixedLineShrinkRequest
	GetTs() *string
}

type BindAxbFixedLineShrinkRequest struct {
	// 主叫侧放音编码，AXB业务时必须设置。  放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 接通后被叫侧放音编码,接通后被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，B号码为被叫侧接听时的放音编号为2。
	//
	// example:
	//
	// 示例值示例值
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 号池ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 区号,去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费。
	//
	// example:
	//
	// 10
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// 绑定类型，值如下： mode101：客户携带X号码 mode102：平台分配X号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// 消息请求唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// 15433678436
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大250字符长度
	//
	// example:
	//
	// remark
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定时间，格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码,接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,B被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，B号码为被叫时主叫侧的放音编号为2。
	//
	// example:
	//
	// 示例值示例值示例值
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18456713271
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 18971362645
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数  格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// example:
	//
	// 19767562345
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxbFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxbFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxbFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxbFixedLineShrinkRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxbFixedLineShrinkRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxbFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxbFixedLineShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *BindAxbFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxbFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxbFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxbFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxbFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxbFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxbFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxbFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxbFixedLineShrinkRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxbFixedLineShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxbFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucode(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAppId(v string) *BindAxbFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAreacode(v string) *BindAxbFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetBindType(v string) *BindAxbFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExpiration(v string) *BindAxbFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExtraShrink(v string) *BindAxbFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOrderId(v string) *BindAxbFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetRemark(v string) *BindAxbFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetSubts(v string) *BindAxbFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelA(v string) *BindAxbFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelB(v string) *BindAxbFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelX(v string) *BindAxbFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTs(v string) *BindAxbFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
