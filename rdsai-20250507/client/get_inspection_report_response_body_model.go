// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetInspectionReportResponseBodyData) *GetInspectionReportResponseBody
	GetData() []*GetInspectionReportResponseBodyData
	SetMarkdownText(v string) *GetInspectionReportResponseBody
	GetMarkdownText() *string
	SetRequestId(v string) *GetInspectionReportResponseBody
	GetRequestId() *string
	SetTaskId(v string) *GetInspectionReportResponseBody
	GetTaskId() *string
	SetTemplateId(v string) *GetInspectionReportResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetInspectionReportResponseBody
	GetTemplateName() *string
}

type GetInspectionReportResponseBody struct {
	// The result details.
	Data []*GetInspectionReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The Markdown text.
	//
	// example:
	//
	// # RDS Batch Inspection Summary Report\\n\\n> This batch inspection checked **1*	- instance; found **1*	- instance with warnings (1 warning item in total)……
	MarkdownText *string `json:"MarkdownText,omitempty" xml:"MarkdownText,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The inspection report ID.
	//
	// example:
	//
	// 9d246af2-a0cd-4f69-857d-3785048f****
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetInspectionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBody) GetData() []*GetInspectionReportResponseBodyData {
	return s.Data
}

func (s *GetInspectionReportResponseBody) GetMarkdownText() *string {
	return s.MarkdownText
}

func (s *GetInspectionReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInspectionReportResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInspectionReportResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetInspectionReportResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetInspectionReportResponseBody) SetData(v []*GetInspectionReportResponseBodyData) *GetInspectionReportResponseBody {
	s.Data = v
	return s
}

func (s *GetInspectionReportResponseBody) SetMarkdownText(v string) *GetInspectionReportResponseBody {
	s.MarkdownText = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetRequestId(v string) *GetInspectionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetTaskId(v string) *GetInspectionReportResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetTemplateId(v string) *GetInspectionReportResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetInspectionReportResponseBody) SetTemplateName(v string) *GetInspectionReportResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetInspectionReportResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInspectionReportResponseBodyData struct {
	// The request result.
	Data []*GetInspectionReportResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The inspection end time in the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-01-31T02:05:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine type.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Test instance
	InstanceDesc *string `json:"InstanceDesc,omitempty" xml:"InstanceDesc,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2zep6e5u6l2yu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The level summary.
	LevelSummary *GetInspectionReportResponseBodyDataLevelSummary `json:"LevelSummary,omitempty" xml:"LevelSummary,omitempty" type:"Struct"`
	// The Markdown text.
	//
	// 	- If the InstanceId parameter is not specified: the reports for all instances in the inspection report are returned, but the MarkdownText field is empty ("").
	//
	// 	- If the InstanceId parameter is specified: the report for the specified instance is returned, and the MarkdownText field contains the specific content.
	//
	// example:
	//
	// # RDS Instance Inspection Report\\n\\n**Instance ID**: `rm-2zep6e5u6l2yu****`\\n\\n**Engine Type**: MySQL\\n\\n**Region**: cn-beijing\\n\\n**Time Range (UTC)**: 2026-02-04T03:31:00Z ~ 2026-02-05T03:31:00Z\\n\\n> This inspection includes **60*	- checks, of which 1 warning, 57 normal, and 2 failed to retrieve data.……
	MarkdownText *string `json:"MarkdownText,omitempty" xml:"MarkdownText,omitempty"`
	// The region information.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The inspection start time in the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInspectionReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBodyData) GetData() []*GetInspectionReportResponseBodyDataData {
	return s.Data
}

func (s *GetInspectionReportResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInspectionReportResponseBodyData) GetEngineType() *string {
	return s.EngineType
}

func (s *GetInspectionReportResponseBodyData) GetInstanceDesc() *string {
	return s.InstanceDesc
}

func (s *GetInspectionReportResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInspectionReportResponseBodyData) GetLevelSummary() *GetInspectionReportResponseBodyDataLevelSummary {
	return s.LevelSummary
}

func (s *GetInspectionReportResponseBodyData) GetMarkdownText() *string {
	return s.MarkdownText
}

func (s *GetInspectionReportResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetInspectionReportResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInspectionReportResponseBodyData) SetData(v []*GetInspectionReportResponseBodyDataData) *GetInspectionReportResponseBodyData {
	s.Data = v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetEndTime(v string) *GetInspectionReportResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetEngineType(v string) *GetInspectionReportResponseBodyData {
	s.EngineType = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetInstanceDesc(v string) *GetInspectionReportResponseBodyData {
	s.InstanceDesc = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetInstanceId(v string) *GetInspectionReportResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetLevelSummary(v *GetInspectionReportResponseBodyDataLevelSummary) *GetInspectionReportResponseBodyData {
	s.LevelSummary = v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetMarkdownText(v string) *GetInspectionReportResponseBodyData {
	s.MarkdownText = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetRegion(v string) *GetInspectionReportResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) SetStartTime(v string) *GetInspectionReportResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetInspectionReportResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LevelSummary != nil {
		if err := s.LevelSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInspectionReportResponseBodyDataData struct {
	// The group ID.
	//
	// example:
	//
	// instance_info
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The attached resource names.
	Items []*GetInspectionReportResponseBodyDataDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetInspectionReportResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBodyDataData) GetGroup() *string {
	return s.Group
}

func (s *GetInspectionReportResponseBodyDataData) GetItems() []*GetInspectionReportResponseBodyDataDataItems {
	return s.Items
}

func (s *GetInspectionReportResponseBodyDataData) SetGroup(v string) *GetInspectionReportResponseBodyDataData {
	s.Group = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataData) SetItems(v []*GetInspectionReportResponseBodyDataDataItems) *GetInspectionReportResponseBodyDataData {
	s.Items = v
	return s
}

func (s *GetInspectionReportResponseBodyDataData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInspectionReportResponseBodyDataDataItems struct {
	// The request result.
	Data []*GetInspectionReportResponseBodyDataDataItemsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The alert level.
	//
	// example:
	//
	// Normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The result message.
	//
	// example:
	//
	// Instance running status is normal
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The category name.
	//
	// example:
	//
	// instance_runningstatus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInspectionReportResponseBodyDataDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBodyDataDataItems) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBodyDataDataItems) GetData() []*GetInspectionReportResponseBodyDataDataItemsData {
	return s.Data
}

func (s *GetInspectionReportResponseBodyDataDataItems) GetLevel() *string {
	return s.Level
}

func (s *GetInspectionReportResponseBodyDataDataItems) GetMessage() *string {
	return s.Message
}

func (s *GetInspectionReportResponseBodyDataDataItems) GetName() *string {
	return s.Name
}

func (s *GetInspectionReportResponseBodyDataDataItems) SetData(v []*GetInspectionReportResponseBodyDataDataItemsData) *GetInspectionReportResponseBodyDataDataItems {
	s.Data = v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItems) SetLevel(v string) *GetInspectionReportResponseBodyDataDataItems {
	s.Level = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItems) SetMessage(v string) *GetInspectionReportResponseBodyDataDataItems {
	s.Message = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItems) SetName(v string) *GetInspectionReportResponseBodyDataDataItems {
	s.Name = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItems) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInspectionReportResponseBodyDataDataItemsData struct {
	// The tag key.
	//
	// example:
	//
	// DBInstanceStatus
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Running
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInspectionReportResponseBodyDataDataItemsData) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBodyDataDataItemsData) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBodyDataDataItemsData) GetKey() *string {
	return s.Key
}

func (s *GetInspectionReportResponseBodyDataDataItemsData) GetValue() *string {
	return s.Value
}

func (s *GetInspectionReportResponseBodyDataDataItemsData) SetKey(v string) *GetInspectionReportResponseBodyDataDataItemsData {
	s.Key = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItemsData) SetValue(v string) *GetInspectionReportResponseBodyDataDataItemsData {
	s.Value = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataDataItemsData) Validate() error {
	return dara.Validate(s)
}

type GetInspectionReportResponseBodyDataLevelSummary struct {
	// The number of error items.
	//
	// example:
	//
	// 2
	Error *int64 `json:"Error,omitempty" xml:"Error,omitempty"`
	// The number of failed items.
	//
	// example:
	//
	// 3
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of normal items.
	//
	// example:
	//
	// 10
	Normal *int64 `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// The number of warning items.
	//
	// example:
	//
	// 1
	Warning *int64 `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s GetInspectionReportResponseBodyDataLevelSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponseBodyDataLevelSummary) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) GetError() *int64 {
	return s.Error
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) GetNormal() *int64 {
	return s.Normal
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) GetWarning() *int64 {
	return s.Warning
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) SetError(v int64) *GetInspectionReportResponseBodyDataLevelSummary {
	s.Error = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) SetFailed(v int64) *GetInspectionReportResponseBodyDataLevelSummary {
	s.Failed = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) SetNormal(v int64) *GetInspectionReportResponseBodyDataLevelSummary {
	s.Normal = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) SetWarning(v int64) *GetInspectionReportResponseBodyDataLevelSummary {
	s.Warning = &v
	return s
}

func (s *GetInspectionReportResponseBodyDataLevelSummary) Validate() error {
	return dara.Validate(s)
}
