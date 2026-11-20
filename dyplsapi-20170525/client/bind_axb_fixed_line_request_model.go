// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxbFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxbFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxbFixedLineRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxbFixedLineRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxbFixedLineRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxbFixedLineRequest
	GetExpiration() *string
	SetExtra(v *BindAxbFixedLineRequestExtra) *BindAxbFixedLineRequest
	GetExtra() *BindAxbFixedLineRequestExtra
	SetOrderId(v string) *BindAxbFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxbFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxbFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxbFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxbFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxbFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxbFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxbFixedLineRequest
	GetTelA() *string
	SetTelB(v string) *BindAxbFixedLineRequest
	GetTelB() *string
	SetTelX(v string) *BindAxbFixedLineRequest
	GetTelX() *string
	SetTs(v string) *BindAxbFixedLineRequest
	GetTs() *string
}

type BindAxbFixedLineRequest struct {
	// The caller-side announcement code. This parameter is required for the fixed-line AXB privacy number service. The code must include announcements for three scenarios. Specify the codes in the order of "B->X, A->X, other numbers->X", separated by commas.
	//
	// > For example, "1,2,3" indicates that the announcement code for B->X is 1, the announcement code for A->X is 2, and the announcement code for other numbers->X is 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// The callee-side announcement code after the call is connected. The code must include announcements for two scenarios. Specify the codes in the order of "A as callee, B as callee", separated by commas.
	//
	// > For example, "1,2" indicates that the announcement code when Number A answers as the callee is 1, and the announcement code when Number B answers as the callee is 2.
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
	// The area code. You can specify the location of Number X (specified city -> provincial capital -> national routing). If you do not specify a location, a number is randomly assigned nationwide.
	//
	// example:
	//
	// 0571
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// The binding type. Valid values:
	//
	// - mode101: The customer specifies Number X.
	//
	// - mode102: The platform allocates Number X.
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
	// - Other values: If you specify another value, such as 60, the binding is valid for 60 seconds after Number X is successfully allocated. After 60 seconds, the binding is automatically unbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The extended parameters.
	Extra *BindAxbFixedLineRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
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
	// The caller-side announcement code after the call is connected. The code must include announcements for two scenarios. Specify the codes in the order of "A as callee, B as callee", separated by commas.
	//
	// > For example, "1,2" indicates that the caller-side announcement code when Number A is the callee is 1, and the caller-side announcement code when Number B is the callee is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// Number A of the fixed-line AXB privacy number.
	//
	// > - Number A can be a mobile phone number or a fixed-line number. For a fixed-line number, add the area code. No hyphen is required between the area code and the fixed-line number.
	//
	// > - If your business requires modifying a bound number, you can call the [UpdateAxbBindFixedLine](https://api.aliyun.com/document/Dyplsapi/2017-05-25/UpdateAxbBindFixedLine) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Fixed-line number: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// Number B of the fixed-line AXB privacy number.
	//
	// > - Number B can be a mobile phone number or a fixed-line number. For a fixed-line number, add the area code. No hyphen is required between the area code and the fixed-line number.
	//
	// > - If your business requires modifying a bound number, you can call the [UpdateAxbBindFixedLine](https://api.aliyun.com/document/Dyplsapi/2017-05-25/UpdateAxbBindFixedLine) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mobile phone number: 182****2371
	//
	// Fixed-line number: 0571******36
	TelB *string `json:"TelB,omitempty" xml:"TelB,omitempty"`
	// Number X of the fixed-line AXB privacy number. Number X is a privacy number that you have [purchased](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
	//
	// > This parameter is required when the customer specifies Number X (that is, the BindType parameter is set to mode101).
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

func (s BindAxbFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxbFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxbFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxbFixedLineRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxbFixedLineRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxbFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxbFixedLineRequest) GetExtra() *BindAxbFixedLineRequestExtra {
	return s.Extra
}

func (s *BindAxbFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxbFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxbFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxbFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxbFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxbFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxbFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxbFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxbFixedLineRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxbFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxbFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxbFixedLineRequest) SetAnucode(v string) *BindAxbFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAnucodecalled(v string) *BindAxbFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAppId(v string) *BindAxbFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetAreacode(v string) *BindAxbFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetBindType(v string) *BindAxbFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExpiration(v string) *BindAxbFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetExtra(v *BindAxbFixedLineRequestExtra) *BindAxbFixedLineRequest {
	s.Extra = v
	return s
}

func (s *BindAxbFixedLineRequest) SetOrderId(v string) *BindAxbFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetOwnerId(v int64) *BindAxbFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetRemark(v string) *BindAxbFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetSubts(v string) *BindAxbFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelA(v string) *BindAxbFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelB(v string) *BindAxbFixedLineRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTelX(v string) *BindAxbFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineRequest) SetTs(v string) *BindAxbFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *BindAxbFixedLineRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxbFixedLineRequestExtra struct {
	// The call recording control. Valid values:
	//
	// - 0 (default): Do not record the call.
	//
	// - 1: Record the call.
	//
	// > After you enable call recording, you can play and download voice files in the call records. Recording files are billed at CNY 0.010 per minute.
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxbFixedLineRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineRequestExtra) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineRequestExtra) GetCallrecording() *string {
	return s.Callrecording
}

func (s *BindAxbFixedLineRequestExtra) SetCallrecording(v string) *BindAxbFixedLineRequestExtra {
	s.Callrecording = &v
	return s
}

func (s *BindAxbFixedLineRequestExtra) Validate() error {
	return dara.Validate(s)
}
