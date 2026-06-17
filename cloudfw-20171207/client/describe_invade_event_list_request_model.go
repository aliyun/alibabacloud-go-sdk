// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsIP(v string) *DescribeInvadeEventListRequest
	GetAssetsIP() *string
	SetAssetsInstanceId(v string) *DescribeInvadeEventListRequest
	GetAssetsInstanceId() *string
	SetAssetsInstanceName(v string) *DescribeInvadeEventListRequest
	GetAssetsInstanceName() *string
	SetCurrentPage(v string) *DescribeInvadeEventListRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeInvadeEventListRequest
	GetEndTime() *string
	SetEventKey(v string) *DescribeInvadeEventListRequest
	GetEventKey() *string
	SetEventName(v string) *DescribeInvadeEventListRequest
	GetEventName() *string
	SetEventUuid(v string) *DescribeInvadeEventListRequest
	GetEventUuid() *string
	SetIsIgnore(v string) *DescribeInvadeEventListRequest
	GetIsIgnore() *string
	SetLang(v string) *DescribeInvadeEventListRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeInvadeEventListRequest
	GetMemberUid() *int64
	SetPageSize(v string) *DescribeInvadeEventListRequest
	GetPageSize() *string
	SetProcessStatusList(v []*int32) *DescribeInvadeEventListRequest
	GetProcessStatusList() []*int32
	SetRiskLevel(v []*int32) *DescribeInvadeEventListRequest
	GetRiskLevel() []*int32
	SetSourceIp(v string) *DescribeInvadeEventListRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInvadeEventListRequest
	GetStartTime() *string
}

type DescribeInvadeEventListRequest struct {
	// The IP address of the affected asset.
	//
	// example:
	//
	// 10.0.XX.XX
	AssetsIP *string `json:"AssetsIP,omitempty" xml:"AssetsIP,omitempty"`
	// The ID of the affected instance.
	//
	// example:
	//
	// ins_1321_asedb_****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the affected instance.
	//
	// example:
	//
	// ECS_test
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The number of the page to return.
	//
	// Default: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. This must be a UNIX timestamp in seconds. If you omit this parameter, the query defaults to the current time.
	//
	// example:
	//
	// 1656837360
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A unique identifier for the breach awareness event.
	//
	// example:
	//
	// 69d189e2-ec17-4676-a2fe-02969234****
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the breach awareness event.
	//
	// example:
	//
	// event_test
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The universally unique identifier (UUID) of the breach awareness event.
	//
	// example:
	//
	// fadd-dfdd-****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// Specifies whether to query for ignored breach awareness events. Valid values:
	//
	// - **true**: Ignored.
	//
	// - **false**: Not ignored.
	//
	// example:
	//
	// true
	IsIgnore *string `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 135809047715****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return per page.
	//
	// Default: 6. Maximum: 10.
	//
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array of processing statuses to filter events by. Only events with a status specified in this array are returned.
	//
	// example:
	//
	// 1358090477156271
	ProcessStatusList []*int32 `json:"ProcessStatusList,omitempty" xml:"ProcessStatusList,omitempty" type:"Repeated"`
	// An array of risk levels to filter events by. Only events with a risk level specified in this array are returned.
	RiskLevel []*int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The source IP address that initiated the event.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start of the time range to query. This must be a UNIX timestamp in seconds. If you omit this parameter, the query defaults to the last 30 days.
	//
	// example:
	//
	// 1656750960
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListRequest) GetAssetsIP() *string {
	return s.AssetsIP
}

func (s *DescribeInvadeEventListRequest) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventListRequest) GetAssetsInstanceName() *string {
	return s.AssetsInstanceName
}

func (s *DescribeInvadeEventListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInvadeEventListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInvadeEventListRequest) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeInvadeEventListRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeInvadeEventListRequest) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventListRequest) GetIsIgnore() *string {
	return s.IsIgnore
}

func (s *DescribeInvadeEventListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEventListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeInvadeEventListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInvadeEventListRequest) GetProcessStatusList() []*int32 {
	return s.ProcessStatusList
}

func (s *DescribeInvadeEventListRequest) GetRiskLevel() []*int32 {
	return s.RiskLevel
}

func (s *DescribeInvadeEventListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEventListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvadeEventListRequest) SetAssetsIP(v string) *DescribeInvadeEventListRequest {
	s.AssetsIP = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceId(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceName(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetCurrentPage(v string) *DescribeInvadeEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEndTime(v string) *DescribeInvadeEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventKey(v string) *DescribeInvadeEventListRequest {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventName(v string) *DescribeInvadeEventListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventUuid(v string) *DescribeInvadeEventListRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetIsIgnore(v string) *DescribeInvadeEventListRequest {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetLang(v string) *DescribeInvadeEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetMemberUid(v int64) *DescribeInvadeEventListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetPageSize(v string) *DescribeInvadeEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetProcessStatusList(v []*int32) *DescribeInvadeEventListRequest {
	s.ProcessStatusList = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetRiskLevel(v []*int32) *DescribeInvadeEventListRequest {
	s.RiskLevel = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetSourceIp(v string) *DescribeInvadeEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetStartTime(v string) *DescribeInvadeEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEventListRequest) Validate() error {
	return dara.Validate(s)
}
