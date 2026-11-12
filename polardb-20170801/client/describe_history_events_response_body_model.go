// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeHistoryEventsResponseBodyItems) *DescribeHistoryEventsResponseBody
	GetItems() []*DescribeHistoryEventsResponseBodyItems
	SetPageNumber(v int32) *DescribeHistoryEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHistoryEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHistoryEventsResponseBody
	GetTotalCount() *int32
}

type DescribeHistoryEventsResponseBody struct {
	// The list of events.
	//
	// example:
	//
	// []
	Items []*DescribeHistoryEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponseBody) GetItems() []*DescribeHistoryEventsResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryEventsResponseBody) SetItems(v []*DescribeHistoryEventsResponseBodyItems) *DescribeHistoryEventsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetPageNumber(v int32) *DescribeHistoryEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetPageSize(v int32) *DescribeHistoryEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetRequestId(v string) *DescribeHistoryEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetTotalCount(v int32) *DescribeHistoryEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) Validate() error {
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

type DescribeHistoryEventsResponseBodyItems struct {
	// The data overview.
	//
	// example:
	//
	// {}
	Data *DescribeHistoryEventsResponseBodyItemsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the task.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region ID of the region where the task is located.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The source of the event.
	//
	// example:
	//
	// loanBill
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The database version.
	//
	// example:
	//
	// 8.0
	Specversion *string `json:"Specversion,omitempty" xml:"Specversion,omitempty"`
	// The name of the event to be processed.
	//
	// example:
	//
	// QiTian
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The running time of the queried task. Unit: seconds.
	//
	// example:
	//
	// 1758680209206
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The event type.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHistoryEventsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponseBodyItems) GetData() *DescribeHistoryEventsResponseBodyItemsData {
	return s.Data
}

func (s *DescribeHistoryEventsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeHistoryEventsResponseBodyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeHistoryEventsResponseBodyItems) GetSource() *string {
	return s.Source
}

func (s *DescribeHistoryEventsResponseBodyItems) GetSpecversion() *string {
	return s.Specversion
}

func (s *DescribeHistoryEventsResponseBodyItems) GetSubject() *string {
	return s.Subject
}

func (s *DescribeHistoryEventsResponseBodyItems) GetTime() *string {
	return s.Time
}

func (s *DescribeHistoryEventsResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *DescribeHistoryEventsResponseBodyItems) SetData(v *DescribeHistoryEventsResponseBodyItemsData) *DescribeHistoryEventsResponseBodyItems {
	s.Data = v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetId(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetRegion(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetSource(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Source = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetSpecversion(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Specversion = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetSubject(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Subject = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetTime(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Time = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) SetType(v string) *DescribeHistoryEventsResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItems) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHistoryEventsResponseBodyItemsData struct {
	// The cloud service type of the application group. Valid values:
	//
	// - **web**: Web applications.
	//
	// - **native**: Native applications.
	//
	// example:
	//
	// web
	CmsProduct *string `json:"CmsProduct,omitempty" xml:"CmsProduct,omitempty"`
	// The type of the database.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The pagination parameter.
	//
	// example:
	//
	// 1
	DetailImpact *string `json:"DetailImpact,omitempty" xml:"DetailImpact,omitempty"`
	// The details of the instance operation.
	//
	// example:
	//
	// The instance needs to be restarted due to system maintenance
	DetailReason *string `json:"DetailReason,omitempty" xml:"DetailReason,omitempty"`
	// The end time of the query. The end time must be later than the start time, and the interval between the start time and the end time must be within 24 hours. Specify the time in the `YYYY-MM-DDThh:mmZ` format (UTC time).
	//
	// example:
	//
	// 2025-12-24T02:24:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The category of the system event. Valid values:
	//
	// - **Exception**: Exception events.
	//
	// - **Optimize**: Optimization events.
	//
	// - **Notification**: Notification events.
	//
	// - **Maintenance**: Scheduled O&M events.
	//
	// example:
	//
	// Exception
	EventCategory *string `json:"EventCategory,omitempty" xml:"EventCategory,omitempty"`
	// The event code.
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The event details.
	//
	// example:
	//
	// The instance will be restarted within the maintenance window
	EventDetail *string `json:"EventDetail,omitempty" xml:"EventDetail,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 600324
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The impact overview of the event.
	//
	// example:
	//
	// The instance connection may be briefly interrupted
	EventImpact *string `json:"EventImpact,omitempty" xml:"EventImpact,omitempty"`
	// The event level. Valid values:
	//
	// - **INFO**: Notification.
	//
	// - **WARN**: Warning.
	//
	// - **CRITICAL**: Critical.
	//
	// example:
	//
	// INFO
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The source of the event operation.
	//
	// example:
	//
	// System maintenance
	EventReason *string `json:"EventReason,omitempty" xml:"EventReason,omitempty"`
	// The status of the event. Valid values:
	//
	// - **Inquiring**: Inquiring.
	//
	// - **Scheduled**: Scheduled.
	//
	// - **Running**: Running.
	//
	// - **Succeed**: Succeeded.
	//
	// - **Failed**: Failed.
	//
	// - **Canceled**: Canceled.
	//
	// example:
	//
	// Inquiring
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The system event type. Valid values:
	//
	// - **SystemMaintenance.Reboot**: The instance is restarted due to system maintenance.
	//
	// - **SystemMaintenance.Redeploy**: Redeployment of the instance due to system maintenance.
	//
	// - **SystemFailure.Reboot**: The instance is restarted due to a system fault.
	//
	// - **SystemFailure.Redeploy**: Redeployment of the instance due to a system fault.
	//
	// - **SystemFailure.Delete**: The instance is released because instance creation failed.
	//
	// - **InstanceFailure.Reboot**: The instance is restarted due to an instance fault.
	//
	// - **InstanceExpiration.Stop**: The subscription instance is stopped due to expiration.
	//
	// - **InstanceExpiration.Delete**: The subscription instance is released due to expiration.
	//
	// - **AccountUnbalanced.Stop**: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// - **AccountUnbalanced.Delete**: The pay-as-you-go instance is released due to an overdue payment.
	//
	// example:
	//
	// SystemFailure.Delete
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-15T14:53:06+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-01-07T15:10:32+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The handling status.
	//
	// example:
	//
	// done
	HandleStatus *string `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	// Indicates whether the event has a lifecycle.
	//
	// example:
	//
	// 0
	HasLifeCycle *int32 `json:"HasLifeCycle,omitempty" xml:"HasLifeCycle,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pc-2ze150h1p29t***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// dhimgsearch
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the item is closed. Valid values:
	//
	// - **0**: Closed.
	//
	// - **1**: Open.
	//
	// example:
	//
	// 1
	IsClosed *int32 `json:"IsClosed,omitempty" xml:"IsClosed,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// polardb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - **Instance**: Instance resources.
	//
	// - **Host**: Host resources.
	//
	// - **User**: User resources.
	//
	// example:
	//
	// Host
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the source data.
	//
	// example:
	//
	// loanBill
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the query. Format: `yyyy-MM-ddTHH:mm:ssZ` (UTC time).
	//
	// example:
	//
	// 2025-02-27T02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the user to which the resource belongs.
	//
	// example:
	//
	// 1540497309282125
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHistoryEventsResponseBodyItemsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponseBodyItemsData) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetCmsProduct() *string {
	return s.CmsProduct
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetDbType() *string {
	return s.DbType
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetDetailImpact() *string {
	return s.DetailImpact
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetDetailReason() *string {
	return s.DetailReason
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventCategory() *string {
	return s.EventCategory
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventDetail() *string {
	return s.EventDetail
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventId() *string {
	return s.EventId
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventImpact() *string {
	return s.EventImpact
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventReason() *string {
	return s.EventReason
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetEventType() *string {
	return s.EventType
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetHandleStatus() *string {
	return s.HandleStatus
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetHasLifeCycle() *int32 {
	return s.HasLifeCycle
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetIsClosed() *int32 {
	return s.IsClosed
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetProduct() *string {
	return s.Product
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHistoryEventsResponseBodyItemsData) GetUid() *string {
	return s.Uid
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetCmsProduct(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.CmsProduct = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetDbType(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.DbType = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetDetailImpact(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.DetailImpact = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetDetailReason(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.DetailReason = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEndTime(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EndTime = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventCategory(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventCategory = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventCode(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventCode = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventDetail(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventDetail = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventId(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventImpact(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventImpact = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventLevel(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventLevel = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventReason(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventReason = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventStatus(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventStatus = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetEventType(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.EventType = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetGmtCreated(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetGmtModified(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.GmtModified = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetHandleStatus(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.HandleStatus = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetHasLifeCycle(v int32) *DescribeHistoryEventsResponseBodyItemsData {
	s.HasLifeCycle = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetInstanceId(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetInstanceName(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.InstanceName = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetIsClosed(v int32) *DescribeHistoryEventsResponseBodyItemsData {
	s.IsClosed = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetProduct(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.Product = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetRegionId(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetResourceType(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.ResourceType = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetSourceType(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.SourceType = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetStartTime(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.StartTime = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) SetUid(v string) *DescribeHistoryEventsResponseBodyItemsData {
	s.Uid = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyItemsData) Validate() error {
	return dara.Validate(s)
}
