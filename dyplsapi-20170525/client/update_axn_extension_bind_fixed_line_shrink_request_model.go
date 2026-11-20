// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnExtensionBindFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetAppId() *string
	SetExpiration(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraaxxShrink(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetExtraaxxShrink() *string
	SetOrderId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetSubId() *string
	SetSubts(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTelA() *string
	SetTs(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest
	GetTs() *string
}

type UpdateAxnExtensionBindFixedLineShrinkRequest struct {
	// The announcement code for the caller side. This parameter is required for the AXB service. It must contain codes for three scenarios. Specify the codes in the order of "B->X, A->X, and other numbers->X", separated by commas.
	//
	// > "1,2,3" indicates that the announcement code for B->X is 1, the announcement code for A->X is 2, and the announcement code for other numbers->X is 3.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// The announcement code for the callee side after the call is connected. It must contain codes for two scenarios. Specify the codes in the order of "A as callee and B as callee", separated by commas.
	//
	// > "1,2" indicates that the announcement code when number A answers as the callee is 1, and the announcement code when number B answers as the callee is 2.
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
	// - Other values: If you specify another value, such as 60, the binding is valid for 60 seconds after number X is allocated. After 60 seconds, the binding is automatically unbound.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The extension parameters.
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
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
	// GHX0*************************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The binding time. Format: yyyyMMddHHmmss. The time is in UTC+8 and uses the 24-hour format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// The announcement code for the caller side after the call is connected. It must contain codes for two scenarios. Specify the codes in the order of "A as callee and B as callee", separated by commas.
	//
	// > "1,2" indicates that the announcement code for the caller side when number A is the callee is 1, and the announcement code for the caller side when number B is the callee is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// The number A of the landline AXN extension.
	//
	// > Number A can be set to a mobile phone number or a landline number. For a landline number, the area code must be included, and no hyphen is required between the area code and the number.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Landline number: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// The business timestamp. Format: yyyyMMddHHmmssSSS. The time is in UTC+8 and uses the 24-hour format, accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetExtraaxxShrink() *string {
	return s.ExtraaxxShrink
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetSubId() *string {
	return s.SubId
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucode(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAnucodecalled(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetAppId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExpiration(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetExtraaxxShrink(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOrderId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetRemark(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetResourceOwnerId(v int64) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubId(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.SubId = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetSubts(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTAnucodeConnect(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTelA(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) SetTs(v string) *UpdateAxnExtensionBindFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
