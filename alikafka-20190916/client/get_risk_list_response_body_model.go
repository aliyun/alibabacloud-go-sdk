// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetRiskListResponseBody
	GetCode() *int64
	SetData(v *GetRiskListResponseBodyData) *GetRiskListResponseBody
	GetData() *GetRiskListResponseBodyData
	SetMessage(v string) *GetRiskListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRiskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRiskListResponseBody
	GetSuccess() *bool
}

type GetRiskListResponseBody struct {
	// The return code. A value of 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRiskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRiskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRiskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetRiskListResponseBody) GetData() *GetRiskListResponseBodyData {
	return s.Data
}

func (s *GetRiskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRiskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRiskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRiskListResponseBody) SetCode(v int64) *GetRiskListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRiskListResponseBody) SetData(v *GetRiskListResponseBodyData) *GetRiskListResponseBody {
	s.Data = v
	return s
}

func (s *GetRiskListResponseBody) SetMessage(v string) *GetRiskListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRiskListResponseBody) SetRequestId(v string) *GetRiskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskListResponseBody) SetSuccess(v bool) *GetRiskListResponseBody {
	s.Success = &v
	return s
}

func (s *GetRiskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRiskListResponseBodyData struct {
	// The list of threat items for the instance.
	RiskList []*GetRiskListResponseBodyDataRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 11
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRiskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRiskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRiskListResponseBodyData) GetRiskList() []*GetRiskListResponseBodyDataRiskList {
	return s.RiskList
}

func (s *GetRiskListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetRiskListResponseBodyData) SetRiskList(v []*GetRiskListResponseBodyDataRiskList) *GetRiskListResponseBodyData {
	s.RiskList = v
	return s
}

func (s *GetRiskListResponseBodyData) SetTotal(v int64) *GetRiskListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRiskListResponseBodyData) Validate() error {
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRiskListResponseBodyDataRiskList struct {
	// The timestamp when the threat was created. Unit: milliseconds.
	//
	// example:
	//
	// 1702545932000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The metric rating. Valid values:
	//
	// - A: Healthy.
	//
	// - B: Suboptimal.
	//
	// - F: Poor.
	//
	// example:
	//
	// A
	GradeType *string `json:"GradeType,omitempty" xml:"GradeType,omitempty"`
	// Indicates whether the instance is healthy.
	//
	// This is a Boolean parameter. Valid values:
	//
	// - true: The instance is healthy.
	//
	// - false: The instance is unhealthy.
	//
	// example:
	//
	// true
	Health *bool `json:"Health,omitempty" xml:"Health,omitempty"`
	// The list of instance IDs.
	//
	// example:
	//
	// alikafka_pre-cn-m7r1tzxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp of the last alert. Unit: milliseconds.
	//
	// example:
	//
	// 1683270264
	LastAlarmTime *int64 `json:"LastAlarmTime,omitempty" xml:"LastAlarmTime,omitempty"`
	// The risk level. Valid values:
	//
	// - 0: Urgent.
	//
	// - 1: Important.
	//
	// - 2: Normal.
	//
	// example:
	//
	// 1
	LevelType *int64 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The timestamp when the threat was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1637719920000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the threat item.
	//
	// > There are 24 types of names.
	//
	// >
	//
	// > - For more information, see the supplementary notes at the end of this document.
	//
	// example:
	//
	// inputIo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner.
	//
	// example:
	//
	// 0123123123xxx
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// A cascading structure. The system determines whether to nest another layer of report data based on the values of outer fields.
	RelationList []*string `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	// The recommended fix.
	//
	// example:
	//
	// 相关问题里的Topic存在碎片化发送问题，请参考文档进行优化
	ReportTips *string `json:"ReportTips,omitempty" xml:"ReportTips,omitempty"`
	// The report type of the threat item. Valid values:
	//
	// - topic: Optimization is required for a specific topic.
	//
	// - group: Optimization is required for a specific group.
	//
	// - doc: Optimization must be performed based on a document.
	//
	// - commonBuy: An upgrade or a similar operation is required for the returned threat item.
	//
	// - mdsKey: You only need to fix the threat based on the suggestions in ReportTips.
	//
	// example:
	//
	// doc
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The value of the report.
	//
	// > - If ReportType is doc, ReportValue returns the document ID. You can construct the URL to the document by replacing the ${reportValue} variable in the following URL with the returned value: <props="china">https\\://help.aliyun.com/document_detail/${reportValue}.html<props="intl">https\\://www\\.alibabacloud.com/help/document_detail/${reportValue}.html
	//
	// >
	//
	// > - If ReportType is commonBuy, an upgrade or a similar operation is required.
	//
	// >
	//
	// > - If ReportType is topic, the value of ReportValue is the name of the topic that needs to be fixed.
	//
	// >
	//
	// > - If ReportType is group, the value of ReportValue is the name of the group that needs to be fixed.
	//
	// >
	//
	// > - If ReportType is mdsKey, you only need to fix the threat based on the suggestions in ReportTips.
	//
	// example:
	//
	// test
	ReportValue *string `json:"ReportValue,omitempty" xml:"ReportValue,omitempty"`
	// The status of the threat item. This parameter indicates whether the threat has been fixed. Valid values:
	//
	// - 0: To be fixed.
	//
	// - -1: Ignored.
	//
	// - 1: Fixed.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the threat item.
	//
	// > There are 24 types of threats.
	//
	// >
	//
	// > - For more information, see the supplementary notes at the end of this document.
	//
	// example:
	//
	// inputIo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value calculated by the system.
	//
	// > If ReportType is doc, check the relationList and value fields. The value field returns a number that indicates the number of topics or groups in the `relationList` field that require optimization.
	//
	// >
	//
	// > - When ReportType is commonBuy, check the value of Value. The value is a percentage.
	//
	// >
	//
	// > - When ReportType is topic, check the value of Value. The value identifies the Topic that needs to be fixed.
	//
	// >
	//
	// > - When ReportType is group, check the value of Value. The value identifies the Group that needs to be fixed.
	//
	// example:
	//
	// 44
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRiskListResponseBodyDataRiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRiskListResponseBodyDataRiskList) GoString() string {
	return s.String()
}

func (s *GetRiskListResponseBodyDataRiskList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRiskListResponseBodyDataRiskList) GetGradeType() *string {
	return s.GradeType
}

func (s *GetRiskListResponseBodyDataRiskList) GetHealth() *bool {
	return s.Health
}

func (s *GetRiskListResponseBodyDataRiskList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRiskListResponseBodyDataRiskList) GetLastAlarmTime() *int64 {
	return s.LastAlarmTime
}

func (s *GetRiskListResponseBodyDataRiskList) GetLevelType() *int64 {
	return s.LevelType
}

func (s *GetRiskListResponseBodyDataRiskList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetRiskListResponseBodyDataRiskList) GetName() *string {
	return s.Name
}

func (s *GetRiskListResponseBodyDataRiskList) GetOwner() *string {
	return s.Owner
}

func (s *GetRiskListResponseBodyDataRiskList) GetRelationList() []*string {
	return s.RelationList
}

func (s *GetRiskListResponseBodyDataRiskList) GetReportTips() *string {
	return s.ReportTips
}

func (s *GetRiskListResponseBodyDataRiskList) GetReportType() *string {
	return s.ReportType
}

func (s *GetRiskListResponseBodyDataRiskList) GetReportValue() *string {
	return s.ReportValue
}

func (s *GetRiskListResponseBodyDataRiskList) GetStatus() *int64 {
	return s.Status
}

func (s *GetRiskListResponseBodyDataRiskList) GetType() *string {
	return s.Type
}

func (s *GetRiskListResponseBodyDataRiskList) GetValue() *string {
	return s.Value
}

func (s *GetRiskListResponseBodyDataRiskList) SetCreateTime(v int64) *GetRiskListResponseBodyDataRiskList {
	s.CreateTime = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetGradeType(v string) *GetRiskListResponseBodyDataRiskList {
	s.GradeType = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetHealth(v bool) *GetRiskListResponseBodyDataRiskList {
	s.Health = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetInstanceId(v string) *GetRiskListResponseBodyDataRiskList {
	s.InstanceId = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetLastAlarmTime(v int64) *GetRiskListResponseBodyDataRiskList {
	s.LastAlarmTime = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetLevelType(v int64) *GetRiskListResponseBodyDataRiskList {
	s.LevelType = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetModifiedTime(v int64) *GetRiskListResponseBodyDataRiskList {
	s.ModifiedTime = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetName(v string) *GetRiskListResponseBodyDataRiskList {
	s.Name = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetOwner(v string) *GetRiskListResponseBodyDataRiskList {
	s.Owner = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetRelationList(v []*string) *GetRiskListResponseBodyDataRiskList {
	s.RelationList = v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetReportTips(v string) *GetRiskListResponseBodyDataRiskList {
	s.ReportTips = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetReportType(v string) *GetRiskListResponseBodyDataRiskList {
	s.ReportType = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetReportValue(v string) *GetRiskListResponseBodyDataRiskList {
	s.ReportValue = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetStatus(v int64) *GetRiskListResponseBodyDataRiskList {
	s.Status = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetType(v string) *GetRiskListResponseBodyDataRiskList {
	s.Type = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) SetValue(v string) *GetRiskListResponseBodyDataRiskList {
	s.Value = &v
	return s
}

func (s *GetRiskListResponseBodyDataRiskList) Validate() error {
	return dara.Validate(s)
}
