// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GetInspectionReportRequest
	GetReportId() *string
}

type GetInspectionReportRequest struct {
	// example:
	//
	// 91bec4c5a168494e8128468e3995df87
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s GetInspectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportRequest) GoString() string {
	return s.String()
}

func (s *GetInspectionReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetInspectionReportRequest) SetReportId(v string) *GetInspectionReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetInspectionReportRequest) Validate() error {
	return dara.Validate(s)
}
