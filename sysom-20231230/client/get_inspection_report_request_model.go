// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetInspectionReportRequest
	GetXDebugId() *string
	SetReportId(v string) *GetInspectionReportRequest
	GetReportId() *string
	SetXSysomInvokeSource(v string) *GetInspectionReportRequest
	GetXSysomInvokeSource() *string
}

type GetInspectionReportRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The inspection report ID.
	//
	// example:
	//
	// 91bec4c5a168494e8128468e3995df87
	ReportId           *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetInspectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportRequest) GoString() string {
	return s.String()
}

func (s *GetInspectionReportRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetInspectionReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetInspectionReportRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetInspectionReportRequest) SetXDebugId(v string) *GetInspectionReportRequest {
	s.XDebugId = &v
	return s
}

func (s *GetInspectionReportRequest) SetReportId(v string) *GetInspectionReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetInspectionReportRequest) SetXSysomInvokeSource(v string) *GetInspectionReportRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetInspectionReportRequest) Validate() error {
	return dara.Validate(s)
}
