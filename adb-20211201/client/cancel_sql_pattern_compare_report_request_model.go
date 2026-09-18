// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPatternCompareReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelSqlPatternCompareReportRequest
	GetDBClusterId() *string
	SetRegionId(v string) *CancelSqlPatternCompareReportRequest
	GetRegionId() *string
	SetReportId(v int64) *CancelSqlPatternCompareReportRequest
	GetReportId() *int64
}

type CancelSqlPatternCompareReportRequest struct {
	// The AnalyticDB for MySQL instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze1234567890****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SQL Pattern comparison report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s CancelSqlPatternCompareReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPatternCompareReportRequest) GoString() string {
	return s.String()
}

func (s *CancelSqlPatternCompareReportRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelSqlPatternCompareReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelSqlPatternCompareReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *CancelSqlPatternCompareReportRequest) SetDBClusterId(v string) *CancelSqlPatternCompareReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelSqlPatternCompareReportRequest) SetRegionId(v string) *CancelSqlPatternCompareReportRequest {
	s.RegionId = &v
	return s
}

func (s *CancelSqlPatternCompareReportRequest) SetReportId(v int64) *CancelSqlPatternCompareReportRequest {
	s.ReportId = &v
	return s
}

func (s *CancelSqlPatternCompareReportRequest) Validate() error {
	return dara.Validate(s)
}
