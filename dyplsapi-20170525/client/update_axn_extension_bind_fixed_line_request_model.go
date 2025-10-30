// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnExtensionBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetExpiration() *string
	SetExtraaxx(v *UpdateAxnExtensionBindFixedLineRequestExtraaxx) *UpdateAxnExtensionBindFixedLineRequest
	GetExtraaxx() *UpdateAxnExtensionBindFixedLineRequestExtraaxx
	SetOrderId(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetTelA() *string
	SetTs(v string) *UpdateAxnExtensionBindFixedLineRequest
	GetTs() *string
}

type UpdateAxnExtensionBindFixedLineRequest struct {
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
	Extraaxx *UpdateAxnExtensionBindFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
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

func (s UpdateAxnExtensionBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetExtraaxx() *UpdateAxnExtensionBindFixedLineRequestExtraaxx {
	return s.Extraaxx
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxnExtensionBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAnucode(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetAppId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetExpiration(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetExtraaxx(v *UpdateAxnExtensionBindFixedLineRequestExtraaxx) *UpdateAxnExtensionBindFixedLineRequest {
	s.Extraaxx = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetOrderId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetRemark(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetSubId(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetSubts(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTelA(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) SetTs(v string) *UpdateAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequest) Validate() error {
	if s.Extraaxx != nil {
		if err := s.Extraaxx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAxnExtensionBindFixedLineRequestExtraaxx struct {
	// 可选。 A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的B号码
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 录音控制，仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineRequestExtraaxx) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineRequestExtraaxx) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) GetCallback() *string {
	return s.Callback
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) GetCallrecording() *string {
	return s.Callrecording
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) SetCallback(v string) *UpdateAxnExtensionBindFixedLineRequestExtraaxx {
	s.Callback = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) SetCallrecording(v string) *UpdateAxnExtensionBindFixedLineRequestExtraaxx {
	s.Callrecording = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineRequestExtraaxx) Validate() error {
	return dara.Validate(s)
}
