// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotListConversationsResponseBody
	GetCode() *int64
	SetData(v *KopilotListConversationsResponseBodyData) *KopilotListConversationsResponseBody
	GetData() *KopilotListConversationsResponseBodyData
	SetRequestId(v string) *KopilotListConversationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotListConversationsResponseBody
	GetSuccess() *bool
}

type KopilotListConversationsResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned when the call is successful.
	Data *KopilotListConversationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2DF166F2-F581-5254-AAB6-B482083FA7B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotListConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotListConversationsResponseBody) GetData() *KopilotListConversationsResponseBodyData {
	return s.Data
}

func (s *KopilotListConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotListConversationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotListConversationsResponseBody) SetCode(v int64) *KopilotListConversationsResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotListConversationsResponseBody) SetData(v *KopilotListConversationsResponseBodyData) *KopilotListConversationsResponseBody {
	s.Data = v
	return s
}

func (s *KopilotListConversationsResponseBody) SetRequestId(v string) *KopilotListConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotListConversationsResponseBody) SetSuccess(v bool) *KopilotListConversationsResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotListConversationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyData struct {
	// The overview of tasks, notification channels, and quotas for the current account. Returned only when the request includes the overview and the feature is available.
	AutomationOverview *KopilotListConversationsResponseBodyDataAutomationOverview `json:"AutomationOverview,omitempty" xml:"AutomationOverview,omitempty" type:"Struct"`
	// The list of session IDs.
	ConversationIds []*string `json:"ConversationIds,omitempty" xml:"ConversationIds,omitempty" type:"Repeated"`
	// The number of entries returned on the current page in a paged query.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// A mapping that uses session IDs as keys and associated task summaries as values.
	ScheduledTaskInfoBySessionId map[string]*DataScheduledTaskInfoBySessionIdValue `json:"ScheduledTaskInfoBySessionId,omitempty" xml:"ScheduledTaskInfoBySessionId,omitempty"`
	// The quota for the number of scheduled tasks of the current primary account in this environment, counted across regions.
	ScheduledTaskQuota *KopilotListConversationsResponseBodyDataScheduledTaskQuota `json:"ScheduledTaskQuota,omitempty" xml:"ScheduledTaskQuota,omitempty" type:"Struct"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The user UID.
	//
	// example:
	//
	// 1734455674565
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s KopilotListConversationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyData) GetAutomationOverview() *KopilotListConversationsResponseBodyDataAutomationOverview {
	return s.AutomationOverview
}

func (s *KopilotListConversationsResponseBodyData) GetConversationIds() []*string {
	return s.ConversationIds
}

func (s *KopilotListConversationsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *KopilotListConversationsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *KopilotListConversationsResponseBodyData) GetScheduledTaskInfoBySessionId() map[string]*DataScheduledTaskInfoBySessionIdValue {
	return s.ScheduledTaskInfoBySessionId
}

func (s *KopilotListConversationsResponseBodyData) GetScheduledTaskQuota() *KopilotListConversationsResponseBodyDataScheduledTaskQuota {
	return s.ScheduledTaskQuota
}

func (s *KopilotListConversationsResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *KopilotListConversationsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *KopilotListConversationsResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *KopilotListConversationsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *KopilotListConversationsResponseBodyData) SetAutomationOverview(v *KopilotListConversationsResponseBodyDataAutomationOverview) *KopilotListConversationsResponseBodyData {
	s.AutomationOverview = v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetConversationIds(v []*string) *KopilotListConversationsResponseBodyData {
	s.ConversationIds = v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetCount(v int32) *KopilotListConversationsResponseBodyData {
	s.Count = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetPage(v int32) *KopilotListConversationsResponseBodyData {
	s.Page = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetScheduledTaskInfoBySessionId(v map[string]*DataScheduledTaskInfoBySessionIdValue) *KopilotListConversationsResponseBodyData {
	s.ScheduledTaskInfoBySessionId = v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetScheduledTaskQuota(v *KopilotListConversationsResponseBodyDataScheduledTaskQuota) *KopilotListConversationsResponseBodyData {
	s.ScheduledTaskQuota = v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetSize(v int32) *KopilotListConversationsResponseBodyData {
	s.Size = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetTotal(v int64) *KopilotListConversationsResponseBodyData {
	s.Total = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetTotalPages(v int32) *KopilotListConversationsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetUserId(v string) *KopilotListConversationsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) Validate() error {
	if s.AutomationOverview != nil {
		if err := s.AutomationOverview.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduledTaskQuota != nil {
		if err := s.ScheduledTaskQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyDataAutomationOverview struct {
	// The time when the overview was generated, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	AsOf *string `json:"AsOf,omitempty" xml:"AsOf,omitempty"`
	// The quota for the number of notification channels of the current primary account in this environment, counted across regions.
	DestinationQuota *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota `json:"DestinationQuota,omitempty" xml:"DestinationQuota,omitempty" type:"Struct"`
	// The paginated results of notification channels. Webhook URLs and signing keys are not returned.
	Destinations *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Struct"`
	// The home region for tasks and notification channels.
	//
	// example:
	//
	// cn-beijing
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The quota for the number of scheduled tasks of the current primary account in this environment, counted across regions.
	Quota *KopilotListConversationsResponseBodyDataAutomationOverviewQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The paginated results of scheduled tasks for the account.
	Tasks *KopilotListConversationsResponseBodyDataAutomationOverviewTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverview) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverview) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetAsOf() *string {
	return s.AsOf
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetDestinationQuota() *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota {
	return s.DestinationQuota
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetDestinations() *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations {
	return s.Destinations
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetQuota() *KopilotListConversationsResponseBodyDataAutomationOverviewQuota {
	return s.Quota
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) GetTasks() *KopilotListConversationsResponseBodyDataAutomationOverviewTasks {
	return s.Tasks
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetAsOf(v string) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.AsOf = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetDestinationQuota(v *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.DestinationQuota = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetDestinations(v *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.Destinations = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetHomeRegion(v string) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.HomeRegion = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetQuota(v *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.Quota = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) SetTasks(v *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) *KopilotListConversationsResponseBodyDataAutomationOverview {
	s.Tasks = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverview) Validate() error {
	if s.DestinationQuota != nil {
		if err := s.DestinationQuota.Validate(); err != nil {
			return err
		}
	}
	if s.Destinations != nil {
		if err := s.Destinations.Validate(); err != nil {
			return err
		}
	}
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota struct {
	// The maximum number of tasks or channels allowed, subject to the actual configuration.
	//
	// example:
	//
	// 3
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The remaining quota, calculated as the limit minus the used quantity. The minimum value is 0.
	//
	// example:
	//
	// 2
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// The number of non-deleted notification channels. Disabled channels still consume quota.
	//
	// example:
	//
	// 1
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) GetLimit() *int32 {
	return s.Limit
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) GetRemaining() *int64 {
	return s.Remaining
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) GetUsed() *int64 {
	return s.Used
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) SetLimit(v int32) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota {
	s.Limit = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) SetRemaining(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota {
	s.Remaining = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) SetUsed(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota {
	s.Used = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationQuota) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataAutomationOverviewDestinations struct {
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The list of tasks or notification channels on the current page.
	Items []*KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The cursor for the next page. This field is empty if no more pages are available.
	//
	// example:
	//
	// 123
	NextCursor *string `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	// The total number of non-deleted records in the list.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) GetHasMore() *bool {
	return s.HasMore
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) GetItems() []*KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	return s.Items
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) GetNextCursor() *string {
	return s.NextCursor
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) GetTotal() *int64 {
	return s.Total
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) SetHasMore(v bool) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations {
	s.HasMore = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) SetItems(v []*KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations {
	s.Items = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) SetNextCursor(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations {
	s.NextCursor = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) SetTotal(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations {
	s.Total = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinations) Validate() error {
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

type KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems struct {
	// The time when the record was created, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The unique identifier of the notification channel.
	//
	// example:
	//
	// dst_0123456789abcdef0123456789abcdef
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The most recent delivery record. This field is empty if no delivery has been made.
	LastDelivery *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery `json:"LastDelivery,omitempty" xml:"LastDelivery,omitempty" type:"Struct"`
	// The name of the notification channel.
	//
	// example:
	//
	// Inspection Notification Group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration status of the channel. Valid values:
	//
	// - ACTIVE: enabled.
	//
	// - DISABLED: disabled.
	//
	// Being enabled does not indicate that a delivery has been successfully sent.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the notification channel. DINGTALK_WEBHOOK indicates a DingTalk custom chatbot.
	//
	// example:
	//
	// DINGTALK_WEBHOOK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the record was last updated, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The version number used for concurrent update verification of the record.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetDestinationId() *string {
	return s.DestinationId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetLastDelivery() *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	return s.LastDelivery
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetName() *string {
	return s.Name
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetType() *string {
	return s.Type
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) GetVersion() *int64 {
	return s.Version
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetCreatedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.CreatedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetDestinationId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.DestinationId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetLastDelivery(v *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.LastDelivery = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetName(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.Name = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetStatus(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.Status = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetType(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.Type = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetUpdatedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.UpdatedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) SetVersion(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems {
	s.Version = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItems) Validate() error {
	if s.LastDelivery != nil {
		if err := s.LastDelivery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery struct {
	// The time when the notification service accepted the delivery, in UTC ISO 8601 format. This field is empty if the delivery has not been accepted.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	AcceptedAt *string `json:"AcceptedAt,omitempty" xml:"AcceptedAt,omitempty"`
	// The time when the record was created, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The unique identifier of the delivery record.
	//
	// example:
	//
	// dlv_0123456789abcdef0123456789abcdef
	DeliveryId *string `json:"DeliveryId,omitempty" xml:"DeliveryId,omitempty"`
	// The time of the most recent delivery attempt, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	LastAttemptAt *string `json:"LastAttemptAt,omitempty" xml:"LastAttemptAt,omitempty"`
	// The delivery type. Valid values:
	//
	// - AUTO: automatic delivery.
	//
	// - MANUAL: manual delivery.
	//
	// - TEST: connectivity test.
	//
	// example:
	//
	// AUTO
	SendKind *string `json:"SendKind,omitempty" xml:"SendKind,omitempty"`
	// The most recent delivery status. ACCEPTED indicates that the notification service has accepted the delivery, but does not mean the recipient has read it.
	//
	// example:
	//
	// ACCEPTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetAcceptedAt() *string {
	return s.AcceptedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetDeliveryId() *string {
	return s.DeliveryId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetLastAttemptAt() *string {
	return s.LastAttemptAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetSendKind() *string {
	return s.SendKind
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetAcceptedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.AcceptedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetCreatedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.CreatedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetDeliveryId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.DeliveryId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetLastAttemptAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.LastAttemptAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetSendKind(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.SendKind = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) SetStatus(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery {
	s.Status = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewDestinationsItemsLastDelivery) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataAutomationOverviewQuota struct {
	// The maximum number of tasks or channels allowed, subject to the actual configuration.
	//
	// example:
	//
	// 3
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The remaining quota, calculated as the limit minus the used quantity. The minimum value is 0.
	//
	// example:
	//
	// 2
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// The number of task quota slots consumed. Tasks in DRAFT, ENABLED, PAUSED, or NEEDS_AUTH status are counted. Completed and deleted tasks do not consume quota.
	//
	// example:
	//
	// 1
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewQuota) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewQuota) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) GetLimit() *int32 {
	return s.Limit
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) GetRemaining() *int64 {
	return s.Remaining
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) GetUsed() *int64 {
	return s.Used
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) SetLimit(v int32) *KopilotListConversationsResponseBodyDataAutomationOverviewQuota {
	s.Limit = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) SetRemaining(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewQuota {
	s.Remaining = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) SetUsed(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewQuota {
	s.Used = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewQuota) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataAutomationOverviewTasks struct {
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The list of tasks or notification channels on the current page.
	Items []*KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The cursor for the next page. This field is empty if no more pages are available.
	//
	// example:
	//
	// 123
	NextCursor *string `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	// The total number of non-deleted records in the list.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasks) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasks) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) GetHasMore() *bool {
	return s.HasMore
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) GetItems() []*KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	return s.Items
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) GetNextCursor() *string {
	return s.NextCursor
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) GetTotal() *int64 {
	return s.Total
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) SetHasMore(v bool) *KopilotListConversationsResponseBodyDataAutomationOverviewTasks {
	s.HasMore = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) SetItems(v []*KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) *KopilotListConversationsResponseBodyDataAutomationOverviewTasks {
	s.Items = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) SetNextCursor(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasks {
	s.NextCursor = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) SetTotal(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewTasks {
	s.Total = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasks) Validate() error {
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

type KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems struct {
	// The currently queued or running execution record. This field is empty if no active run exists.
	ActiveRun *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun `json:"ActiveRun,omitempty" xml:"ActiveRun,omitempty" type:"Struct"`
	// The task execution configuration.
	Configuration *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The time when the record was created, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The most recent completed run record, including failed runs. This field is empty if no record exists.
	LastCompletedRun *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun `json:"LastCompletedRun,omitempty" xml:"LastCompletedRun,omitempty" type:"Struct"`
	// The name of the scheduled task.
	//
	// example:
	//
	// Kafka Resource Inspection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next scheduled execution time, in UTC ISO 8601 format. This field is empty if no next run is scheduled.
	//
	// example:
	//
	// 2026-09-17T12:15:00Z
	NextRunAt *string `json:"NextRunAt,omitempty" xml:"NextRunAt,omitempty"`
	// The human-readable description of the execution schedule.
	//
	// example:
	//
	// Every 900 seconds
	ScheduleDescription *string `json:"ScheduleDescription,omitempty" xml:"ScheduleDescription,omitempty"`
	// The session mode for displaying results. Valid values:
	//
	// - SHARED: shared session.
	//
	// - PER_RUN: independent session for each run.
	//
	// example:
	//
	// SHARED
	SessionMode *string `json:"SessionMode,omitempty" xml:"SessionMode,omitempty"`
	// The ID of the source session from which the task was created.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	SourceSessionId *string `json:"SourceSessionId,omitempty" xml:"SourceSessionId,omitempty"`
	// The status of the scheduled task. Valid values:
	//
	// - DRAFT: Draft.
	//
	// - ENABLED: Enabled.
	//
	// - PAUSED: Paused.
	//
	// - NEEDS_AUTH: Pending authorization.
	//
	// - COMPLETED: Completed.
	//
	// This status is independent of the running status.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the target session that stores run results in shared mode.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	TargetSessionId *string `json:"TargetSessionId,omitempty" xml:"TargetSessionId,omitempty"`
	// The unique identifier of the scheduled task.
	//
	// example:
	//
	// task_0123456789abcdef0123456789abcdef
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The time when the record was last updated, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The version number used for concurrent update verification of the current record.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetActiveRun() *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	return s.ActiveRun
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetConfiguration() *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	return s.Configuration
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetLastCompletedRun() *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	return s.LastCompletedRun
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetName() *string {
	return s.Name
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetScheduleDescription() *string {
	return s.ScheduleDescription
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetSessionMode() *string {
	return s.SessionMode
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetSourceSessionId() *string {
	return s.SourceSessionId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetTargetSessionId() *string {
	return s.TargetSessionId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetTaskId() *string {
	return s.TaskId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) GetVersion() *int64 {
	return s.Version
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetActiveRun(v *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.ActiveRun = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetConfiguration(v *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.Configuration = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetCreatedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.CreatedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetLastCompletedRun(v *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.LastCompletedRun = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetName(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.Name = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetNextRunAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.NextRunAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetScheduleDescription(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.ScheduleDescription = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetSessionMode(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.SessionMode = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetSourceSessionId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.SourceSessionId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetStatus(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.Status = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetTargetSessionId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.TargetSessionId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetTaskId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.TaskId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetUpdatedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.UpdatedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) SetVersion(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems {
	s.Version = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItems) Validate() error {
	if s.ActiveRun != nil {
		if err := s.ActiveRun.Validate(); err != nil {
			return err
		}
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.LastCompletedRun != nil {
		if err := s.LastCompletedRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun struct {
	// The error code of a failed run. This field is empty if no error occurred.
	//
	// example:
	//
	// UPSTREAM_TIMEOUT
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The user-facing failure reason. This field is empty if no error occurred.
	//
	// example:
	//
	// Model or tool calling invoke timed out. Try again later
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the run finished, in UTC ISO 8601 format. This field is typically empty for queued or running executions.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The unique identifier of a single run.
	//
	// example:
	//
	// run_0123456789abcdef0123456789abcdef
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The status of a single run. Valid values:
	//
	// - QUEUED: queued.
	//
	// - RUNNING: running.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GetRunId() *string {
	return s.RunId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) SetErrorCode(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	s.ErrorCode = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) SetErrorMessage(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	s.ErrorMessage = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) SetFinishedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	s.FinishedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) SetRunId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	s.RunId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) SetStatus(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun {
	s.Status = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsActiveRun) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration struct {
	// The six-field cron expression with the seconds field fixed to 0. This field is used only for the CRON schedule type.
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The list of notification channel IDs that receive task results. A maximum of 3 IDs are supported.
	DestinationIds []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	// The list of instance IDs within the query scope.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The instruction for the scheduled task execution.
	//
	// example:
	//
	// Query the current risks of the specified instances and summarize the results
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The fixed interval in seconds. Valid values: 900 to 31536000. This field is used only for the FIXED_INTERVAL schedule type.
	//
	// example:
	//
	// 900
	IntervalSeconds *int32 `json:"IntervalSeconds,omitempty" xml:"IntervalSeconds,omitempty"`
	// The time budget for a single task run, in seconds. Default value: 600. Valid values: 30 to 1800.
	//
	// example:
	//
	// 600
	MaxRunSeconds *int32 `json:"MaxRunSeconds,omitempty" xml:"MaxRunSeconds,omitempty"`
	// The token budget for a single task run. Default value: 3000000. Valid values: 1000 to 10000000.
	//
	// example:
	//
	// 3000000
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// The name of the scheduled task.
	//
	// example:
	//
	// Kafka Resource Inspection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region where the queried resources reside.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource scope mode. Valid values:
	//
	// - ACCOUNT: account-level query.
	//
	// - NONE: no resource task.
	//
	// - EXPLICIT: specified instances.
	//
	// example:
	//
	// ACCOUNT
	ResourceMode *string `json:"ResourceMode,omitempty" xml:"ResourceMode,omitempty"`
	// The one-time execution time in ISO 8601 format with time zone. This field is used only for the ONCE schedule type.
	//
	// example:
	//
	// 2026-09-18T12:00:00Z
	RunAt *string `json:"RunAt,omitempty" xml:"RunAt,omitempty"`
	// The schedule type. Valid values:
	//
	// - ONCE: one-time execution.
	//
	// - CRON: cron expression.
	//
	// - FIXED_INTERVAL: fixed interval.
	//
	// example:
	//
	// FIXED_INTERVAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The session mode for displaying results. Valid values:
	//
	// - SHARED: shared session.
	//
	// - PER_RUN: independent session for each run.
	//
	// example:
	//
	// SHARED
	SessionMode *string `json:"SessionMode,omitempty" xml:"SessionMode,omitempty"`
	// The ID of the target session that stores run results in shared mode.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	TargetSessionId *string `json:"TargetSessionId,omitempty" xml:"TargetSessionId,omitempty"`
	// The scheduling time zone. Default value: Asia/Shanghai.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetCronExpression() *string {
	return s.CronExpression
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetDestinationIds() []*string {
	return s.DestinationIds
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetInstruction() *string {
	return s.Instruction
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetIntervalSeconds() *int32 {
	return s.IntervalSeconds
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetMaxRunSeconds() *int32 {
	return s.MaxRunSeconds
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetName() *string {
	return s.Name
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetResourceMode() *string {
	return s.ResourceMode
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetRunAt() *string {
	return s.RunAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetSessionMode() *string {
	return s.SessionMode
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetTargetSessionId() *string {
	return s.TargetSessionId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) GetTimezone() *string {
	return s.Timezone
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetCronExpression(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.CronExpression = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetDestinationIds(v []*string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.DestinationIds = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetInstanceIds(v []*string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.InstanceIds = v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetInstruction(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.Instruction = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetIntervalSeconds(v int32) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.IntervalSeconds = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetMaxRunSeconds(v int32) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.MaxRunSeconds = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetMaxTokens(v int64) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.MaxTokens = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetName(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.Name = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetRegionId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.RegionId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetResourceMode(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.ResourceMode = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetRunAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.RunAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetScheduleType(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetSessionMode(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.SessionMode = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetTargetSessionId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.TargetSessionId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) SetTimezone(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration {
	s.Timezone = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsConfiguration) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun struct {
	// The error code of a failed run. This field is empty if no error occurred.
	//
	// example:
	//
	// UPSTREAM_TIMEOUT
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The user-facing failure reason. This field is empty if no error occurred.
	//
	// example:
	//
	// Model or tool calling invoke timed out. Try again later
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the run finished, in UTC ISO 8601 format. This field is empty if the run has not finished.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The unique identifier of a single run.
	//
	// example:
	//
	// run_0123456789abcdef0123456789abcdef
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The status of the most recent completed run. For example, SUCCEEDED indicates success and FAILED indicates failure.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GetRunId() *string {
	return s.RunId
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) SetErrorCode(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	s.ErrorCode = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) SetErrorMessage(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	s.ErrorMessage = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) SetFinishedAt(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	s.FinishedAt = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) SetRunId(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	s.RunId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) SetStatus(v string) *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun {
	s.Status = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataAutomationOverviewTasksItemsLastCompletedRun) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationsResponseBodyDataScheduledTaskQuota struct {
	// The maximum number of tasks or channels allowed, subject to the actual configuration.
	//
	// example:
	//
	// 3
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The remaining quota, calculated as the limit minus the used quantity. The minimum value is 0.
	//
	// example:
	//
	// 2
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// The number of task quota slots consumed. Tasks in DRAFT, ENABLED, PAUSED, or NEEDS_AUTH status are counted. Completed and deleted tasks do not consume quota.
	//
	// example:
	//
	// 1
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s KopilotListConversationsResponseBodyDataScheduledTaskQuota) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyDataScheduledTaskQuota) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) GetLimit() *int32 {
	return s.Limit
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) GetRemaining() *int64 {
	return s.Remaining
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) GetUsed() *int64 {
	return s.Used
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) SetLimit(v int32) *KopilotListConversationsResponseBodyDataScheduledTaskQuota {
	s.Limit = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) SetRemaining(v int64) *KopilotListConversationsResponseBodyDataScheduledTaskQuota {
	s.Remaining = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) SetUsed(v int64) *KopilotListConversationsResponseBodyDataScheduledTaskQuota {
	s.Used = &v
	return s
}

func (s *KopilotListConversationsResponseBodyDataScheduledTaskQuota) Validate() error {
	return dara.Validate(s)
}
