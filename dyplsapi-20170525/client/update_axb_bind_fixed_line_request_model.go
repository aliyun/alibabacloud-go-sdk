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
	Extra *UpdateAxbBindFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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
	// The call recording control. Valid values:
	//
	// - 0 (default): Call recording is disabled.
	//
	// - 1: Call recording is enabled.
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
