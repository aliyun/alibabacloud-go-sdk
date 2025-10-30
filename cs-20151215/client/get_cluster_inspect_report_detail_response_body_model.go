// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterInspectReportDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckItemResults(v []*GetClusterInspectReportDetailResponseBodyCheckItemResults) *GetClusterInspectReportDetailResponseBody
	GetCheckItemResults() []*GetClusterInspectReportDetailResponseBodyCheckItemResults
	SetEndTime(v string) *GetClusterInspectReportDetailResponseBody
	GetEndTime() *string
	SetNextToken(v string) *GetClusterInspectReportDetailResponseBody
	GetNextToken() *string
	SetReportId(v string) *GetClusterInspectReportDetailResponseBody
	GetReportId() *string
	SetRequestId(v string) *GetClusterInspectReportDetailResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetClusterInspectReportDetailResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetClusterInspectReportDetailResponseBody
	GetStatus() *string
	SetSummary(v *GetClusterInspectReportDetailResponseBodySummary) *GetClusterInspectReportDetailResponseBody
	GetSummary() *GetClusterInspectReportDetailResponseBodySummary
}

type GetClusterInspectReportDetailResponseBody struct {
	// The results.
	CheckItemResults []*GetClusterInspectReportDetailResponseBodyCheckItemResults `json:"checkItemResults,omitempty" xml:"checkItemResults,omitempty" type:"Repeated"`
	// The completion time of the inspection report.
	//
	// example:
	//
	// 2024-12-18T19:41:12.778433+08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The token that is used to display the returned tags on multiple pages.
	//
	// example:
	//
	// AK8uQQrxgFK8sbARvnCj6w9R3kPme4I3
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the inspection report.
	//
	// example:
	//
	// 782df89346054a0000562063a****
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The start time of the inspection report.
	//
	// example:
	//
	// 2024-12-18T19:40:16.778333+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the inspection report. Valid values:
	//
	// 	- completed: The inspection report is generated.
	//
	// 	- running: The inspection report is generating.
	//
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Overview of inspection reports.
	Summary *GetClusterInspectReportDetailResponseBodySummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
}

func (s GetClusterInspectReportDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterInspectReportDetailResponseBody) GetCheckItemResults() []*GetClusterInspectReportDetailResponseBodyCheckItemResults {
	return s.CheckItemResults
}

func (s *GetClusterInspectReportDetailResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetClusterInspectReportDetailResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetClusterInspectReportDetailResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GetClusterInspectReportDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterInspectReportDetailResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetClusterInspectReportDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetClusterInspectReportDetailResponseBody) GetSummary() *GetClusterInspectReportDetailResponseBodySummary {
	return s.Summary
}

func (s *GetClusterInspectReportDetailResponseBody) SetCheckItemResults(v []*GetClusterInspectReportDetailResponseBodyCheckItemResults) *GetClusterInspectReportDetailResponseBody {
	s.CheckItemResults = v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetEndTime(v string) *GetClusterInspectReportDetailResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetNextToken(v string) *GetClusterInspectReportDetailResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetReportId(v string) *GetClusterInspectReportDetailResponseBody {
	s.ReportId = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetRequestId(v string) *GetClusterInspectReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetStartTime(v string) *GetClusterInspectReportDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetStatus(v string) *GetClusterInspectReportDetailResponseBody {
	s.Status = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) SetSummary(v *GetClusterInspectReportDetailResponseBodySummary) *GetClusterInspectReportDetailResponseBody {
	s.Summary = v
	return s
}

func (s *GetClusterInspectReportDetailResponseBody) Validate() error {
	if s.CheckItemResults != nil {
		for _, item := range s.CheckItemResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterInspectReportDetailResponseBodyCheckItemResults struct {
	// The category of the inspection item. Valid values:
	//
	// 	- security: Security compliance
	//
	// 	- performance: Performance efficiency
	//
	// 	- stability: Business stability
	//
	// 	- limitation: Service limits
	//
	// 	- cost: Cost optimization
	//
	// example:
	//
	// stability
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The unique identifier of the inspection item.
	//
	// example:
	//
	// APIServerClbInstanceStatus
	CheckItemUid *string `json:"checkItemUid,omitempty" xml:"checkItemUid,omitempty"`
	// The description of the inspection item.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The fixing suggestion.
	Fix *string `json:"fix,omitempty" xml:"fix,omitempty"`
	// The level of the inspection item. Valid values:
	//
	// 	- advice: Suggestions
	//
	// 	- warning: Low severity
	//
	// 	- error: Medium severity
	//
	// 	- critical: High severity
	//
	// example:
	//
	// critical
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The name of the inspection item.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The inspection results. Valid values:
	//
	// 	- true: The inspection item is abnormal.
	//
	// 	- false: The inspection item is normal.
	//
	// 	- disable: The inspection item is not enabled.
	//
	// example:
	//
	// false
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The resource type of the inspection object.
	//
	// example:
	//
	// CLB
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The inspection objects.
	Targets []*string `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s GetClusterInspectReportDetailResponseBodyCheckItemResults) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectReportDetailResponseBodyCheckItemResults) GoString() string {
	return s.String()
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetCategory() *string {
	return s.Category
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetCheckItemUid() *string {
	return s.CheckItemUid
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetDescription() *string {
	return s.Description
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetFix() *string {
	return s.Fix
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetLevel() *string {
	return s.Level
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetName() *string {
	return s.Name
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetResult() *string {
	return s.Result
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetTargetType() *string {
	return s.TargetType
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) GetTargets() []*string {
	return s.Targets
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetCategory(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Category = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetCheckItemUid(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.CheckItemUid = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetDescription(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Description = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetFix(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Fix = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetLevel(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Level = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetName(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Name = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetResult(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Result = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetTargetType(v string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.TargetType = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) SetTargets(v []*string) *GetClusterInspectReportDetailResponseBodyCheckItemResults {
	s.Targets = v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodyCheckItemResults) Validate() error {
	return dara.Validate(s)
}

type GetClusterInspectReportDetailResponseBodySummary struct {
	// The number of check items whose inspection result is advice.
	//
	// example:
	//
	// 0
	AdviceCount *int32 `json:"adviceCount,omitempty" xml:"adviceCount,omitempty"`
	// Check the status code of the inspection task.
	//
	// example:
	//
	// warning
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The number of check items whose inspection result is error.
	//
	// example:
	//
	// 0
	ErrorCount *int32 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// The number of check items whose inspection result is normal.
	//
	// example:
	//
	// 10
	NormalCount *int32 `json:"normalCount,omitempty" xml:"normalCount,omitempty"`
	// The number of check items whose inspection result is warning.
	//
	// example:
	//
	// 1
	WarnCount *int32 `json:"warnCount,omitempty" xml:"warnCount,omitempty"`
}

func (s GetClusterInspectReportDetailResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectReportDetailResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetClusterInspectReportDetailResponseBodySummary) GetAdviceCount() *int32 {
	return s.AdviceCount
}

func (s *GetClusterInspectReportDetailResponseBodySummary) GetCode() *string {
	return s.Code
}

func (s *GetClusterInspectReportDetailResponseBodySummary) GetErrorCount() *int32 {
	return s.ErrorCount
}

func (s *GetClusterInspectReportDetailResponseBodySummary) GetNormalCount() *int32 {
	return s.NormalCount
}

func (s *GetClusterInspectReportDetailResponseBodySummary) GetWarnCount() *int32 {
	return s.WarnCount
}

func (s *GetClusterInspectReportDetailResponseBodySummary) SetAdviceCount(v int32) *GetClusterInspectReportDetailResponseBodySummary {
	s.AdviceCount = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodySummary) SetCode(v string) *GetClusterInspectReportDetailResponseBodySummary {
	s.Code = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodySummary) SetErrorCount(v int32) *GetClusterInspectReportDetailResponseBodySummary {
	s.ErrorCount = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodySummary) SetNormalCount(v int32) *GetClusterInspectReportDetailResponseBodySummary {
	s.NormalCount = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodySummary) SetWarnCount(v int32) *GetClusterInspectReportDetailResponseBodySummary {
	s.WarnCount = &v
	return s
}

func (s *GetClusterInspectReportDetailResponseBodySummary) Validate() error {
	return dara.Validate(s)
}
