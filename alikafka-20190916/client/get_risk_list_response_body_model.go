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
	// example:
	//
	// 200
	Code *int64                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRiskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetRiskListResponseBodyData struct {
	RiskList []*GetRiskListResponseBodyDataRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type GetRiskListResponseBodyDataRiskList struct {
	// example:
	//
	// 1702545932000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// A
	GradeType *string `json:"GradeType,omitempty" xml:"GradeType,omitempty"`
	// example:
	//
	// true
	Health *bool `json:"Health,omitempty" xml:"Health,omitempty"`
	// example:
	//
	// alikafka_pre-cn-m7r1tzxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1683270264
	LastAlarmTime *int64 `json:"LastAlarmTime,omitempty" xml:"LastAlarmTime,omitempty"`
	// example:
	//
	// 1
	LevelType *int64 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// example:
	//
	// 1637719920000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// inputIo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0123123123xxx
	Owner        *string   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RelationList []*string `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	ReportTips   *string   `json:"ReportTips,omitempty" xml:"ReportTips,omitempty"`
	// example:
	//
	// doc
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// test
	ReportValue *string `json:"ReportValue,omitempty" xml:"ReportValue,omitempty"`
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// inputIo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
