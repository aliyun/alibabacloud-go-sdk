// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnExtensionFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnExtensionFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnExtensionFixedLineRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnExtensionFixedLineRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnExtensionFixedLineRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnExtensionFixedLineRequest
	GetExpiration() *string
	SetExtraaxx(v *BindAxnExtensionFixedLineRequestExtraaxx) *BindAxnExtensionFixedLineRequest
	GetExtraaxx() *BindAxnExtensionFixedLineRequestExtraaxx
	SetOrderId(v string) *BindAxnExtensionFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnExtensionFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnExtensionFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnExtensionFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnExtensionFixedLineRequest
	GetTelA() *string
	SetTelX(v string) *BindAxnExtensionFixedLineRequest
	GetTelX() *string
	SetTelXext(v string) *BindAxnExtensionFixedLineRequest
	GetTelXext() *string
	SetTs(v string) *BindAxnExtensionFixedLineRequest
	GetTs() *string
}

type BindAxnExtensionFixedLineRequest struct {
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
	Extraaxx *BindAxnExtensionFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
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

func (s BindAxnExtensionFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnExtensionFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnExtensionFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnExtensionFixedLineRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnExtensionFixedLineRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnExtensionFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnExtensionFixedLineRequest) GetExtraaxx() *BindAxnExtensionFixedLineRequestExtraaxx {
	return s.Extraaxx
}

func (s *BindAxnExtensionFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnExtensionFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnExtensionFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnExtensionFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnExtensionFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnExtensionFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnExtensionFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnExtensionFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnExtensionFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnExtensionFixedLineRequest) GetTelXext() *string {
	return s.TelXext
}

func (s *BindAxnExtensionFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucode(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAppId(v string) *BindAxnExtensionFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAreacode(v string) *BindAxnExtensionFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetBindType(v string) *BindAxnExtensionFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExpiration(v string) *BindAxnExtensionFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExtraaxx(v *BindAxnExtensionFixedLineRequestExtraaxx) *BindAxnExtensionFixedLineRequest {
	s.Extraaxx = v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOrderId(v string) *BindAxnExtensionFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetRemark(v string) *BindAxnExtensionFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetSubts(v string) *BindAxnExtensionFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelA(v string) *BindAxnExtensionFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelX(v string) *BindAxnExtensionFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelXext(v string) *BindAxnExtensionFixedLineRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTs(v string) *BindAxnExtensionFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) Validate() error {
	if s.Extraaxx != nil {
		if err := s.Extraaxx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnExtensionFixedLineRequestExtraaxx struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的B号码
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

func (s BindAxnExtensionFixedLineRequestExtraaxx) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequestExtraaxx) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) GetCallback() *string {
	return s.Callback
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) GetCallrecording() *string {
	return s.Callrecording
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallback(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callback = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallrecording(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callrecording = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) Validate() error {
	return dara.Validate(s)
}
