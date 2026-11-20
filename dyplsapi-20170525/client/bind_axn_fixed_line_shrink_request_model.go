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
	// The audio prompt codes for the caller side. The value must include codes for three scenarios. Specify the codes in the order of B->X, A->X, and other numbers->X, separated by commas.
	//
	// > For example, 1,2,3 indicates that the audio prompt code for B->X is 1, the code for A->X is 2, and the code for other numbers->X is 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// The audio prompt codes for the callee side. The value must include codes for two scenarios. Specify the codes in the order of A as the callee and other numbers as the callee, separated by commas.
	//
	// > For example, 1,2 indicates that the audio prompt code when number A answers as the callee is 1, and the code when other numbers answer as the callee is 2.
	//
	// example:
	//
	// 1,2
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// The ID of the number pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The area code. You can specify the location of number X (specified city -> provincial capital -> national routing). If you do not specify the location, a number is randomly assigned nationwide.
	//
	// example:
	//
	// 0571
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// The binding type. Valid values:
	//
	// - mode101: The customer specifies number X.
	//
	// - mode102: The platform assigns number X.
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// The expiration time. Unit: seconds. Valid values:
	//
	// - 0: The binding never expires.
	//
	// - Other values: If you specify another value, such as 60, the binding is valid for 60 seconds after number X is successfully assigned. After 60 seconds, the binding is automatically unbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The extended parameters.
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The unique identifier of the message request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The external business extension field. This parameter is returned in the call record receipt message. The value cannot exceed 100 characters in length.
	//
	// example:
	//
	// 19394
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding time. The format is yyyyMMddHHmmss. The time is in UTC+8 and uses the 24-hour clock.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// The audio prompt codes for the caller side after the call is connected. The value must include codes for two scenarios. Specify the codes in the order of A as the callee and other numbers as the callee, separated by commas.
	//
	// > For example, 1,2 indicates that the audio prompt code for the caller side when number A is the callee is 1, and the code for the caller side when other numbers are the callee is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// Number A of the fixed-line AXN privacy number.
	//
	//
	// > - Number A can be a mobile phone number or a fixed-line number. A fixed-line number must include the area code, and no hyphen is required between the area code and the number.
	//
	// > - If your business requires modifying a bound number, you can call the [UpdateAxnBindFixedLine](https://api.aliyun.com/document/Dyplsapi/2017-05-25/UpdateAxnBindFixedLine) operation to modify it.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Fixed-line number: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// Number N of the fixed-line AXN privacy number.
	//
	// example:
	//
	// 155****2222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// Number X of the fixed-line AXN privacy number. Number X is the privacy number that you have purchased in the console.
	//
	// > This parameter is required in the mode where the customer specifies number X (that is, the BindType parameter is set to mode101).
	//
	// example:
	//
	// 0571******23
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// The business timestamp. The format is yyyyMMddHHmmssSSS. The time is in UTC+8 and uses the 24-hour clock, accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
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
