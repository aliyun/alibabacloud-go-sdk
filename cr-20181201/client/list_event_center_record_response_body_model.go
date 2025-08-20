// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEventCenterRecordResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListEventCenterRecordResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListEventCenterRecordResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListEventCenterRecordResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListEventCenterRecordResponseBodyRecords) *ListEventCenterRecordResponseBody
	GetRecords() []*ListEventCenterRecordResponseBodyRecords
	SetRequestId(v string) *ListEventCenterRecordResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEventCenterRecordResponseBody
	GetTotalCount() *int32
}

type ListEventCenterRecordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of historical events.
	//
	// example:
	//
	// []
	Records []*ListEventCenterRecordResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 665C7A5E-BAEC-5BCD-AF9F-5F9260D672BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total entries of historical events.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEventCenterRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEventCenterRecordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListEventCenterRecordResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListEventCenterRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEventCenterRecordResponseBody) GetRecords() []*ListEventCenterRecordResponseBodyRecords {
	return s.Records
}

func (s *ListEventCenterRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventCenterRecordResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEventCenterRecordResponseBody) SetCode(v string) *ListEventCenterRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetIsSuccess(v bool) *ListEventCenterRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetPageNo(v int32) *ListEventCenterRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetPageSize(v int32) *ListEventCenterRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetRecords(v []*ListEventCenterRecordResponseBodyRecords) *ListEventCenterRecordResponseBody {
	s.Records = v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetRequestId(v string) *ListEventCenterRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetTotalCount(v int32) *ListEventCenterRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventCenterRecordResponseBodyRecords struct {
	// The time when the event was created.
	//
	// example:
	//
	// 1638188622000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event notification channel.
	//
	// example:
	//
	// EVENT_BRIDGE
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The ID of the event notification.
	//
	// example:
	//
	// 7d478419-61df-49e5-b92b-30ce730c2127
	EventNotifyId *string `json:"EventNotifyId,omitempty" xml:"EventNotifyId,omitempty"`
	// The notification method. Valid values:
	//
	// 	- `http`: The notification is sent over HTTP.
	//
	// 	- `https`: The notification is sent over HTTPS.
	//
	// 	- `dingding`: The notification is sent by using DingTalk.
	//
	// example:
	//
	// http
	EventNotifyMethod *string `json:"EventNotifyMethod,omitempty" xml:"EventNotifyMethod,omitempty"`
	// The type of the event.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-gl34plsa****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// mychain
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the event record.
	//
	// example:
	//
	// crecrr-ctdbzwtkpr*****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// ruby-2.4.0
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtla*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// chain-demo
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tags.
	//
	// example:
	//
	// ruby-2.4.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The time when the event was last updated.
	//
	// example:
	//
	// 1638188622000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEventCenterRecordResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponseBodyRecords) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListEventCenterRecordResponseBodyRecords) GetEventChannel() *string {
	return s.EventChannel
}

func (s *ListEventCenterRecordResponseBodyRecords) GetEventNotifyId() *string {
	return s.EventNotifyId
}

func (s *ListEventCenterRecordResponseBodyRecords) GetEventNotifyMethod() *string {
	return s.EventNotifyMethod
}

func (s *ListEventCenterRecordResponseBodyRecords) GetEventType() *string {
	return s.EventType
}

func (s *ListEventCenterRecordResponseBodyRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventCenterRecordResponseBodyRecords) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEventCenterRecordResponseBodyRecords) GetRecordId() *string {
	return s.RecordId
}

func (s *ListEventCenterRecordResponseBodyRecords) GetRepoName() *string {
	return s.RepoName
}

func (s *ListEventCenterRecordResponseBodyRecords) GetRuleId() *string {
	return s.RuleId
}

func (s *ListEventCenterRecordResponseBodyRecords) GetRuleName() *string {
	return s.RuleName
}

func (s *ListEventCenterRecordResponseBodyRecords) GetTag() *string {
	return s.Tag
}

func (s *ListEventCenterRecordResponseBodyRecords) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListEventCenterRecordResponseBodyRecords) SetCreateTime(v int64) *ListEventCenterRecordResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventChannel(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventChannel = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventNotifyId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventNotifyId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventNotifyMethod(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventNotifyMethod = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventType(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventType = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetInstanceId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.InstanceId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetNamespace(v string) *ListEventCenterRecordResponseBodyRecords {
	s.Namespace = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRecordId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRepoName(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RepoName = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRuleId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RuleId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRuleName(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RuleName = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetTag(v string) *ListEventCenterRecordResponseBodyRecords {
	s.Tag = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetUpdateTime(v int64) *ListEventCenterRecordResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
