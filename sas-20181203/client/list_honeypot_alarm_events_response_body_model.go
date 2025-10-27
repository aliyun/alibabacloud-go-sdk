// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAlarmEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotAlarmEvents(v []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) *ListHoneypotAlarmEventsResponseBody
	GetHoneypotAlarmEvents() []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents
	SetPageInfo(v *ListHoneypotAlarmEventsResponseBodyPageInfo) *ListHoneypotAlarmEventsResponseBody
	GetPageInfo() *ListHoneypotAlarmEventsResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotAlarmEventsResponseBody
	GetRequestId() *string
}

type ListHoneypotAlarmEventsResponseBody struct {
	// The alert events.
	HoneypotAlarmEvents []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents `json:"HoneypotAlarmEvents,omitempty" xml:"HoneypotAlarmEvents,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListHoneypotAlarmEventsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 513C9554-55A4-5504-B7C4-6E17EB4FC7A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHoneypotAlarmEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsResponseBody) GetHoneypotAlarmEvents() []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	return s.HoneypotAlarmEvents
}

func (s *ListHoneypotAlarmEventsResponseBody) GetPageInfo() *ListHoneypotAlarmEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotAlarmEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotAlarmEventsResponseBody) SetHoneypotAlarmEvents(v []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) *ListHoneypotAlarmEventsResponseBody {
	s.HoneypotAlarmEvents = v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBody) SetPageInfo(v *ListHoneypotAlarmEventsResponseBodyPageInfo) *ListHoneypotAlarmEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBody) SetRequestId(v string) *ListHoneypotAlarmEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBody) Validate() error {
	if s.HoneypotAlarmEvents != nil {
		for _, item := range s.HoneypotAlarmEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents struct {
	// The event ID.
	//
	// example:
	//
	// 940272
	AlarmEventId *int64 `json:"AlarmEventId,omitempty" xml:"AlarmEventId,omitempty"`
	// The name of the alert event.
	//
	// example:
	//
	// Attack Honeypot
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	// The type of the alert event.
	//
	// example:
	//
	// Initial Access
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// The unique identifier of the alert event.
	//
	// example:
	//
	// 167e6fc0d931917d2059efcd1d00f6ab
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The total number of times that the alert event was generated.
	//
	// example:
	//
	// 11
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The timestamp that indicates the time when the alert event was first detected. Unit: milliseconds.
	//
	// example:
	//
	// 1658193602000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The timestamp that indicates the time when the alert event was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1660610772000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The risk information.
	MergeFieldList []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList `json:"MergeFieldList,omitempty" xml:"MergeFieldList,omitempty" type:"Repeated"`
	// The handling status of the alert event. Valid values:
	//
	// 	- **1**: pending
	//
	// 	- **2**: ignored
	//
	// 	- **4**: confirmed
	//
	// example:
	//
	// 1
	OperateStatus *int32 `json:"OperateStatus,omitempty" xml:"OperateStatus,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **2**: low
	//
	// 	- **3**: medium
	//
	// 	- **4**: high
	//
	// example:
	//
	// 2
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetAlarmEventId() *int64 {
	return s.AlarmEventId
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetAlarmEventName() *string {
	return s.AlarmEventName
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetEventCount() *int32 {
	return s.EventCount
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetMergeFieldList() []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList {
	return s.MergeFieldList
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetOperateStatus() *int32 {
	return s.OperateStatus
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetAlarmEventId(v int64) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.AlarmEventId = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetAlarmEventName(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.AlarmEventName = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetAlarmEventType(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.AlarmEventType = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetAlarmUniqueInfo(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetEventCount(v int32) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.EventCount = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetFirstTime(v int64) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.FirstTime = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetLastTime(v int64) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.LastTime = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetMergeFieldList(v []*ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.MergeFieldList = v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetOperateStatus(v int32) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.OperateStatus = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) SetRiskLevel(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents {
	s.RiskLevel = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEvents) Validate() error {
	if s.MergeFieldList != nil {
		for _, item := range s.MergeFieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList struct {
	// The extended value that corresponds to the field key.
	//
	// example:
	//
	// dest_ip_ext
	FieldExtInfo *string `json:"FieldExtInfo,omitempty" xml:"FieldExtInfo,omitempty"`
	// The key of the field.
	//
	// example:
	//
	// dest_ip_count
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	// The type of the field. You can ignore this internal parameter.
	//
	// example:
	//
	// level1_item3
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The value that corresponds to the field key.
	//
	// example:
	//
	// 1
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) GetFieldExtInfo() *string {
	return s.FieldExtInfo
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) GetFieldKey() *string {
	return s.FieldKey
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) GetFieldType() *string {
	return s.FieldType
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) SetFieldExtInfo(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList {
	s.FieldExtInfo = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) SetFieldKey(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList {
	s.FieldKey = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) SetFieldType(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList {
	s.FieldType = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) SetFieldValue(v string) *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList {
	s.FieldValue = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyHoneypotAlarmEventsMergeFieldList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotAlarmEventsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotAlarmEventsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) SetCount(v int32) *ListHoneypotAlarmEventsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotAlarmEventsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotAlarmEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotAlarmEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotAlarmEventsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
