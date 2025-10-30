// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnBindFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraShrink(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetExtraShrink() *string
	SetOrderId(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetTelA() *string
	SetTelB(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetTelB() *string
	SetTs(v string) *UpdateAxnBindFixedLineShrinkRequest
	GetTs() *string
}

type UpdateAxnBindFixedLineShrinkRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3。
	//
	// example:
	//
	// 示例值示例值示例值
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码  被叫放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫侧接听时的放音编号为1，其他号码为被叫侧接听时的放音编号为2。
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
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 额外字段
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
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
	// 1234
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧放音编码  接通后主叫侧放音编码必须包含2个场景的编码。按照“A被叫,其他号码被叫”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2”表示A号码为被叫时主叫侧的放音编号为1，其他号码被叫为被叫时主叫侧的放音编号为2
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
	// B号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnBindFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetTelB() *string {
	return s.TelB
}

func (s *UpdateAxnBindFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetExtraShrink(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTelB(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxnBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxnBindFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
