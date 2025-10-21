// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDiagnosisHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetJobDiagnosisHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetJobDiagnosisHeaders
	GetWorkspace() *string
}

type GetJobDiagnosisHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetJobDiagnosisHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetJobDiagnosisHeaders) GoString() string {
	return s.String()
}

func (s *GetJobDiagnosisHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetJobDiagnosisHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetJobDiagnosisHeaders) SetCommonHeaders(v map[string]*string) *GetJobDiagnosisHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobDiagnosisHeaders) SetWorkspace(v string) *GetJobDiagnosisHeaders {
	s.Workspace = &v
	return s
}

func (s *GetJobDiagnosisHeaders) Validate() error {
	return dara.Validate(s)
}
