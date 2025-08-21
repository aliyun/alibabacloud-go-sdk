// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CloseDiagnosisRequest
	GetClientToken() *string
	SetLang(v string) *CloseDiagnosisRequest
	GetLang() *string
}

type CloseDiagnosisRequest struct {
	// The ID of the request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// spanish
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s CloseDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloseDiagnosisRequest) GetLang() *string {
	return s.Lang
}

func (s *CloseDiagnosisRequest) SetClientToken(v string) *CloseDiagnosisRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseDiagnosisRequest) SetLang(v string) *CloseDiagnosisRequest {
	s.Lang = &v
	return s
}

func (s *CloseDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
