// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeNisInspectionReportStatusResponseBody
	GetEndTime() *string
	SetInspectionProject(v string) *DescribeNisInspectionReportStatusResponseBody
	GetInspectionProject() *string
	SetInspectionReportId(v string) *DescribeNisInspectionReportStatusResponseBody
	GetInspectionReportId() *string
	SetInspectionTaskId(v string) *DescribeNisInspectionReportStatusResponseBody
	GetInspectionTaskId() *string
	SetInspectionTaskName(v string) *DescribeNisInspectionReportStatusResponseBody
	GetInspectionTaskName() *string
	SetRequestId(v string) *DescribeNisInspectionReportStatusResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeNisInspectionReportStatusResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeNisInspectionReportStatusResponseBody
	GetStatus() *string
}

type DescribeNisInspectionReportStatusResponseBody struct {
	// example:
	//
	// 2024-07-18 15:13:07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// basic
	InspectionProject *string `json:"InspectionProject,omitempty" xml:"InspectionProject,omitempty"`
	// example:
	//
	// nir-2ca527b8de114ba4afb9
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// example:
	//
	// ni-8svmpe0yso****r7fh79
	InspectionTaskId   *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	InspectionTaskName *string `json:"InspectionTaskName,omitempty" xml:"InspectionTaskName,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2024-07-18 15:12:28
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNisInspectionReportStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetInspectionProject() *string {
	return s.InspectionProject
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetInspectionTaskName() *string {
	return s.InspectionTaskName
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeNisInspectionReportStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetEndTime(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetInspectionProject(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.InspectionProject = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetInspectionReportId(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetInspectionTaskId(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.InspectionTaskId = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetInspectionTaskName(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.InspectionTaskName = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetRequestId(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetStartTime(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) SetStatus(v string) *DescribeNisInspectionReportStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
