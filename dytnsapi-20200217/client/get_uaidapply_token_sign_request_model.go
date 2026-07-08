// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDApplyTokenSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetUAIDApplyTokenSignRequest
	GetAuthCode() *string
	SetCarrier(v string) *GetUAIDApplyTokenSignRequest
	GetCarrier() *string
	SetClientType(v string) *GetUAIDApplyTokenSignRequest
	GetClientType() *string
	SetFormat(v string) *GetUAIDApplyTokenSignRequest
	GetFormat() *string
	SetOutId(v string) *GetUAIDApplyTokenSignRequest
	GetOutId() *string
	SetOwnerId(v int64) *GetUAIDApplyTokenSignRequest
	GetOwnerId() *int64
	SetParamKey(v string) *GetUAIDApplyTokenSignRequest
	GetParamKey() *string
	SetParamStr(v string) *GetUAIDApplyTokenSignRequest
	GetParamStr() *string
	SetResourceOwnerAccount(v string) *GetUAIDApplyTokenSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetUAIDApplyTokenSignRequest
	GetResourceOwnerId() *int64
	SetTime(v string) *GetUAIDApplyTokenSignRequest
	GetTime() *string
}

type GetUAIDApplyTokenSignRequest struct {
	// The authorization code.
	//
	// > To obtain this authorization code, navigate to **Tag Plaza*	- in the [**Phone Number Verification Service**](https://dytns.console.aliyun.com/analysis/square) console, select a tag, and submit an application. You receive the code after your application is approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// HwD97InG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The mobile carrier. Valid values:
	//
	// - **CM**: China Mobile
	//
	// - **CU**: China Unicom
	//
	// - **CT**: China Telecom
	//
	// This parameter is required.
	//
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The client type. Valid values:
	//
	// - `30100`: Android
	//
	// - `30300`: iOS
	//
	// - `20200`: H5
	//
	// - `10010`: Web
	//
	// This parameter is required.
	//
	// example:
	//
	// 30300
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 示例值示例值
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The external ID.
	//
	// > For China Mobile (CM), this parameter corresponds to `traceId` and `msgId`. The values of `OutId`, `traceId`, and `msgId` must be the same.
	//
	// This parameter is required.
	//
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required if the carrier is China Telecom (CT). For details, see the China Telecom documentation.
	//
	// example:
	//
	// 955EC1A869852EA8BC66F********D7C6E92017BBD5B001C736EFEAFB775C232
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// This parameter is required if the carrier is China Telecom (CT). For details, see the China Telecom documentation.
	//
	// example:
	//
	// B2F0D4CD7A216D16CE2AF4BBC********29A454FDDD991F919106C12CB89ABA8
	ParamStr             *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The event timestamp, accurate to the millisecond.<br>
	//
	// Format: `yyyyMMddHHmmssSSS`.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 20121227180001165
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetUAIDApplyTokenSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDApplyTokenSignRequest) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetUAIDApplyTokenSignRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *GetUAIDApplyTokenSignRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetUAIDApplyTokenSignRequest) GetFormat() *string {
	return s.Format
}

func (s *GetUAIDApplyTokenSignRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetUAIDApplyTokenSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetUAIDApplyTokenSignRequest) GetParamKey() *string {
	return s.ParamKey
}

func (s *GetUAIDApplyTokenSignRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *GetUAIDApplyTokenSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetUAIDApplyTokenSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetUAIDApplyTokenSignRequest) GetTime() *string {
	return s.Time
}

func (s *GetUAIDApplyTokenSignRequest) SetAuthCode(v string) *GetUAIDApplyTokenSignRequest {
	s.AuthCode = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetCarrier(v string) *GetUAIDApplyTokenSignRequest {
	s.Carrier = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetClientType(v string) *GetUAIDApplyTokenSignRequest {
	s.ClientType = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetFormat(v string) *GetUAIDApplyTokenSignRequest {
	s.Format = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetOutId(v string) *GetUAIDApplyTokenSignRequest {
	s.OutId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetOwnerId(v int64) *GetUAIDApplyTokenSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetParamKey(v string) *GetUAIDApplyTokenSignRequest {
	s.ParamKey = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetParamStr(v string) *GetUAIDApplyTokenSignRequest {
	s.ParamStr = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetResourceOwnerAccount(v string) *GetUAIDApplyTokenSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetResourceOwnerId(v int64) *GetUAIDApplyTokenSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) SetTime(v string) *GetUAIDApplyTokenSignRequest {
	s.Time = &v
	return s
}

func (s *GetUAIDApplyTokenSignRequest) Validate() error {
	return dara.Validate(s)
}
