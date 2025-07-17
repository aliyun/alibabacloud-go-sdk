// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDIJobEventsResponseBodyPagingInfo) *ListDIJobEventsResponseBody
	GetPagingInfo() *ListDIJobEventsResponseBodyPagingInfo
	SetRequestId(v string) *ListDIJobEventsResponseBody
	GetRequestId() *string
}

type ListDIJobEventsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIJobEventsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 645F6D68-9C29-5961-80B1-BDD4B794C22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBody) GetPagingInfo() *ListDIJobEventsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDIJobEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIJobEventsResponseBody) SetPagingInfo(v *ListDIJobEventsResponseBodyPagingInfo) *ListDIJobEventsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobEventsResponseBody) SetRequestId(v string) *ListDIJobEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIJobEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDIJobEventsResponseBodyPagingInfo struct {
	// The events returned. The value of this parameter is an array.
	DIJobEvent []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent `json:"DIJobEvent,omitempty" xml:"DIJobEvent,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2524
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfo) GetDIJobEvent() []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	return s.DIJobEvent
}

func (s *ListDIJobEventsResponseBodyPagingInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIJobEventsResponseBodyPagingInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIJobEventsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetDIJobEvent(v []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent) *ListDIJobEventsResponseBodyPagingInfo {
	s.DIJobEvent = v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDIJobEventsResponseBodyPagingInfoDIJobEvent struct {
	// The processing result of the DDL event. Valid values: Critical, Ignore, Normal, and Warning.
	//
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The alert notification method. Valid values: Phone, Mail, Sms, Ding, and Webhook.
	//
	// example:
	//
	// Phone
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The time when the event was created.
	//
	// example:
	//
	// 1663573162
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The alert details.
	//
	// example:
	//
	// aggregator:avg [**] for 5 minutes, service maybe abnormal
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The DDL statement of the destination table.
	//
	// example:
	//
	// alter table table2 ***
	DstSql *string `json:"DstSql,omitempty" xml:"DstSql,omitempty"`
	// The name of the destination table.
	//
	// example:
	//
	// table2
	DstTable *string `json:"DstTable,omitempty" xml:"DstTable,omitempty"`
	// The error logs for failovers.
	//
	// example:
	//
	// 2024-05-29 15:11:31,377 [main] INFO com.*.**.di.core.metrics.:21 []  {****}
	//
	// 2024-05-29 15:11:31,384 [main] INFO *.aliyun.*.di.*.*.metrics.*:27 [] - Open MarioDiReporter
	//
	// 2024-05-29 15:11:33,248 [flink-akka.*.*-dispatcher-17] INFO
	FailoverMessage *string `json:"FailoverMessage,omitempty" xml:"FailoverMessage,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The severity level of the alert. Valid values: Warning and Critical.
	//
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The DDL statement of the source table.
	//
	// example:
	//
	// alter table table1 ***
	SrcSql *string `json:"SrcSql,omitempty" xml:"SrcSql,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// table1
	SrcTable *string `json:"SrcTable,omitempty" xml:"SrcTable,omitempty"`
	// The sending status of an alert notification. Valid values: Success, Fail, and Silence.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the alert event.
	//
	// 	- Heartbeat
	//
	// 	- Delay
	//
	// 	- FailoverCount
	//
	// 	- DdlReport
	//
	// 	- ResourceUtilization
	//
	// example:
	//
	// Delay
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetAction() *string {
	return s.Action
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetChannels() *string {
	return s.Channels
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetDetail() *string {
	return s.Detail
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetDstSql() *string {
	return s.DstSql
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetDstTable() *string {
	return s.DstTable
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetFailoverMessage() *string {
	return s.FailoverMessage
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetId() *string {
	return s.Id
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetSeverity() *string {
	return s.Severity
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetSrcSql() *string {
	return s.SrcSql
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetSrcTable() *string {
	return s.SrcTable
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetStatus() *string {
	return s.Status
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GetType() *string {
	return s.Type
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetAction(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Action = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetChannels(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Channels = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetCreateTime(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.CreateTime = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDetail(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Detail = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetFailoverMessage(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.FailoverMessage = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetId(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Id = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSeverity(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Severity = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetStatus(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Status = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetType(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Type = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) Validate() error {
	return dara.Validate(s)
}
