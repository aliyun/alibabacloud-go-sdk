// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnucode(v string) *BindAxnExtensionFixedLineRequest
	GetAnucode() *string
	SetAnucodecalled(v string) *BindAxnExtensionFixedLineRequest
	GetAnucodecalled() *string
	SetAppId(v string) *BindAxnExtensionFixedLineRequest
	GetAppId() *string
	SetAreacode(v string) *BindAxnExtensionFixedLineRequest
	GetAreacode() *string
	SetBindType(v string) *BindAxnExtensionFixedLineRequest
	GetBindType() *string
	SetExpiration(v string) *BindAxnExtensionFixedLineRequest
	GetExpiration() *string
	SetExtraaxx(v *BindAxnExtensionFixedLineRequestExtraaxx) *BindAxnExtensionFixedLineRequest
	GetExtraaxx() *BindAxnExtensionFixedLineRequestExtraaxx
	SetOrderId(v string) *BindAxnExtensionFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *BindAxnExtensionFixedLineRequest
	GetOwnerId() *int64
	SetRemark(v string) *BindAxnExtensionFixedLineRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubts(v string) *BindAxnExtensionFixedLineRequest
	GetSubts() *string
	SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineRequest
	GetTAnucodeConnect() *string
	SetTelA(v string) *BindAxnExtensionFixedLineRequest
	GetTelA() *string
	SetTelX(v string) *BindAxnExtensionFixedLineRequest
	GetTelX() *string
	SetTelXext(v string) *BindAxnExtensionFixedLineRequest
	GetTelXext() *string
	SetTs(v string) *BindAxnExtensionFixedLineRequest
	GetTs() *string
}

type BindAxnExtensionFixedLineRequest struct {
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
	Extraaxx *BindAxnExtensionFixedLineRequestExtraaxx `json:"Extraaxx,omitempty" xml:"Extraaxx,omitempty" type:"Struct"`
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

func (s BindAxnExtensionFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequest) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequest) GetAnucode() *string {
	return s.Anucode
}

func (s *BindAxnExtensionFixedLineRequest) GetAnucodecalled() *string {
	return s.Anucodecalled
}

func (s *BindAxnExtensionFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindAxnExtensionFixedLineRequest) GetAreacode() *string {
	return s.Areacode
}

func (s *BindAxnExtensionFixedLineRequest) GetBindType() *string {
	return s.BindType
}

func (s *BindAxnExtensionFixedLineRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *BindAxnExtensionFixedLineRequest) GetExtraaxx() *BindAxnExtensionFixedLineRequestExtraaxx {
	return s.Extraaxx
}

func (s *BindAxnExtensionFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *BindAxnExtensionFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindAxnExtensionFixedLineRequest) GetRemark() *string {
	return s.Remark
}

func (s *BindAxnExtensionFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindAxnExtensionFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindAxnExtensionFixedLineRequest) GetSubts() *string {
	return s.Subts
}

func (s *BindAxnExtensionFixedLineRequest) GetTAnucodeConnect() *string {
	return s.TAnucodeConnect
}

func (s *BindAxnExtensionFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *BindAxnExtensionFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *BindAxnExtensionFixedLineRequest) GetTelXext() *string {
	return s.TelXext
}

func (s *BindAxnExtensionFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucode(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAnucodecalled(v string) *BindAxnExtensionFixedLineRequest {
	s.Anucodecalled = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAppId(v string) *BindAxnExtensionFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetAreacode(v string) *BindAxnExtensionFixedLineRequest {
	s.Areacode = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetBindType(v string) *BindAxnExtensionFixedLineRequest {
	s.BindType = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExpiration(v string) *BindAxnExtensionFixedLineRequest {
	s.Expiration = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetExtraaxx(v *BindAxnExtensionFixedLineRequestExtraaxx) *BindAxnExtensionFixedLineRequest {
	s.Extraaxx = v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOrderId(v string) *BindAxnExtensionFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetRemark(v string) *BindAxnExtensionFixedLineRequest {
	s.Remark = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerAccount(v string) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetResourceOwnerId(v int64) *BindAxnExtensionFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetSubts(v string) *BindAxnExtensionFixedLineRequest {
	s.Subts = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTAnucodeConnect(v string) *BindAxnExtensionFixedLineRequest {
	s.TAnucodeConnect = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelA(v string) *BindAxnExtensionFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelX(v string) *BindAxnExtensionFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTelXext(v string) *BindAxnExtensionFixedLineRequest {
	s.TelXext = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) SetTs(v string) *BindAxnExtensionFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequest) Validate() error {
	if s.Extraaxx != nil {
		if err := s.Extraaxx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAxnExtensionFixedLineRequestExtraaxx struct {
	// Callback control. When A calls through X, that is, A dials X as the caller. Valid values:
	//
	// - 0 (default): Cannot make outbound calls.
	//
	// - 1: Connect to the most recent N number.
	//
	// - 2: Call back the fixed number TelB.
	//
	// example:
	//
	// 0
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// Recording control. Valid values:
	//
	// - 0 (default): No recording.
	//
	// - 1: Recording.
	//
	// > After you enable the call recording feature, you can play and download audio files in the call records. Recording files are billed at CNY 0.010/minute.
	//
	// example:
	//
	// 0
	Callrecording *string `json:"Callrecording,omitempty" xml:"Callrecording,omitempty"`
}

func (s BindAxnExtensionFixedLineRequestExtraaxx) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineRequestExtraaxx) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) GetCallback() *string {
	return s.Callback
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) GetCallrecording() *string {
	return s.Callrecording
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallback(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callback = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) SetCallrecording(v string) *BindAxnExtensionFixedLineRequestExtraaxx {
	s.Callrecording = &v
	return s
}

func (s *BindAxnExtensionFixedLineRequestExtraaxx) Validate() error {
	return dara.Validate(s)
}
