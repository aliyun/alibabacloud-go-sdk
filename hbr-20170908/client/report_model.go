// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReport interface {
	dara.Model
	String() string
	GoString() string
	SetFailedFiles(v string) *Report
	GetFailedFiles() *string
	SetSkippedFiles(v string) *Report
	GetSkippedFiles() *string
	SetSuccessFiles(v string) *Report
	GetSuccessFiles() *string
	SetTotalFiles(v string) *Report
	GetTotalFiles() *string
}

type Report struct {
	// example:
	//
	// temp/report/r-0000dnz7p4pk31u6madf_failed.csv
	FailedFiles *string `json:"FailedFiles,omitempty" xml:"FailedFiles,omitempty"`
	// example:
	//
	// temp/report/r-0000dnz7p4pk31u6madf_skipped.csv
	SkippedFiles *string `json:"SkippedFiles,omitempty" xml:"SkippedFiles,omitempty"`
	// example:
	//
	// temp/report/r-0000dnz7p4pk31u6madf_success.csv
	SuccessFiles *string `json:"SuccessFiles,omitempty" xml:"SuccessFiles,omitempty"`
	// example:
	//
	// temp/report/r-0000dnz7p4pk31u6madf_total.csv
	TotalFiles *string `json:"TotalFiles,omitempty" xml:"TotalFiles,omitempty"`
}

func (s Report) String() string {
	return dara.Prettify(s)
}

func (s Report) GoString() string {
	return s.String()
}

func (s *Report) GetFailedFiles() *string {
	return s.FailedFiles
}

func (s *Report) GetSkippedFiles() *string {
	return s.SkippedFiles
}

func (s *Report) GetSuccessFiles() *string {
	return s.SuccessFiles
}

func (s *Report) GetTotalFiles() *string {
	return s.TotalFiles
}

func (s *Report) SetFailedFiles(v string) *Report {
	s.FailedFiles = &v
	return s
}

func (s *Report) SetSkippedFiles(v string) *Report {
	s.SkippedFiles = &v
	return s
}

func (s *Report) SetSuccessFiles(v string) *Report {
	s.SuccessFiles = &v
	return s
}

func (s *Report) SetTotalFiles(v string) *Report {
	s.TotalFiles = &v
	return s
}

func (s *Report) Validate() error {
	return dara.Validate(s)
}
