// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListDataQualityAlertRulesResponseBodyPageInfo) *ListDataQualityAlertRulesResponseBody
	GetPageInfo() *ListDataQualityAlertRulesResponseBodyPageInfo
	SetRequestId(v string) *ListDataQualityAlertRulesResponseBody
	GetRequestId() *string
}

type ListDataQualityAlertRulesResponseBody struct {
	// Alert rule configurations.
	PageInfo *ListDataQualityAlertRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBody) GetPageInfo() *ListDataQualityAlertRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListDataQualityAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityAlertRulesResponseBody) SetPageInfo(v *ListDataQualityAlertRulesResponseBodyPageInfo) *ListDataQualityAlertRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBody) SetRequestId(v string) *ListDataQualityAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityAlertRulesResponseBodyPageInfo struct {
	// The list of alert rule configurations.
	DataQualityAlertRules []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules `json:"DataQualityAlertRules,omitempty" xml:"DataQualityAlertRules,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 335
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityAlertRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) GetDataQualityAlertRules() []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	return s.DataQualityAlertRules
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) SetDataQualityAlertRules(v []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) *ListDataQualityAlertRulesResponseBodyPageInfo {
	s.DataQualityAlertRules = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) SetPageNumber(v int32) *ListDataQualityAlertRulesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) SetPageSize(v int32) *ListDataQualityAlertRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListDataQualityAlertRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules struct {
	// The alert conditions.
	//
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The ID of the data quality monitor alert rule.
	//
	// example:
	//
	// 26433
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Alert notification configurations.
	Notification *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The project ID.
	//
	// example:
	//
	// 59094
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Monitored targets of the data quality alert rule.
	Target *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GetNotification() *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification {
	return s.Notification
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) GetTarget() *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget {
	return s.Target
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) SetCondition(v string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	s.Condition = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) SetId(v int64) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	s.Id = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) SetNotification(v *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	s.Notification = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) SetProjectId(v int64) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) SetTarget(v *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules {
	s.Target = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRules) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification struct {
	// In Channels, you can set both Email and Sms at the same time. In other cases, only one channel can be set.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The alert recipients.
	Receivers []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) GetChannels() []*string {
	return s.Channels
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) GetReceivers() []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers {
	return s.Receivers
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) SetChannels(v []*string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification {
	s.Channels = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) SetReceivers(v []*ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification {
	s.Receivers = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotification) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers struct {
	// Additional configurations required for the alert recipients. When ReceiverType is DingdingUrl, you can set `{"atAll":true}` to mention all members.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of alert recipients.
	//
	// 	- ShiftSchedule
	//
	// 	- WebhookUrl
	//
	// 	- FeishuUrl
	//
	// 	- TaskOwner
	//
	// 	- WeixinUrl
	//
	// 	- DingdingUrl
	//
	// 	- DataQualityScanOwner
	//
	// 	- AliUid
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The value of alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) SetExtension(v string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) SetReceiverType(v string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) SetReceiverValues(v []*string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget struct {
	// The list of monitored target IDs
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the monitored target. Only DataQualityScan is supported.
	//
	// example:
	//
	// DataQualityScan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) GetType() *string {
	return s.Type
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) SetIds(v []*int64) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget {
	s.Ids = v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) SetType(v string) *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget {
	s.Type = &v
	return s
}

func (s *ListDataQualityAlertRulesResponseBodyPageInfoDataQualityAlertRulesTarget) Validate() error {
	return dara.Validate(s)
}
