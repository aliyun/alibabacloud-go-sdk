// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxbBindFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraShrink(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetExtraShrink() *string
	SetOrderId(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetTelA() *string
	SetTelB(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetTelB() *string
	SetTs(v string) *UpdateAxbBindFixedLineShrinkRequest
	GetTs() *string
}

type UpdateAxbBindFixedLineShrinkRequest struct {
	// 主叫侧放音编码
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// 被叫侧放音编码
	//
	// example:
	//
	// 1,2
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// 号池ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 过期时间,单位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
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
	// 1234567890
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
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// A100X558X0000400023
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式为yyyyMMddHHmmss。时间采用北京时间，24小时制。
	//
	// This parameter is required.
	//
	// example:
	//
	// 20150920190126
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// 接通后主叫侧的放音编码
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// 真实号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13900000000
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 对端号码，telA,telB不允许同时更新
	//
	// example:
	//
	// 13005711234
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxbBindFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetTelB() *string {
	return s.TelB
}

func (s *UpdateAxbBindFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetExtraShrink(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTelB(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxbBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxbBindFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
