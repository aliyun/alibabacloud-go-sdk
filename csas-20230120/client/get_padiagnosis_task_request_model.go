// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPADiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnoseId(v string) *GetPADiagnosisTaskRequest
	GetDiagnoseId() *string
}

type GetPADiagnosisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// diag-58d0750e8786919a
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
}

func (s GetPADiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskRequest) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *GetPADiagnosisTaskRequest) SetDiagnoseId(v string) *GetPADiagnosisTaskRequest {
	s.DiagnoseId = &v
	return s
}

func (s *GetPADiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
