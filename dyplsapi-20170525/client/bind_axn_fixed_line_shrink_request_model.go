// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnFixedLineShrinkRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnFixedLineShrinkRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnFixedLineShrinkRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraShrink(v string) *BindAxnFixedLineShrinkRequest
	GetExtraShrink() *string
	SetOrderId(v string) *BindAxnFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnFixedLineShrinkRequest
	GetTelA() *string
	SetTelB(v string) *BindAxnFixedLineShrinkRequest
	GetTelB() *string
	SetTelX(v string) *BindAxnFixedLineShrinkRequest
	GetTelX() *string
	SetTs(v string) *BindAxnFixedLineShrinkRequest
	GetTs() *string
}

type BindAxnFixedLineShrinkRequest struct {
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

func (s BindAxnFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxnFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnFixedLineShrinkRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnFixedLineShrinkRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnFixedLineShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *BindAxnFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnFixedLineShrinkRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxnFixedLineShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnFixedLineShrinkRequest) SetAnucode(v string) *BindAxnFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxnFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAppId(v string) *BindAxnFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetAreacode(v string) *BindAxnFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetBindType(v string) *BindAxnFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetExpiration(v string) *BindAxnFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetExtraShrink(v string) *BindAxnFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetOrderId(v string) *BindAxnFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxnFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetRemark(v string) *BindAxnFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxnFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxnFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetSubts(v string) *BindAxnFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxnFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelA(v string) *BindAxnFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelB(v string) *BindAxnFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTelX(v string) *BindAxnFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) SetTs(v string) *BindAxnFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
