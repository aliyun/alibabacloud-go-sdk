// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDiagnosticReportListRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDiagnosticReportListRequest
	GetEndTime() *string
	SetPageNo(v string) *DescribeDiagnosticReportListRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeDiagnosticReportListRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeDiagnosticReportListRequest
	GetStartTime() *string
}

type DescribeDiagnosticReportListRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1668420492000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1668398892000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosticReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDiagnosticReportListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosticReportListRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeDiagnosticReportListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDiagnosticReportListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnosticReportListRequest) SetDBInstanceId(v string) *DescribeDiagnosticReportListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetEndTime(v string) *DescribeDiagnosticReportListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetPageNo(v string) *DescribeDiagnosticReportListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetPageSize(v string) *DescribeDiagnosticReportListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) SetStartTime(v string) *DescribeDiagnosticReportListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosticReportListRequest) Validate() error {
	return dara.Validate(s)
}
