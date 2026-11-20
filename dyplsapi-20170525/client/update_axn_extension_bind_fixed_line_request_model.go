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
	Extraaxx *UpdateAxnExtensionBindFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
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
	// The callback control. Number A calls through number X, which means number A is the caller and number X is the callee. Valid values:
	//
	// - 0 (default): Outbound calls are not allowed.
	//
	// - 1: Connects to the most recent number N.
	//
	// - 2: Calls back the fixed number TelB.
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The call recording control. Valid values:
	//
	// - 0 (default): Does not record calls.
	//
	// - 1: Records calls.
	//
	// > After the call recording feature is enabled, you can play and download voice files in the call records. Recording files are billed at CNY 0.010 per minute based on the call duration.
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
