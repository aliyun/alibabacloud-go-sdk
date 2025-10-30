// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnExtensionBindFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraaxxShrink(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetExtraaxxShrink() *string
	SetOrderId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTelA() *string
	SetTs(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTs() *string
}

type UpdateAxnExtensionBindFixedLineShrinkRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  AXN分机号业务的放音编码,B->X和其他号码->X的编码一致  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
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
	// 号池ID
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// AXx的扩展参数项
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
	// 消息请求唯一标识。
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
	// 1233
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// 可参考绑定响应
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
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
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetExtraaxxShrink() *string {
	return s.ExtraaxxShrink
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExtraaxxShrink(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
