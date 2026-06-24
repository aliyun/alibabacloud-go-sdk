// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiagnosisSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateDiagnosisSettingsRequest
	GetClientToken() *string
	SetBody(v string) *UpdateDiagnosisSettingsRequest
	GetBody() *string
	SetLang(v string) *UpdateDiagnosisSettingsRequest
	GetLang() *string
}

type UpdateDiagnosisSettingsRequest struct {
	// A unique token used to ensure idempotence of the request. The client generates this value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
	// The language of the response. Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s UpdateDiagnosisSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiagnosisSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDiagnosisSettingsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateDiagnosisSettingsRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDiagnosisSettingsRequest) SetClientToken(v string) *UpdateDiagnosisSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDiagnosisSettingsRequest) SetBody(v string) *UpdateDiagnosisSettingsRequest {
	s.Body = &v
	return s
}

func (s *UpdateDiagnosisSettingsRequest) SetLang(v string) *UpdateDiagnosisSettingsRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDiagnosisSettingsRequest) Validate() error {
	return dara.Validate(s)
}
