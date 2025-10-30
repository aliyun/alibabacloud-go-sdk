// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxbBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxbBindFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxbBindFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxbBindFixedLineRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxbBindFixedLineRequest
	GetExpiration() *string
	SetExtra(v *UpdateAxbBindFixedLineRequestExtra) *UpdateAxbBindFixedLineRequest
	GetExtra() *UpdateAxbBindFixedLineRequestExtra
	SetOrderId(v string) *UpdateAxbBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxbBindFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxbBindFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxbBindFixedLineRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxbBindFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxbBindFixedLineRequest
	GetTelA() *string
	SetTelB(v string) *UpdateAxbBindFixedLineRequest
	GetTelB() *string
	SetTs(v string) *UpdateAxbBindFixedLineRequest
	GetTs() *string
}

type UpdateAxbBindFixedLineRequest struct {
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
	Extra *UpdateAxbBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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

func (s UpdateAxbBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxbBindFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxbBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxbBindFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxbBindFixedLineRequest) GetExtra() *UpdateAxbBindFixedLineRequestExtra {
	return s.Extra
}

func (s *UpdateAxbBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxbBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxbBindFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxbBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxbBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxbBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxbBindFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxbBindFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxbBindFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxbBindFixedLineRequest) GetTelB() *string {
	return s.TelB
}

func (s *UpdateAxbBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxbBindFixedLineRequest) SetAnucode(v string) *UpdateAxbBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxbBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetAppId(v string) *UpdateAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetExpiration(v string) *UpdateAxbBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetExtra(v *UpdateAxbBindFixedLineRequestExtra) *UpdateAxbBindFixedLineRequest {
	s.Extra = v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetOrderId(v string) *UpdateAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetRemark(v string) *UpdateAxbBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetSubId(v string) *UpdateAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetSubts(v string) *UpdateAxbBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxbBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTelA(v string) *UpdateAxbBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTelB(v string) *UpdateAxbBindFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) SetTs(v string) *UpdateAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxbBindFixedLineRequestExtra struct {
	// 录音控制， 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxbBindFixedLineRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxbBindFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *UpdateAxbBindFixedLineRequestExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *UpdateAxbBindFixedLineRequestExtra) SetCallrecording(v string) *UpdateAxbBindFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

func (s *UpdateAxbBindFixedLineRequestExtra) Validate() error {
	return dara.Validate(s)
}
