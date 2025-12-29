// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDConversionSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetUAIDConversionSignRequest
	GetAuthCode() *string
	SetCarrier(v string) *GetUAIDConversionSignRequest
	GetCarrier() *string
	SetClientType(v string) *GetUAIDConversionSignRequest
	GetClientType() *string
	SetFormat(v string) *GetUAIDConversionSignRequest
	GetFormat() *string
	SetOutId(v string) *GetUAIDConversionSignRequest
	GetOutId() *string
	SetOwnerId(v int64) *GetUAIDConversionSignRequest
	GetOwnerId() *int64
	SetParamKey(v string) *GetUAIDConversionSignRequest
	GetParamKey() *string
	SetParamStr(v string) *GetUAIDConversionSignRequest
	GetParamStr() *string
	SetResourceOwnerAccount(v string) *GetUAIDConversionSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetUAIDConversionSignRequest
	GetResourceOwnerId() *int64
	SetTime(v string) *GetUAIDConversionSignRequest
	GetTime() *string
}

type GetUAIDConversionSignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	ParamKey *string `json:"ParamKey,omitempty" xml:"ParamKey,omitempty"`
	// example:
	//
	// 示例值示例值
	ParamStr             *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetUAIDConversionSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDConversionSignRequest) GoString() string {
	return s.String()
}

func (s *GetUAIDConversionSignRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetUAIDConversionSignRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *GetUAIDConversionSignRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetUAIDConversionSignRequest) GetFormat() *string {
	return s.Format
}

func (s *GetUAIDConversionSignRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetUAIDConversionSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetUAIDConversionSignRequest) GetParamKey() *string {
	return s.ParamKey
}

func (s *GetUAIDConversionSignRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *GetUAIDConversionSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetUAIDConversionSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetUAIDConversionSignRequest) GetTime() *string {
	return s.Time
}

func (s *GetUAIDConversionSignRequest) SetAuthCode(v string) *GetUAIDConversionSignRequest {
	s.AuthCode = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetCarrier(v string) *GetUAIDConversionSignRequest {
	s.Carrier = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetClientType(v string) *GetUAIDConversionSignRequest {
	s.ClientType = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetFormat(v string) *GetUAIDConversionSignRequest {
	s.Format = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetOutId(v string) *GetUAIDConversionSignRequest {
	s.OutId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetOwnerId(v int64) *GetUAIDConversionSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetParamKey(v string) *GetUAIDConversionSignRequest {
	s.ParamKey = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetParamStr(v string) *GetUAIDConversionSignRequest {
	s.ParamStr = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetResourceOwnerAccount(v string) *GetUAIDConversionSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetResourceOwnerId(v int64) *GetUAIDConversionSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUAIDConversionSignRequest) SetTime(v string) *GetUAIDConversionSignRequest {
	s.Time = &v
	return s
}

func (s *GetUAIDConversionSignRequest) Validate() error {
	return dara.Validate(s)
}
