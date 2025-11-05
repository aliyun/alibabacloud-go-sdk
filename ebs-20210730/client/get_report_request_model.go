// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetReportRequest
	GetAppName() *string
	SetRegionId(v string) *GetReportRequest
	GetRegionId() *string
	SetReportId(v string) *GetReportRequest
	GetReportId() *string
	SetReportType(v string) *GetReportRequest
	GetReportType() *string
}

type GetReportRequest struct {
	// Optional, AppName only takes effect when ReportType=present.
	//
	// example:
	//
	// App1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Region name.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// When ReportType=history, ReportId is required to query historical reports based on ReportId.
	//
	// example:
	//
	// report-74fbea80e802xxxx
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// Optional values: history/present.
	//
	// example:
	//
	// history
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportRequest) GoString() string {
	return s.String()
}

func (s *GetReportRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetReportRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetReportRequest) SetAppName(v string) *GetReportRequest {
	s.AppName = &v
	return s
}

func (s *GetReportRequest) SetRegionId(v string) *GetReportRequest {
	s.RegionId = &v
	return s
}

func (s *GetReportRequest) SetReportId(v string) *GetReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetReportRequest) SetReportType(v string) *GetReportRequest {
	s.ReportType = &v
	return s
}

func (s *GetReportRequest) Validate() error {
	return dara.Validate(s)
}
