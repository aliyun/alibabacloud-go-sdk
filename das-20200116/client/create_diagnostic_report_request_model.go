// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDiagnosticReportRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *CreateDiagnosticReportRequest
	GetEndTime() *string
	SetStartTime(v string) *CreateDiagnosticReportRequest
	GetStartTime() *string
}

type CreateDiagnosticReportRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to create the diagnostic report. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The start time must be later than the end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993001
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to create the diagnostic report. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDiagnosticReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnosticReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateDiagnosticReportRequest) SetDBInstanceId(v string) *CreateDiagnosticReportRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetEndTime(v string) *CreateDiagnosticReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetStartTime(v string) *CreateDiagnosticReportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) Validate() error {
	return dara.Validate(s)
}
