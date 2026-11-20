// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetExpiration() *string
	SetExtraaxxShrink(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetExtraaxxShrink() *string
	SetOrderId(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelA() *string
	SetTelX(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelX() *string
	SetTelXext(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTelXext() *string
	SetTs(v string) *BindAxnExtensionFixedLineShrinkRequest
	GetTs() *string
}

type BindAxnExtensionFixedLineShrinkRequest struct {
	// Caller-side ringtone code. This parameter is required for AXB scenarios. Must contain codes for 3 scenarios. Enter the codes in the order of "B->X, A->X, other numbers->X", separated by commas.
	//
	// > "1,2,3" indicates that the ringtone code for B->X is 1, for A->X is 2, and for other numbers->X is 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	Anucode *string `json:"Anucode,omitempty" xml:"Anucode,omitempty"`
	// Called-side ringtone code after connection. Must contain codes for 2 scenarios. Enter the codes in the order of "A as called party, B as called party", separated by commas.
	//
	// > "1,2" indicates that the ringtone code when A is the called party is 1, and the ringtone code when B is the called party is 2.
	//
	// example:
	//
	// 1,2
	Anucodecalled *string `json:"Anucodecalled,omitempty" xml:"Anucodecalled,omitempty"`
	// Number pool ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Area code. Supports specifying the home location of the X number (specified city -> provincial capital -> nationwide routing). If no home location is specified, a number is randomly assigned nationwide.
	//
	// example:
	//
	// 0571
	Areacode *string `json:"Areacode,omitempty" xml:"Areacode,omitempty"`
	// Binding type. Valid values:
	//
	// - mode101: Customer specifies the X number.
	//
	// - mode102: Platform assigns the X number.
	//
	// This parameter is required.
	//
	// example:
	//
	// mode101
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// Expiration time. Unit: seconds. Valid values:
	//
	// - 0: Never expires.
	//
	// - Other values: When other values are specified, such as 60, the binding is valid for 60 seconds after the X number is successfully assigned. After 60 seconds, the binding is automatically released.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// Extension parameters.
	ExtraaxxShrink *string `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty"`
	// Unique identifier of the message request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// External business extension field. This parameter is returned in call record receipt messages. Must not exceed 100 characters.
	//
	// example:
	//
	// 19394
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Binding time. Format: yyyyMMddHHmmss. Uses Beijing time, 24-hour clock.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116
	Subts *string `json:"Subts,omitempty" xml:"Subts,omitempty"`
	// Caller-side ringtone code after connection. Must contain codes for 2 scenarios. Enter the codes in the order of "A as called party, B as called party", separated by commas.
	//
	// > "1,2" indicates that the caller-side ringtone code when A is the called party is 1, and the caller-side ringtone code when B is the called party is 2.
	//
	// example:
	//
	// 1,2
	TAnucodeConnect *string `json:"TAnucodeConnect,omitempty" xml:"TAnucodeConnect,omitempty"`
	// The A number for the fixed-line AXN extension number.
	//
	// > The A number can be set to a mobile number or a fixed-line number. Fixed-line numbers must include the area code, and no hyphen is needed between the area code and the number.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mobile: 182****2371
	//
	// Fixed line: 0571******36
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// The X number for the fixed-line AXN extension number. The X number is a privacy number that you have [purchased](https://dyplsnext.console.aliyun.com/numberbuy) in the console.
	//
	// example:
	//
	// 0571******23
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// Extension number. This parameter is required when the BindType parameter is set to mode101.
	//
	// example:
	//
	// 1009
	TelXext *string `json:"TelXext,omitempty" xml:"TelXext,omitempty"`
	// Business timestamp. Format: yyyyMMddHHmmssSSS. Uses Beijing time, 24-hour clock, accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s BindAxnExtensionFixedLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetExtraaxxShrink() *string {
	return s.ExtraaxxShrink
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTelXext() *string {
	return s.TelXext
}

func (s *BindAxnExtensionFixedLineShrinkRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAppId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetAreacode(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetBindType(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExpiration(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetExtraaxxShrink(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ExtraaxxShrink = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOrderId(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetRemark(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetSubts(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelA(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelX(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTelXext(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) SetTs(v string) *BindAxnExtensionFixedLineShrinkRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnExtensionFixedLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
