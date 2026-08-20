// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountSafetyIncidentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountSafetyIncidentResponseBody
	GetCode() *string
	SetData(v *QueryAccountSafetyIncidentResponseBodyData) *QueryAccountSafetyIncidentResponseBody
	GetData() *QueryAccountSafetyIncidentResponseBodyData
	SetMessage(v string) *QueryAccountSafetyIncidentResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountSafetyIncidentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountSafetyIncidentResponseBody
	GetSuccess() *bool
}

type QueryAccountSafetyIncidentResponseBody struct {
	// The status code.
	//
	// > 200: success. Other values (such as 500 or 400): error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryAccountSafetyIncidentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The prompt message.
	//
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FBDD713-00A5-5C98-B661-3FD31A349B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountSafetyIncidentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountSafetyIncidentResponseBody) GetData() *QueryAccountSafetyIncidentResponseBodyData {
	return s.Data
}

func (s *QueryAccountSafetyIncidentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountSafetyIncidentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountSafetyIncidentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountSafetyIncidentResponseBody) SetCode(v string) *QueryAccountSafetyIncidentResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBody) SetData(v *QueryAccountSafetyIncidentResponseBodyData) *QueryAccountSafetyIncidentResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBody) SetMessage(v string) *QueryAccountSafetyIncidentResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBody) SetRequestId(v string) *QueryAccountSafetyIncidentResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBody) SetSuccess(v bool) *QueryAccountSafetyIncidentResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountSafetyIncidentResponseBodyData struct {
	// The event data.
	List []*QueryAccountSafetyIncidentResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *QueryAccountSafetyIncidentResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s QueryAccountSafetyIncidentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponseBodyData) GetList() []*QueryAccountSafetyIncidentResponseBodyDataList {
	return s.List
}

func (s *QueryAccountSafetyIncidentResponseBodyData) GetPageInfo() *QueryAccountSafetyIncidentResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *QueryAccountSafetyIncidentResponseBodyData) SetList(v []*QueryAccountSafetyIncidentResponseBodyDataList) *QueryAccountSafetyIncidentResponseBodyData {
	s.List = v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyData) SetPageInfo(v *QueryAccountSafetyIncidentResponseBodyDataPageInfo) *QueryAccountSafetyIncidentResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type QueryAccountSafetyIncidentResponseBodyDataList struct {
	// The control action name code.
	//
	// example:
	//
	// success_service
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The control action name.
	//
	// example:
	//
	// 处罚直接成功
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The control removal time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	AntiPunishTime *string `json:"AntiPunishTime,omitempty" xml:"AntiPunishTime,omitempty"`
	// The called API operation.
	//
	// example:
	//
	// AddDomainRecord
	CallApi *string `json:"CallApi,omitempty" xml:"CallApi,omitempty"`
	// The control time information.
	DateExtras *QueryAccountSafetyIncidentResponseBodyDataListDateExtras `json:"DateExtras,omitempty" xml:"DateExtras,omitempty" type:"Struct"`
	// The event ID.
	//
	// example:
	//
	// 4ba4065e0b2206c05f86d5eaa00ae520
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event impact.
	//
	// example:
	//
	// ak leak.
	EventImpact *string `json:"EventImpact,omitempty" xml:"EventImpact,omitempty"`
	// The control event name.
	//
	// example:
	//
	// ak leak.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The event reason.
	//
	// example:
	//
	// ak leak.
	EventReason *string `json:"EventReason,omitempty" xml:"EventReason,omitempty"`
	// The event subtype name.
	//
	// example:
	//
	// 可疑身份调用敏感
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The exception call time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	ExceptionCallTime *string `json:"ExceptionCallTime,omitempty" xml:"ExceptionCallTime,omitempty"`
	// The exception IP address.
	//
	// example:
	//
	// 39.1X4.63.XX9
	ExceptionIp *string `json:"ExceptionIp,omitempty" xml:"ExceptionIp,omitempty"`
	// The control start time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	PunishTime *string `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	// The hardening suggestion.
	//
	// example:
	//
	// suggestion
	Reinforcement *string `json:"Reinforcement,omitempty" xml:"Reinforcement,omitempty"`
	// The cloud resource ID.
	//
	// example:
	//
	// i-2zeanc2b2vgfpbvp60cs
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The control object type.
	//
	// example:
	//
	// customer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The event status. Valid values:
	//
	// - **Executing**: In progress.
	//
	// - **Removed**: Removed.
	//
	// - **Alerting**: Alerting.
	//
	// - **Ended**: Ended.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The handling suggestion.
	//
	// example:
	//
	// suggestion
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The help topic name.
	//
	// example:
	//
	// help
	UserGuideName *string `json:"UserGuideName,omitempty" xml:"UserGuideName,omitempty"`
	// The help topic URL.
	//
	// example:
	//
	// https://xxx.aliyun.com/
	UserGuideUrl *string `json:"UserGuideUrl,omitempty" xml:"UserGuideUrl,omitempty"`
}

func (s QueryAccountSafetyIncidentResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetActionName() *string {
	return s.ActionName
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetAntiPunishTime() *string {
	return s.AntiPunishTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetCallApi() *string {
	return s.CallApi
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetDateExtras() *QueryAccountSafetyIncidentResponseBodyDataListDateExtras {
	return s.DateExtras
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetEventId() *string {
	return s.EventId
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetEventImpact() *string {
	return s.EventImpact
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetEventName() *string {
	return s.EventName
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetEventReason() *string {
	return s.EventReason
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetEventType() *string {
	return s.EventType
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetExceptionCallTime() *string {
	return s.ExceptionCallTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetExceptionIp() *string {
	return s.ExceptionIp
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetPunishTime() *string {
	return s.PunishTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetReinforcement() *string {
	return s.Reinforcement
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetTip() *string {
	return s.Tip
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetUserGuideName() *string {
	return s.UserGuideName
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) GetUserGuideUrl() *string {
	return s.UserGuideUrl
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetActionCode(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ActionCode = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetActionName(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ActionName = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetAntiPunishTime(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.AntiPunishTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetCallApi(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.CallApi = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetDateExtras(v *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.DateExtras = v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetEventId(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetEventImpact(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.EventImpact = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetEventName(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetEventReason(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.EventReason = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetEventType(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetExceptionCallTime(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ExceptionCallTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetExceptionIp(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ExceptionIp = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetPunishTime(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.PunishTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetReinforcement(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.Reinforcement = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetResourceId(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ResourceId = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetResourceType(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetStatus(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetTip(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.Tip = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetUserGuideName(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.UserGuideName = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) SetUserGuideUrl(v string) *QueryAccountSafetyIncidentResponseBodyDataList {
	s.UserGuideUrl = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataList) Validate() error {
	if s.DateExtras != nil {
		if err := s.DateExtras.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountSafetyIncidentResponseBodyDataListDateExtras struct {
	// The alert end time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	AlertEndTime *string `json:"AlertEndTime,omitempty" xml:"AlertEndTime,omitempty"`
	// The first alert time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	AlertStartTime *string `json:"AlertStartTime,omitempty" xml:"AlertStartTime,omitempty"`
	// The latest detection time.
	//
	// > Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-16 15:15:00
	LastCheckTime *string `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
}

func (s QueryAccountSafetyIncidentResponseBodyDataListDateExtras) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponseBodyDataListDateExtras) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) GetAlertEndTime() *string {
	return s.AlertEndTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) GetAlertStartTime() *string {
	return s.AlertStartTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) GetLastCheckTime() *string {
	return s.LastCheckTime
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) SetAlertEndTime(v string) *QueryAccountSafetyIncidentResponseBodyDataListDateExtras {
	s.AlertEndTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) SetAlertStartTime(v string) *QueryAccountSafetyIncidentResponseBodyDataListDateExtras {
	s.AlertStartTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) SetLastCheckTime(v string) *QueryAccountSafetyIncidentResponseBodyDataListDateExtras {
	s.LastCheckTime = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataListDateExtras) Validate() error {
	return dara.Validate(s)
}

type QueryAccountSafetyIncidentResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The number of assets displayed on each page in a paging query.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of events.
	//
	// example:
	//
	// 20
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAccountSafetyIncidentResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountSafetyIncidentResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) GetCurrent() *string {
	return s.Current
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) GetTotal() *string {
	return s.Total
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) SetCurrent(v string) *QueryAccountSafetyIncidentResponseBodyDataPageInfo {
	s.Current = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) SetPageSize(v string) *QueryAccountSafetyIncidentResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) SetTotal(v string) *QueryAccountSafetyIncidentResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

func (s *QueryAccountSafetyIncidentResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
