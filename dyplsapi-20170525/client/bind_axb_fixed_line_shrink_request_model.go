// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxbFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxbFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxbFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxbFixedLineShrinkRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxbFixedLineShrinkRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxbFixedLineShrinkRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxbFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraShrink(v string) *BindAxbFixedLineShrinkRequest
	GetExtraShrink() *string
	SetOrderId(v string) *BindAxbFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxbFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxbFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxbFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxbFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxbFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxbFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxbFixedLineShrinkRequest
	GetTelA() *string
	SetTelB(v string) *BindAxbFixedLineShrinkRequest
	GetTelB() *string
	SetTelX(v string) *BindAxbFixedLineShrinkRequest
	GetTelX() *string
	SetTs(v string) *BindAxbFixedLineShrinkRequest
	GetTs() *string
}

type BindAxbFixedLineShrinkRequest struct {
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

func (s BindAxbFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxbFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxbFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxbFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxbFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxbFixedLineShrinkRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxbFixedLineShrinkRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxbFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxbFixedLineShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *BindAxbFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxbFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxbFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxbFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxbFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxbFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxbFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxbFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxbFixedLineShrinkRequest) GetTelB() *string {
	return s.TelB
}

func (s *BindAxbFixedLineShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxbFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucode(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxbFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAppId(v string) *BindAxbFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetAreacode(v string) *BindAxbFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetBindType(v string) *BindAxbFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExpiration(v string) *BindAxbFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetExtraShrink(v string) *BindAxbFixedLineShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOrderId(v string) *BindAxbFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetRemark(v string) *BindAxbFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxbFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetSubts(v string) *BindAxbFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxbFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelA(v string) *BindAxbFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelB(v string) *BindAxbFixedLineShrinkRequest {
	s.TelB = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTelX(v string) *BindAxbFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) SetTs(v string) *BindAxbFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *BindAxbFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
