// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJWTPayload interface {
	dara.Model
	String() string
	GoString() string
	SetAud(v string) *JWTPayload
	GetAud() *string
	SetAutoCreate(v bool) *JWTPayload
	GetAutoCreate() *bool
	SetExp(v int64) *JWTPayload
	GetExp() *int64
	SetIat(v int64) *JWTPayload
	GetIat() *int64
	SetIss(v string) *JWTPayload
	GetIss() *string
	SetJti(v string) *JWTPayload
	GetJti() *string
	SetNbf(v int64) *JWTPayload
	GetNbf() *int64
	SetSub(v string) *JWTPayload
	GetSub() *string
	SetSubType(v string) *JWTPayload
	GetSubType() *string
}

type JWTPayload struct {
	Aud        *string `json:"aud,omitempty" xml:"aud,omitempty"`
	AutoCreate *bool   `json:"auto_create,omitempty" xml:"auto_create,omitempty"`
	Exp        *int64  `json:"exp,omitempty" xml:"exp,omitempty"`
	Iat        *int64  `json:"iat,omitempty" xml:"iat,omitempty"`
	Iss        *string `json:"iss,omitempty" xml:"iss,omitempty"`
	Jti        *string `json:"jti,omitempty" xml:"jti,omitempty"`
	Nbf        *int64  `json:"nbf,omitempty" xml:"nbf,omitempty"`
	Sub        *string `json:"sub,omitempty" xml:"sub,omitempty"`
	SubType    *string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}

func (s JWTPayload) String() string {
	return dara.Prettify(s)
}

func (s JWTPayload) GoString() string {
	return s.String()
}

func (s *JWTPayload) GetAud() *string {
	return s.Aud
}

func (s *JWTPayload) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *JWTPayload) GetExp() *int64 {
	return s.Exp
}

func (s *JWTPayload) GetIat() *int64 {
	return s.Iat
}

func (s *JWTPayload) GetIss() *string {
	return s.Iss
}

func (s *JWTPayload) GetJti() *string {
	return s.Jti
}

func (s *JWTPayload) GetNbf() *int64 {
	return s.Nbf
}

func (s *JWTPayload) GetSub() *string {
	return s.Sub
}

func (s *JWTPayload) GetSubType() *string {
	return s.SubType
}

func (s *JWTPayload) SetAud(v string) *JWTPayload {
	s.Aud = &v
	return s
}

func (s *JWTPayload) SetAutoCreate(v bool) *JWTPayload {
	s.AutoCreate = &v
	return s
}

func (s *JWTPayload) SetExp(v int64) *JWTPayload {
	s.Exp = &v
	return s
}

func (s *JWTPayload) SetIat(v int64) *JWTPayload {
	s.Iat = &v
	return s
}

func (s *JWTPayload) SetIss(v string) *JWTPayload {
	s.Iss = &v
	return s
}

func (s *JWTPayload) SetJti(v string) *JWTPayload {
	s.Jti = &v
	return s
}

func (s *JWTPayload) SetNbf(v int64) *JWTPayload {
	s.Nbf = &v
	return s
}

func (s *JWTPayload) SetSub(v string) *JWTPayload {
	s.Sub = &v
	return s
}

func (s *JWTPayload) SetSubType(v string) *JWTPayload {
	s.SubType = &v
	return s
}

func (s *JWTPayload) Validate() error {
	return dara.Validate(s)
}
