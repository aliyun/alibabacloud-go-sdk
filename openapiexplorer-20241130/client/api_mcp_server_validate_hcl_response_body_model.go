// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiMcpServerValidateHclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticReport(v interface{}) *ApiMcpServerValidateHclResponseBody
	GetDiagnosticReport() interface{}
	SetErrors(v []*string) *ApiMcpServerValidateHclResponseBody
	GetErrors() []*string
	SetHash(v string) *ApiMcpServerValidateHclResponseBody
	GetHash() *string
	SetIsValid(v bool) *ApiMcpServerValidateHclResponseBody
	GetIsValid() *bool
	SetParameters(v []interface{}) *ApiMcpServerValidateHclResponseBody
	GetParameters() []interface{}
	SetRequestId(v string) *ApiMcpServerValidateHclResponseBody
	GetRequestId() *string
	SetWarnings(v []*string) *ApiMcpServerValidateHclResponseBody
	GetWarnings() []*string
}

type ApiMcpServerValidateHclResponseBody struct {
	// example:
	//
	// Argument or block definition required: An argument or block definition is required here.
	DiagnosticReport interface{} `json:"diagnosticReport,omitempty" xml:"diagnosticReport,omitempty"`
	Errors           []*string   `json:"errors,omitempty" xml:"errors,omitempty" type:"Repeated"`
	// example:
	//
	// 0628e13692023222bef9d6377dd03da3304b689e1b2df60f584ea27b4163bf07
	Hash *string `json:"hash,omitempty" xml:"hash,omitempty"`
	// example:
	//
	// true
	IsValid    *bool         `json:"isValid,omitempty" xml:"isValid,omitempty"`
	Parameters []interface{} `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Warnings  []*string `json:"warnings,omitempty" xml:"warnings,omitempty" type:"Repeated"`
}

func (s ApiMcpServerValidateHclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApiMcpServerValidateHclResponseBody) GoString() string {
	return s.String()
}

func (s *ApiMcpServerValidateHclResponseBody) GetDiagnosticReport() interface{} {
	return s.DiagnosticReport
}

func (s *ApiMcpServerValidateHclResponseBody) GetErrors() []*string {
	return s.Errors
}

func (s *ApiMcpServerValidateHclResponseBody) GetHash() *string {
	return s.Hash
}

func (s *ApiMcpServerValidateHclResponseBody) GetIsValid() *bool {
	return s.IsValid
}

func (s *ApiMcpServerValidateHclResponseBody) GetParameters() []interface{} {
	return s.Parameters
}

func (s *ApiMcpServerValidateHclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApiMcpServerValidateHclResponseBody) GetWarnings() []*string {
	return s.Warnings
}

func (s *ApiMcpServerValidateHclResponseBody) SetDiagnosticReport(v interface{}) *ApiMcpServerValidateHclResponseBody {
	s.DiagnosticReport = v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetErrors(v []*string) *ApiMcpServerValidateHclResponseBody {
	s.Errors = v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetHash(v string) *ApiMcpServerValidateHclResponseBody {
	s.Hash = &v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetIsValid(v bool) *ApiMcpServerValidateHclResponseBody {
	s.IsValid = &v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetParameters(v []interface{}) *ApiMcpServerValidateHclResponseBody {
	s.Parameters = v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetRequestId(v string) *ApiMcpServerValidateHclResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) SetWarnings(v []*string) *ApiMcpServerValidateHclResponseBody {
	s.Warnings = v
	return s
}

func (s *ApiMcpServerValidateHclResponseBody) Validate() error {
	return dara.Validate(s)
}
