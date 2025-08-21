// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDiagnoseReportRequest
	GetLang() *string
}

type DescribeDiagnoseReportRequest struct {
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s DescribeDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDiagnoseReportRequest) SetLang(v string) *DescribeDiagnoseReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
