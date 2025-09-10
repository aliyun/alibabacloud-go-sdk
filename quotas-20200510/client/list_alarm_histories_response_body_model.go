// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmHistories(v []*ListAlarmHistoriesResponseBodyAlarmHistories) *ListAlarmHistoriesResponseBody
	GetAlarmHistories() []*ListAlarmHistoriesResponseBodyAlarmHistories
	SetMaxResults(v int32) *ListAlarmHistoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlarmHistoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAlarmHistoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlarmHistoriesResponseBody
	GetTotalCount() *int32
}

type ListAlarmHistoriesResponseBody struct {
	// The details of the quota alert records.
	AlarmHistories []*ListAlarmHistoriesResponseBodyAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB38DDF9-B1E0-48C1-9966-19C443C2841E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBody) GetAlarmHistories() []*ListAlarmHistoriesResponseBodyAlarmHistories {
	return s.AlarmHistories
}

func (s *ListAlarmHistoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlarmHistoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlarmHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlarmHistoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlarmHistoriesResponseBody) SetAlarmHistories(v []*ListAlarmHistoriesResponseBodyAlarmHistories) *ListAlarmHistoriesResponseBody {
	s.AlarmHistories = v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetMaxResults(v int32) *ListAlarmHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetNextToken(v string) *ListAlarmHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetRequestId(v string) *ListAlarmHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetTotalCount(v int32) *ListAlarmHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlarmHistoriesResponseBodyAlarmHistories struct {
	// The name of the quota alert.
	//
	// example:
	//
	// security_groups
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the quota alert rule was created.
	//
	// example:
	//
	// 2021-01-24T09:20:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The notification methods of the quota alert.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The quota alert contact.
	//
	// example:
	//
	// accountContact
	NotifyTarget *string `json:"NotifyTarget,omitempty" xml:"NotifyTarget,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 31
	QuotaUsage *float32 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// The threshold to trigger quota alerts.
	//
	// example:
	//
	// 29
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The percentage of the quota alert threshold.
	//
	// example:
	//
	// 80
	ThresholdPercent *float32 `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmHistoriesResponseBodyAlarmHistories) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetAlarmName() *string {
	return s.AlarmName
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetNotifyTarget() *string {
	return s.NotifyTarget
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetQuotaUsage() *float32 {
	return s.QuotaUsage
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) GetThresholdPercent() *float32 {
	return s.ThresholdPercent
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetAlarmName(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetCreateTime(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.CreateTime = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyChannels(v []*string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyChannels = v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetNotifyTarget(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.NotifyTarget = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetProductCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ProductCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaActionCode(v string) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaActionCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetQuotaUsage(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.QuotaUsage = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThreshold(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.Threshold = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) SetThresholdPercent(v float32) *ListAlarmHistoriesResponseBodyAlarmHistories {
	s.ThresholdPercent = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyAlarmHistories) Validate() error {
	return dara.Validate(s)
}
