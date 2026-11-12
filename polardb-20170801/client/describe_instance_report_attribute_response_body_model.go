// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceReportAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatest(v bool) *DescribeInstanceReportAttributeResponseBody
	GetLatest() *bool
	SetReportDescription(v string) *DescribeInstanceReportAttributeResponseBody
	GetReportDescription() *string
	SetReportId(v string) *DescribeInstanceReportAttributeResponseBody
	GetReportId() *string
	SetReportScore(v int64) *DescribeInstanceReportAttributeResponseBody
	GetReportScore() *int64
	SetReportStatus(v string) *DescribeInstanceReportAttributeResponseBody
	GetReportStatus() *string
	SetReportSubItems(v []*DescribeInstanceReportAttributeResponseBodyReportSubItems) *DescribeInstanceReportAttributeResponseBody
	GetReportSubItems() []*DescribeInstanceReportAttributeResponseBodyReportSubItems
	SetReportSummary(v string) *DescribeInstanceReportAttributeResponseBody
	GetReportSummary() *string
	SetRequestId(v string) *DescribeInstanceReportAttributeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DescribeInstanceReportAttributeResponseBody
	GetTaskId() *string
}

type DescribeInstanceReportAttributeResponseBody struct {
	// Indicates whether the report is the latest report.
	//
	// example:
	//
	// false
	Latest *bool `json:"Latest,omitempty" xml:"Latest,omitempty"`
	// The detailed description of the inspection report.
	//
	// example:
	//
	// No critical risks were found in this inspection
	ReportDescription *string `json:"ReportDescription,omitempty" xml:"ReportDescription,omitempty"`
	// The ID of the instance inspection report.
	//
	// example:
	//
	// report-001
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The overall score of the inspection report.
	//
	// example:
	//
	// 90
	ReportScore *int64 `json:"ReportScore,omitempty" xml:"ReportScore,omitempty"`
	// The generation status of the inspection report.
	//
	// example:
	//
	// Completed
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The inspection results for each item in the report.
	//
	// example:
	//
	// []
	ReportSubItems []*DescribeInstanceReportAttributeResponseBodyReportSubItems `json:"ReportSubItems,omitempty" xml:"ReportSubItems,omitempty" type:"Repeated"`
	// The overall summary of the inspection report.
	//
	// example:
	//
	// The instance is running normally
	ReportSummary *string `json:"ReportSummary,omitempty" xml:"ReportSummary,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the inspection task that generated the report.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInstanceReportAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceReportAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceReportAttributeResponseBody) GetLatest() *bool {
	return s.Latest
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportDescription() *string {
	return s.ReportDescription
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportScore() *int64 {
	return s.ReportScore
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportSubItems() []*DescribeInstanceReportAttributeResponseBodyReportSubItems {
	return s.ReportSubItems
}

func (s *DescribeInstanceReportAttributeResponseBody) GetReportSummary() *string {
	return s.ReportSummary
}

func (s *DescribeInstanceReportAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceReportAttributeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInstanceReportAttributeResponseBody) SetLatest(v bool) *DescribeInstanceReportAttributeResponseBody {
	s.Latest = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportDescription(v string) *DescribeInstanceReportAttributeResponseBody {
	s.ReportDescription = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportId(v string) *DescribeInstanceReportAttributeResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportScore(v int64) *DescribeInstanceReportAttributeResponseBody {
	s.ReportScore = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportStatus(v string) *DescribeInstanceReportAttributeResponseBody {
	s.ReportStatus = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportSubItems(v []*DescribeInstanceReportAttributeResponseBodyReportSubItems) *DescribeInstanceReportAttributeResponseBody {
	s.ReportSubItems = v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetReportSummary(v string) *DescribeInstanceReportAttributeResponseBody {
	s.ReportSummary = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetRequestId(v string) *DescribeInstanceReportAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) SetTaskId(v string) *DescribeInstanceReportAttributeResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBody) Validate() error {
	if s.ReportSubItems != nil {
		for _, item := range s.ReportSubItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceReportAttributeResponseBodyReportSubItems struct {
	// The risk level of the inspection item.
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The score of the inspection item.
	//
	// example:
	//
	// 90
	Score *int64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The name of the inspection skill that generated the inspection result.
	//
	// example:
	//
	// performance
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The result summary of the inspection item.
	//
	// example:
	//
	// No anomalies found
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s DescribeInstanceReportAttributeResponseBodyReportSubItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceReportAttributeResponseBodyReportSubItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) GetScore() *int64 {
	return s.Score
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) GetSummary() *string {
	return s.Summary
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) SetRiskLevel(v string) *DescribeInstanceReportAttributeResponseBodyReportSubItems {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) SetScore(v int64) *DescribeInstanceReportAttributeResponseBodyReportSubItems {
	s.Score = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) SetSkillName(v string) *DescribeInstanceReportAttributeResponseBodyReportSubItems {
	s.SkillName = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) SetSummary(v string) *DescribeInstanceReportAttributeResponseBodyReportSubItems {
	s.Summary = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponseBodyReportSubItems) Validate() error {
	return dara.Validate(s)
}
