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
	Extra *UpdateAxnBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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
	// The callback control. This applies when A calls through X, meaning A is the caller and X is the callee. Valid values:
	//
	// - 0 (default): Outbound calls are not allowed.
	//
	// - 1: Connect to the most recent number N.
	//
	// - 2: Call back the fixed number TelB.
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The call recording control. Valid values:
	//
	// - 0 (default): Do not record.
	//
	// - 1: Record.
	//
	// > After you enable call recording, you can play and download audio files in the call records. Audio files are billed at CNY 0.010 per minute.
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
