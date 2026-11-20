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
	Extra *BindAxnFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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
	// The callback control. This applies when User A makes a call through number X, meaning User A calls number X. Valid values:
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
	// - 0 (default): Call recording is disabled.
	//
	// - 1: Call recording is enabled.
	//
	// > After you enable call recording, you can play and download audio files in the call records. Audio files are billed at CNY 0.010 per minute.
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
