// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInspectionTaskReportResponseBodyData) *DescribeInspectionTaskReportResponseBody
	GetData() *DescribeInspectionTaskReportResponseBodyData
	SetRequestId(v string) *DescribeInspectionTaskReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInspectionTaskReportResponseBody
	GetSuccess() *bool
}

type DescribeInspectionTaskReportResponseBody struct {
	Data *DescribeInspectionTaskReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInspectionTaskReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskReportResponseBody) GetData() *DescribeInspectionTaskReportResponseBodyData {
	return s.Data
}

func (s *DescribeInspectionTaskReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionTaskReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInspectionTaskReportResponseBody) SetData(v *DescribeInspectionTaskReportResponseBodyData) *DescribeInspectionTaskReportResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInspectionTaskReportResponseBody) SetRequestId(v string) *DescribeInspectionTaskReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBody) SetSuccess(v bool) *DescribeInspectionTaskReportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInspectionTaskReportResponseBodyData struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// # Tair 智能巡检报告\\n\\n## 总览...
	MarkdownText *string `json:"MarkdownText,omitempty" xml:"MarkdownText,omitempty"`
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status  *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Summary *DescribeInspectionTaskReportResponseBodyDataSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// example:
	//
	// tit-dca42f85c73644e0ab5c80ef6412xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTaskReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetMarkdownText() *string {
	return s.MarkdownText
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetSummary() *DescribeInspectionTaskReportResponseBodyDataSummary {
	return s.Summary
}

func (s *DescribeInspectionTaskReportResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetInstanceIds(v []*string) *DescribeInspectionTaskReportResponseBodyData {
	s.InstanceIds = v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetMarkdownText(v string) *DescribeInspectionTaskReportResponseBodyData {
	s.MarkdownText = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetReportLanguage(v string) *DescribeInspectionTaskReportResponseBodyData {
	s.ReportLanguage = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetStatus(v string) *DescribeInspectionTaskReportResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetSummary(v *DescribeInspectionTaskReportResponseBodyDataSummary) *DescribeInspectionTaskReportResponseBodyData {
	s.Summary = v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) SetTaskId(v string) *DescribeInspectionTaskReportResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyData) Validate() error {
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInspectionTaskReportResponseBodyDataSummary struct {
	// example:
	//
	// 0
	Error *int64 `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 0
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 1
	Normal *int64 `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// example:
	//
	// 0
	Warning *int64 `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s DescribeInspectionTaskReportResponseBodyDataSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskReportResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) GetError() *int64 {
	return s.Error
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) GetFailed() *int64 {
	return s.Failed
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) GetNormal() *int64 {
	return s.Normal
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) GetWarning() *int64 {
	return s.Warning
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) SetError(v int64) *DescribeInspectionTaskReportResponseBodyDataSummary {
	s.Error = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) SetFailed(v int64) *DescribeInspectionTaskReportResponseBodyDataSummary {
	s.Failed = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) SetNormal(v int64) *DescribeInspectionTaskReportResponseBodyDataSummary {
	s.Normal = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) SetWarning(v int64) *DescribeInspectionTaskReportResponseBodyDataSummary {
	s.Warning = &v
	return s
}

func (s *DescribeInspectionTaskReportResponseBodyDataSummary) Validate() error {
	return dara.Validate(s)
}
