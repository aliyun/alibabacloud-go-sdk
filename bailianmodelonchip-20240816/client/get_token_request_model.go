// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetTokenRequest
	GetAppId() *string
	SetDeviceName(v string) *GetTokenRequest
	GetDeviceName() *string
	SetNonce(v string) *GetTokenRequest
	GetNonce() *string
	SetRequestTime(v string) *GetTokenRequest
	GetRequestTime() *string
	SetSignature(v string) *GetTokenRequest
	GetSignature() *string
	SetTokenKey(v string) *GetTokenRequest
	GetTokenKey() *string
	SetTokenType(v string) *GetTokenRequest
	GetTokenType() *string
}

type GetTokenRequest struct {
	// This parameter is required.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5b504f84b69b9a73d3a21a2cff05e190
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748413148546
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5/Smm8gnDSgZY2Blftq9eGYpVnpYCztoLJaJfhlH7id0lNlQxydRLtjUkGPr1qdbQq+oUn6Y1h0KJUdk0rf4am3MzdNr/Uhc47c8bbXnV8SlIC0agGo5skEQZNObzUD+sFxt8uN35/FYf7YRC8R61xY7+NPN2NLJrA/DPhewtVRRgAbb8RjergTcIG6oN1XTzLyC+76Az/3o2dPCxTfMXG3AFQc=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// sk-4asv136677d2411b876e536bc8xxxxx
	TokenKey *string `json:"tokenKey,omitempty" xml:"tokenKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetTokenRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetTokenRequest) GetNonce() *string {
	return s.Nonce
}

func (s *GetTokenRequest) GetRequestTime() *string {
	return s.RequestTime
}

func (s *GetTokenRequest) GetSignature() *string {
	return s.Signature
}

func (s *GetTokenRequest) GetTokenKey() *string {
	return s.TokenKey
}

func (s *GetTokenRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *GetTokenRequest) SetAppId(v string) *GetTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetTokenRequest) SetDeviceName(v string) *GetTokenRequest {
	s.DeviceName = &v
	return s
}

func (s *GetTokenRequest) SetNonce(v string) *GetTokenRequest {
	s.Nonce = &v
	return s
}

func (s *GetTokenRequest) SetRequestTime(v string) *GetTokenRequest {
	s.RequestTime = &v
	return s
}

func (s *GetTokenRequest) SetSignature(v string) *GetTokenRequest {
	s.Signature = &v
	return s
}

func (s *GetTokenRequest) SetTokenKey(v string) *GetTokenRequest {
	s.TokenKey = &v
	return s
}

func (s *GetTokenRequest) SetTokenType(v string) *GetTokenRequest {
	s.TokenType = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
