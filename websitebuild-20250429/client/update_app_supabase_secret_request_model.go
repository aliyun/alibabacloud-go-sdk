// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSupabaseSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateAppSupabaseSecretRequest
	GetBizId() *string
	SetSecretKey(v string) *UpdateAppSupabaseSecretRequest
	GetSecretKey() *string
	SetSecretName(v string) *UpdateAppSupabaseSecretRequest
	GetSecretName() *string
	SetSecretType(v string) *UpdateAppSupabaseSecretRequest
	GetSecretType() *string
	SetSecretValue(v string) *UpdateAppSupabaseSecretRequest
	GetSecretValue() *string
}

type UpdateAppSupabaseSecretRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// ***
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// example:
	//
	// 277356_pre_auth
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// example:
	//
	// Opaque
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// example:
	//
	// 1231
	SecretValue *string `json:"SecretValue,omitempty" xml:"SecretValue,omitempty"`
}

func (s UpdateAppSupabaseSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSupabaseSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppSupabaseSecretRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppSupabaseSecretRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *UpdateAppSupabaseSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateAppSupabaseSecretRequest) GetSecretType() *string {
	return s.SecretType
}

func (s *UpdateAppSupabaseSecretRequest) GetSecretValue() *string {
	return s.SecretValue
}

func (s *UpdateAppSupabaseSecretRequest) SetBizId(v string) *UpdateAppSupabaseSecretRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAppSupabaseSecretRequest) SetSecretKey(v string) *UpdateAppSupabaseSecretRequest {
	s.SecretKey = &v
	return s
}

func (s *UpdateAppSupabaseSecretRequest) SetSecretName(v string) *UpdateAppSupabaseSecretRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateAppSupabaseSecretRequest) SetSecretType(v string) *UpdateAppSupabaseSecretRequest {
	s.SecretType = &v
	return s
}

func (s *UpdateAppSupabaseSecretRequest) SetSecretValue(v string) *UpdateAppSupabaseSecretRequest {
	s.SecretValue = &v
	return s
}

func (s *UpdateAppSupabaseSecretRequest) Validate() error {
	return dara.Validate(s)
}
