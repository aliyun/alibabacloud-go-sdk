// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnFixedLineRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnFixedLineRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnFixedLineRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnFixedLineRequest
	GetExpiration() *string
	SetExtra(v *BindAxnFixedLineRequestExtra) *BindAxnFixedLineRequest
	GetExtra() *BindAxnFixedLineRequestExtra
	SetOrderId(v string) *BindAxnFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnFixedLineRequest
	GetTelA() *string
	SetTelB(v string) *BindAxnFixedLineRequest
	GetTelB() *string
	SetTelX(v string) *BindAxnFixedLineRequest
	GetTelX() *string
	SetTs(v string) *BindAxnFixedLineRequest
	GetTs() *string
}

type BindAxnFixedLineRequest struct {
	// 放音编码必须包含3个场景的编码。按照“B->X,A->X,其他号码->X”的顺序填写编码，编码之间以逗号分隔。  比如：“1,2,3”表示B->X放音编号为1，A->X放音编号为2， 其他号码->X放音编号为3
	//
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
	// 位：秒，必须为数字 0：不会自动解绑 非0：自动解绑周期
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// 扩展参数
	Extra *BindAxnFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// N号码
	//
	// example:
	//
	// 15500002222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// X号码； 平台分配号码模式下，该参数可不带，系统忽略该参数
	//
	// example:
	//
	// 0571409312
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnFixedLineRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnFixedLineRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnFixedLineRequest) GetExtra() *BindAxnFixedLineRequestExtra {
	return s.Extra
}

func (s *BindAxnFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnFixedLineRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxnFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnFixedLineRequest) SetAnucode(v string) *BindAxnFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAnucodecalled(v string) *BindAxnFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAppId(v string) *BindAxnFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetAreacode(v string) *BindAxnFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetBindType(v string) *BindAxnFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetExpiration(v string) *BindAxnFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetExtra(v *BindAxnFixedLineRequestExtra) *BindAxnFixedLineRequest {
	s.Extra = v
	return s
}

func (s *BindAxnFixedLineRequest) SetOrderId(v string) *BindAxnFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetOwnerId(v int64) *BindAxnFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetRemark(v string) *BindAxnFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxnFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetResourceOwnerId(v int64) *BindAxnFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetSubts(v string) *BindAxnFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTAnucodeConnect(v string) *BindAxnFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelA(v string) *BindAxnFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelB(v string) *BindAxnFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTelX(v string) *BindAxnFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnFixedLineRequest) SetTs(v string) *BindAxnFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnFixedLineRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnFixedLineRequestExtra struct {
	// A通过X呼叫，即A主叫X，仅下列值有效。默认是0。 0：不能外呼 1：接续最近的N号码 2：回拨固定号码：telB
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 仅下列值有效。默认是0（不开通录音功能）。 0：不录音 1：录音
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxnFixedLineRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineRequestExtra) GetCallback() *string {
	return s.Callback
}

func (s *BindAxnFixedLineRequestExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *BindAxnFixedLineRequestExtra) SetCallback(v string) *BindAxnFixedLineRequestExtra {
	s.Callback = &v
	return s
}

func (s *BindAxnFixedLineRequestExtra) SetCallrecording(v string) *BindAxnFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

func (s *BindAxnFixedLineRequestExtra) Validate() error {
	return dara.Validate(s)
}
