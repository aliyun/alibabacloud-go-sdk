// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxbFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxbFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxbFixedLineRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxbFixedLineRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxbFixedLineRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxbFixedLineRequest
	GetExpiration() *string
	SetExtra(v *BindAxbFixedLineRequestExtra) *BindAxbFixedLineRequest
	GetExtra() *BindAxbFixedLineRequestExtra
	SetOrderId(v string) *BindAxbFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxbFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxbFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxbFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxbFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxbFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxbFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxbFixedLineRequest
	GetTelA() *string
	SetTelB(v string) *BindAxbFixedLineRequest
	GetTelB() *string
	SetTelX(v string) *BindAxbFixedLineRequest
	GetTelX() *string
	SetTs(v string) *BindAxbFixedLineRequest
	GetTs() *string
}

type BindAxbFixedLineRequest struct {
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
	Extra *BindAxbFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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

func (s BindAxbFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxbFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxbFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxbFixedLineRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxbFixedLineRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxbFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxbFixedLineRequest) GetExtra() *BindAxbFixedLineRequestExtra {
	return s.Extra
}

func (s *BindAxbFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxbFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxbFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxbFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxbFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxbFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxbFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxbFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxbFixedLineRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxbFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxbFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxbFixedLineRequest) SetAnucode(v string) *BindAxbFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAnucodecalled(v string) *BindAxbFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAppId(v string) *BindAxbFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAreacode(v string) *BindAxbFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetBindType(v string) *BindAxbFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExpiration(v string) *BindAxbFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExtra(v *BindAxbFixedLineRequestExtra) *BindAxbFixedLineRequest {
	s.Extra = v
	return s
}

func (s *BindAxbFixedLineRequest) SetOrderId(v string) *BindAxbFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetOwnerId(v int64) *BindAxbFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetRemark(v string) *BindAxbFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetSubts(v string) *BindAxbFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelA(v string) *BindAxbFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelB(v string) *BindAxbFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelX(v string) *BindAxbFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTs(v string) *BindAxbFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *BindAxbFixedLineRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxbFixedLineRequestExtra struct {
	// 录音控制，默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 示例值示例值
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxbFixedLineRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequestExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *BindAxbFixedLineRequestExtra) SetCallrecording(v string) *BindAxbFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

func (s *BindAxbFixedLineRequestExtra) Validate() error {
	return dara.Validate(s)
}
