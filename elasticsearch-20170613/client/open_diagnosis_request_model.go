// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenDiagnosisRequest
	GetClientToken() *string
	SetLang(v string) *OpenDiagnosisRequest
	GetLang() *string
}

type OpenDiagnosisRequest struct {
	// The ID of the request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s OpenDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenDiagnosisRequest) GetLang() *string {
	return s.Lang
}

func (s *OpenDiagnosisRequest) SetClientToken(v string) *OpenDiagnosisRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenDiagnosisRequest) SetLang(v string) *OpenDiagnosisRequest {
	s.Lang = &v
	return s
}

func (s *OpenDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
