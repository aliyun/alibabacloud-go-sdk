// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDiagnosisSettingsRequest
	GetLang() *string
}

type DescribeDiagnosisSettingsRequest struct {
	// The language of the returned result. Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s DescribeDiagnosisSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnosisSettingsRequest) SetLang(v string) *DescribeDiagnosisSettingsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisSettingsRequest) Validate() error {
	return dara.Validate(s)
}
