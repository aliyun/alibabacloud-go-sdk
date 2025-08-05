// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*DescribeInvadeEventListResponseBodyEventList) *DescribeInvadeEventListResponseBody
	GetEventList() []*DescribeInvadeEventListResponseBodyEventList
	SetHighLevelPercent(v int32) *DescribeInvadeEventListResponseBody
	GetHighLevelPercent() *int32
	SetLowLevelPercent(v int32) *DescribeInvadeEventListResponseBody
	GetLowLevelPercent() *int32
	SetMiddleLevelPercent(v int32) *DescribeInvadeEventListResponseBody
	GetMiddleLevelPercent() *int32
	SetPageInfo(v *DescribeInvadeEventListResponseBodyPageInfo) *DescribeInvadeEventListResponseBody
	GetPageInfo() *DescribeInvadeEventListResponseBodyPageInfo
	SetRequestId(v string) *DescribeInvadeEventListResponseBody
	GetRequestId() *string
}

type DescribeInvadeEventListResponseBody struct {
	// An array that consists of breach awareness events.
	EventList []*DescribeInvadeEventListResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The percentage of high-risk events.
	//
	// example:
	//
	// 40
	HighLevelPercent *int32 `json:"HighLevelPercent,omitempty" xml:"HighLevelPercent,omitempty"`
	// The percentage of low-risk events.
	//
	// example:
	//
	// 20
	LowLevelPercent *int32 `json:"LowLevelPercent,omitempty" xml:"LowLevelPercent,omitempty"`
	// The percentage of medium-risk events.
	//
	// example:
	//
	// 40
	MiddleLevelPercent *int32 `json:"MiddleLevelPercent,omitempty" xml:"MiddleLevelPercent,omitempty"`
	// The pagination information.
	PageInfo *DescribeInvadeEventListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvadeEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBody) GetEventList() []*DescribeInvadeEventListResponseBodyEventList {
	return s.EventList
}

func (s *DescribeInvadeEventListResponseBody) GetHighLevelPercent() *int32 {
	return s.HighLevelPercent
}

func (s *DescribeInvadeEventListResponseBody) GetLowLevelPercent() *int32 {
	return s.LowLevelPercent
}

func (s *DescribeInvadeEventListResponseBody) GetMiddleLevelPercent() *int32 {
	return s.MiddleLevelPercent
}

func (s *DescribeInvadeEventListResponseBody) GetPageInfo() *DescribeInvadeEventListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInvadeEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEventListResponseBody) SetEventList(v []*DescribeInvadeEventListResponseBodyEventList) *DescribeInvadeEventListResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetHighLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.HighLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetLowLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.LowLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetMiddleLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.MiddleLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetPageInfo(v *DescribeInvadeEventListResponseBodyPageInfo) *DescribeInvadeEventListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetRequestId(v string) *DescribeInvadeEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInvadeEventListResponseBodyEventList struct {
	// The ID of the affected asset.
	//
	// example:
	//
	// i-ECS****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the affected asset.
	//
	// example:
	//
	// ECS_test
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// 	- **BastionHostIP**: the egress IP address of a bastion host
	//
	// 	- **BastionHostIngressIP**: the ingress IP address of a bastion host
	//
	// 	- **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	//
	// 	- **EcsPublicIP**: the public IP address of an ECS instance
	//
	// 	- **EIP**: the EIP
	//
	// 	- **EniEIP**: the EIP of an elastic network interface (ENI)
	//
	// 	- **NatEIP**: the EIP of a NAT gateway
	//
	// 	- **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	//
	// 	- **SlbPublicIP**: the public IP address of an SLB instance
	//
	// 	- **NatPublicIP**: the public IP address of a NAT gateway
	//
	// 	- **HAVIP**: the high-availability virtual IP address (HAVIP)
	//
	// example:
	//
	// EcsPublicIp
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The ID of the breach awareness event.
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
	// The type of the breach awareness event. Valid values:
	//
	// 	- **IPS**: intrusion prevention event
	//
	// 	- **offline**: disconnection event
	//
	// example:
	//
	// IPS
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The UUID of the breach awareness event.
	//
	// example:
	//
	// fadd-dfdd-****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The time when the breach awareness event first occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1656750960
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether the breach awareness event is ignored. Valid values:
	//
	// 	- **true**: The breach awareness event is ignored.
	//
	// 	- **false**: The breach awareness event is not ignored.
	//
	// example:
	//
	// true
	IsIgnore *bool `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The time when the breach awareness event last occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1656837360
	LastTime *int32 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The private IP address of the affected asset.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The handling status of the breach awareness event. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **20**: handled
	//
	// example:
	//
	// 20
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The public IP address of the affected asset.
	//
	// example:
	//
	// 198.51.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// 	- **BastionHostIP**: the egress IP address of a bastion host
	//
	// 	- **BastionHostIngressIP**: the ingress IP address of a bastion host
	//
	// 	- **EcsEIP**: the EIP of an ECS instance
	//
	// 	- **EcsPublicIP**: the public IP address of an ECS instance
	//
	// 	- **EIP**: the EIP
	//
	// 	- **EniEIP**: the EIP of an ENI
	//
	// 	- **NatEIP**: the EIP of a NAT gateway
	//
	// 	- **SlbEIP**: the EIP of an SLB instance
	//
	// 	- **SlbPublicIP**: the public IP address of an SLB instance
	//
	// 	- **NatPublicIP**: the public IP address of a NAT gateway
	//
	// 	- **HAVIP**: the HAVIP
	//
	// example:
	//
	// EcsPublicIp
	PublicIpType *string `json:"PublicIpType,omitempty" xml:"PublicIpType,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetAssetsInstanceName() *string {
	return s.AssetsInstanceName
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetAssetsType() *string {
	return s.AssetsType
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetEventSrc() *string {
	return s.EventSrc
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetFirstTime() *int32 {
	return s.FirstTime
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetIsIgnore() *bool {
	return s.IsIgnore
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetLastTime() *int32 {
	return s.LastTime
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetPublicIpType() *string {
	return s.PublicIpType
}

func (s *DescribeInvadeEventListResponseBodyEventList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceId(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventKey(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventSrc(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventSrc = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventUuid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetFirstTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.FirstTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetIsIgnore(v bool) *DescribeInvadeEventListResponseBodyEventList {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetLastTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.LastTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetMemberUid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPrivateIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetProcessStatus(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIpType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIpType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetRiskLevel(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}

type DescribeInvadeEventListResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of breach awareness events.
	//
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetPageSize(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
