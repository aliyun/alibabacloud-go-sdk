// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizeReportStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPinnedTime(v int64) *UpdateCustomizeReportStatusRequest
	GetPinnedTime() *int64
	SetReportId(v int64) *UpdateCustomizeReportStatusRequest
	GetReportId() *int64
	SetReportStatus(v int32) *UpdateCustomizeReportStatusRequest
	GetReportStatus() *int32
}

type UpdateCustomizeReportStatusRequest struct {
	// The time when the report is pinned. Unit: milliseconds.
	//
	// example:
	//
	// 1717430400000
	PinnedTime *int64 `json:"PinnedTime,omitempty" xml:"PinnedTime,omitempty"`
	// The ID of the report.
	//
	// >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportStatus *int32 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
}

func (s UpdateCustomizeReportStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizeReportStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomizeReportStatusRequest) GetPinnedTime() *int64 {
	return s.PinnedTime
}

func (s *UpdateCustomizeReportStatusRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *UpdateCustomizeReportStatusRequest) GetReportStatus() *int32 {
	return s.ReportStatus
}

func (s *UpdateCustomizeReportStatusRequest) SetPinnedTime(v int64) *UpdateCustomizeReportStatusRequest {
	s.PinnedTime = &v
	return s
}

func (s *UpdateCustomizeReportStatusRequest) SetReportId(v int64) *UpdateCustomizeReportStatusRequest {
	s.ReportId = &v
	return s
}

func (s *UpdateCustomizeReportStatusRequest) SetReportStatus(v int32) *UpdateCustomizeReportStatusRequest {
	s.ReportStatus = &v
	return s
}

func (s *UpdateCustomizeReportStatusRequest) Validate() error {
	return dara.Validate(s)
}
