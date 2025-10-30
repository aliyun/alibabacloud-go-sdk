// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraaxxShrink(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetExtraaxxShrink() *string
	SetOrderId(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelA() *string
	SetTelX(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelX() *string
	SetTelXext(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelXext() *string
	SetTs(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTs() *string
}

type BindAxnExtensionFixedLineShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002,10003
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2
	//
	// example:
	//
	// 10001,10002
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 号池ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 去掉“0” 例如：北京（10）；在平台分配X号码模式中，平台从号码池中分配该地区的X号码，避免产生呼叫长途费
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
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 接入商自有字段，最大250字符长度
	//
	// example:
	//
	// 12444
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码为被叫时主叫侧的放音编号为2
	//
	// example:
	//
	// 10001,10002
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 15500001111放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// X号码；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 分机号；平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnExtensionFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetExtraaxxShrink() *string {
	return s.ExtraaxxShrink
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelXext() *string {
	return s.TelXext
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAppId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAreacode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetBindType(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExpiration(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExtraaxxShrink(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOrderId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetRemark(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetSubts(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelA(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelX(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelXext(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTs(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
