// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxnBindFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxnBindFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxnBindFixedLineRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxnBindFixedLineRequest
	GetExpiration() *string
	SetExtra(v *UpdateAxnBindFixedLineRequestExtra) *UpdateAxnBindFixedLineRequest
	GetExtra() *UpdateAxnBindFixedLineRequestExtra
	SetOrderId(v string) *UpdateAxnBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxnBindFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxnBindFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxnBindFixedLineRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxnBindFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxnBindFixedLineRequest
	GetTelA() *string
	SetTelB(v string) *UpdateAxnBindFixedLineRequest
	GetTelB() *string
	SetTs(v string) *UpdateAxnBindFixedLineRequest
	GetTs() *string
}

type UpdateAxnBindFixedLineRequest struct {
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
	Extra *UpdateAxnBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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

func (s UpdateAxnBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxnBindFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxnBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxnBindFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxnBindFixedLineRequest) GetExtra() *UpdateAxnBindFixedLineRequestExtra {
	return s.Extra
}

func (s *UpdateAxnBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxnBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxnBindFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxnBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxnBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxnBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxnBindFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxnBindFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxnBindFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxnBindFixedLineRequest) GetTelB() *string {
	return s.TelB
}

func (s *UpdateAxnBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxnBindFixedLineRequest) SetAnucode(v string) *UpdateAxnBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxnBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetAppId(v string) *UpdateAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetExpiration(v string) *UpdateAxnBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetExtra(v *UpdateAxnBindFixedLineRequestExtra) *UpdateAxnBindFixedLineRequest {
	s.Extra = v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetOrderId(v string) *UpdateAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetRemark(v string) *UpdateAxnBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetSubId(v string) *UpdateAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetSubts(v string) *UpdateAxnBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxnBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTelA(v string) *UpdateAxnBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTelB(v string) *UpdateAxnBindFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) SetTs(v string) *UpdateAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxnBindFixedLineRequestExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 1
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxnBindFixedLineRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnBindFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *UpdateAxnBindFixedLineRequestExtra) GetCallback() *string {
	return s.Callback
}

func (s *UpdateAxnBindFixedLineRequestExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *UpdateAxnBindFixedLineRequestExtra) SetCallback(v string) *UpdateAxnBindFixedLineRequestExtra {
	s.Callback = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequestExtra) SetCallrecording(v string) *UpdateAxnBindFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

func (s *UpdateAxnBindFixedLineRequestExtra) Validate() error {
	return dara.Validate(s)
}
