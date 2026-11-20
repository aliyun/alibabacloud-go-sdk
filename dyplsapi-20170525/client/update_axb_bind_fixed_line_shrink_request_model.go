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
	// The caller-side audio playback code. This parameter is required for the AXB privacy number service for landlines. The audio playback code must include codes for three scenarios. Specify the codes in the order of "B->X, A->X, and other numbers->X", and separate the codes with commas.
	//
	// > "1,2,3" indicates that the audio playback code for B->X is 1, the code for A->X is 2, and the code for other numbers->X is 3.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// The callee-side audio playback code after the call is connected. The callee-side audio playback code must include codes for two scenarios. Specify the codes in the order of "A as the callee and B as the callee", and separate the codes with commas.
	//
	// > "1,2" indicates that the audio playback code when Number A answers as the callee is 1, and the code when Number B answers as the callee is 2.
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
	// - 0: The binding never expires.
	//
	// - Other values: If you specify another value, such as 60, the binding is valid for 60 seconds after Number X is allocated. After 60 seconds, the binding is automatically unbound.
	//
	// example:
	//
	// 10
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The extension parameters.
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
	// The binding ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// A100***************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The binding time. The format is yyyyMMddHHmmss. The time is in UTC+8 and uses the 24-hour format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// The caller-side audio playback code after the call is connected. The caller-side audio playback code after connection must include codes for two scenarios. Specify the codes in the order of "A as the callee and B as the callee", and separate the codes with commas.
	//
	// > "1,2" indicates that the caller-side audio playback code when Number A is the callee is 1, and the caller-side audio playback code when Number B is the callee is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// Sets Number A in the binding relationship.
	//
	//
	// > - Number A can be set to a mobile phone number or a landline number. If it is a landline number, you must add the area code. No hyphen is required between the area code and the number.
	//
	// > - You cannot update TelA and TelB at the same time.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Landline number: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// Sets Number B in the binding relationship.
	//
	//
	// > - Number B can be set to a mobile phone number or a landline number. If it is a landline number, you must add the area code. No hyphen is required between the area code and the number.
	//
	// > - You cannot update TelA and TelB at the same time.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Landline number: 0571******36
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// The business timestamp. The format is yyyyMMddHHmmssSSS. The time is in UTC+8 and uses the 24-hour format. The precision is accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
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
