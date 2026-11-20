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
	// The prompt code for the caller side. This parameter is required for AXB services. The codes must include three scenarios. Enter the codes in the order of B->X, A->X, and other numbers->X, separated by commas.
	//
	// > 1,2,3 indicates that the prompt code for B->X is 1, the prompt code for A->X is 2, and the prompt code for other numbers->X is 3.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// The prompt code for the callee side after the call is connected. The codes must include two scenarios. Enter the codes in the order of A as callee and B as callee, separated by commas.
	//
	// > 1,2 indicates that the prompt code when A is the callee is 1, and the prompt code when B is the callee is 2.
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
	// The expiration time. Unit: seconds. Valid values:
	//
	// - 0: Never expires.
	//
	// - Other values: If you specify another value, such as 60, the binding is valid for 60 seconds after the number X is allocated. After 60 seconds, the binding is automatically unbound.
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
	// The external business extension field. This parameter is returned in the call record receipt message. The maximum length is 100 characters.
	//
	// example:
	//
	// 19394
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0*************************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The binding time. The format is yyyyMMddHHmmss. The time is in UTC+8 and uses a 24-hour clock.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// The prompt code for the caller side after the call is connected. The codes must include two scenarios. Enter the codes in the order of A as callee and B as callee, separated by commas.
	//
	// > 1,2 indicates that the prompt code for the caller side when A is the callee is 1, and the prompt code for the caller side when B is the callee is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// The number A of the landline AXN privacy number.
	//
	// > You can set number A to a mobile phone number or a landline number. If you use a landline number, you must include the area code. Do not add a hyphen between the area code and the phone number.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Landline number: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// The number N of the landline AXN privacy number.
	//
	// example:
	//
	// 155****2222
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// The business timestamp. The format is yyyyMMddHHmmssSSS. The time is in UTC+8 and uses a 24-hour clock, accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
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
