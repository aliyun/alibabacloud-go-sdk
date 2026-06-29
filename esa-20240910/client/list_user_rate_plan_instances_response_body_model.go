// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRatePlanInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v []*ListUserRatePlanInstancesResponseBodyInstanceInfo) *ListUserRatePlanInstancesResponseBody
	GetInstanceInfo() []*ListUserRatePlanInstancesResponseBodyInstanceInfo
	SetPageNumber(v int32) *ListUserRatePlanInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserRatePlanInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserRatePlanInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserRatePlanInstancesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListUserRatePlanInstancesResponseBody
	GetTotalPage() *int32
}

type ListUserRatePlanInstancesResponseBody struct {
	// The plan instances that match the specified conditions under the user.
	InstanceInfo []*ListUserRatePlanInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListUserRatePlanInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserRatePlanInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRatePlanInstancesResponseBody) GetInstanceInfo() []*ListUserRatePlanInstancesResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *ListUserRatePlanInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserRatePlanInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserRatePlanInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserRatePlanInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserRatePlanInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListUserRatePlanInstancesResponseBody) SetInstanceInfo(v []*ListUserRatePlanInstancesResponseBodyInstanceInfo) *ListUserRatePlanInstancesResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) SetPageNumber(v int32) *ListUserRatePlanInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) SetPageSize(v int32) *ListUserRatePlanInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) SetRequestId(v string) *ListUserRatePlanInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) SetTotalCount(v int32) *ListUserRatePlanInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) SetTotalPage(v int32) *ListUserRatePlanInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBody) Validate() error {
	if s.InstanceInfo != nil {
		for _, item := range s.InstanceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserRatePlanInstancesResponseBodyInstanceInfo struct {
	// The billing method. Valid values:
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	BillingMode *string `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	// If this field is empty, the plan does not include a bot protection instance. If a value is returned, the plan includes a bot protection instance. Valid values:
	//
	// - enterprise_bot: Web Edition
	//
	// - enterprise_bot_with_app: App Edition.
	//
	// example:
	//
	// enterprise_bot
	BotInstanceLevel *string `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
	// The prepaid bot protection requests included in the plan, in units of 10,000.
	//
	// example:
	//
	// 100
	BotRequest *string `json:"BotRequest,omitempty" xml:"BotRequest,omitempty"`
	// The acceleration regions to which sites can be bound under this plan instance. Multiple values are separated by commas (,). Valid values:
	//
	// - **domestic**: China or the Chinese mainland.
	//
	// - **overseas**: Global (excluding China or the Chinese mainland).
	//
	// - **global**: Global (including China or the Chinese mainland).
	//
	// example:
	//
	// domestic,overseas
	Coverages *string `json:"Coverages,omitempty" xml:"Coverages,omitempty"`
	// The purchase time of the plan instance. The time is in ISO 8601 format and displayed in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The prepaid China network acceleration traffic included in the plan, in GB.
	//
	// example:
	//
	// 100
	CrossborderTraffic *string `json:"CrossborderTraffic,omitempty" xml:"CrossborderTraffic,omitempty"`
	// The Anti-DDoS instance specification for the Chinese mainland included in the plan.
	//
	// example:
	//
	// cn_300
	DdosBurstableDomesticProtection *string `json:"DdosBurstableDomesticProtection,omitempty" xml:"DdosBurstableDomesticProtection,omitempty"`
	// The Anti-DDoS instance specification outside the Chinese mainland included in the plan.
	//
	// example:
	//
	// overseas_300
	DdosBurstableOverseasProtection *string `json:"DdosBurstableOverseasProtection,omitempty" xml:"DdosBurstableOverseasProtection,omitempty"`
	// If this field is empty, the plan does not include an Anti-DDoS instance. If a value is returned, the plan includes an Anti-DDoS instance. The value is `esa_ddos_instance`.
	//
	// example:
	//
	// esa_ddos_instance
	DdosInstanceLevel *string `json:"DdosInstanceLevel,omitempty" xml:"DdosInstanceLevel,omitempty"`
	// The subscription duration of the plan instance. Unit: months.
	//
	// example:
	//
	// 3
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The prepaid Edge Routine (ER) requests included in the plan, in units of 10,000.
	//
	// example:
	//
	// 100
	EdgeRoutineRquest *string `json:"EdgeRoutineRquest,omitempty" xml:"EdgeRoutineRquest,omitempty"`
	// The prepaid WAF requests included in the plan, in units of 10,000.
	//
	// example:
	//
	// 100
	EdgeWafRequest *string `json:"EdgeWafRequest,omitempty" xml:"EdgeWafRequest,omitempty"`
	// The expiration time of the plan instance. The time is in ISO 8601 format and displayed in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The plan instance ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The prepaid Layer 4 proxy traffic included in the plan, in GB, for the Chinese mainland.
	//
	// example:
	//
	// 100
	Layer4Traffic *string `json:"Layer4Traffic,omitempty" xml:"Layer4Traffic,omitempty"`
	// The prepaid Layer 4 proxy traffic included in the plan, in GB, outside the Chinese mainland.
	//
	// example:
	//
	// 100
	Layer4TrafficIntl *string `json:"Layer4TrafficIntl,omitempty" xml:"Layer4TrafficIntl,omitempty"`
	// The plan name associated with the plan instance.
	//
	// example:
	//
	// basic
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The prepaid Layer 7 acceleration traffic included in the plan, in GB.
	//
	// example:
	//
	// 100
	PlanTraffic *string `json:"PlanTraffic,omitempty" xml:"PlanTraffic,omitempty"`
	// The plan type associated with the plan instance. Valid values:
	//
	// - **normal**: fixed-version plan.
	//
	// - **enterprise**: Enterprise Edition plan.
	//
	// example:
	//
	// normal
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The auto-renewal cycle. Unit: months.
	//
	// example:
	//
	// 6
	RenewalDuration *int64 `json:"RenewalDuration,omitempty" xml:"RenewalDuration,omitempty"`
	// The auto-renewal status. Valid values:
	//
	// - nomal: normal
	//
	// - auto_renewal: auto-renewal enabled
	//
	// - not_renewal: auto-renewal disabled.
	//
	// example:
	//
	// nomal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The site quota for the plan instance.
	//
	// example:
	//
	// 1
	SiteQuota *string `json:"SiteQuota,omitempty" xml:"SiteQuota,omitempty"`
	// The list of sites bound to the current plan instance.
	Sites []*ListUserRatePlanInstancesResponseBodyInstanceInfoSites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
	// The prepaid smart routing requests included in the plan, in units of 10,000.
	//
	// example:
	//
	// 100
	SmartRoutingRequest *string `json:"SmartRoutingRequest,omitempty" xml:"SmartRoutingRequest,omitempty"`
	// The prepaid HTTP requests included in the plan, in units of 10,000.
	//
	// example:
	//
	// 100
	StaticRequest *string `json:"StaticRequest,omitempty" xml:"StaticRequest,omitempty"`
	// The instance status. Valid values:
	//
	// - **online**: The plan instance is in normal service.
	//
	// - **offline**: The plan instance has expired but has not exceeded the grace period and is not active.
	//
	// - **disable**: The plan instance has been released.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The plan subscription type. Valid values:
	//
	// - entranceplan: Free Edition (Chinese mainland)
	//
	// - entranceplan_intl: Free Edition (International)
	//
	// - basicplan: Basic Edition
	//
	// - standardplan: Standard Edition
	//
	// - advancedplan: Premium Edition
	//
	// - enterpriseplan: Enterprise Edition.
	//
	// example:
	//
	// basicplan
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s ListUserRatePlanInstancesResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserRatePlanInstancesResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetBillingMode() *string {
	return s.BillingMode
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetBotInstanceLevel() *string {
	return s.BotInstanceLevel
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetBotRequest() *string {
	return s.BotRequest
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetCoverages() *string {
	return s.Coverages
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetCrossborderTraffic() *string {
	return s.CrossborderTraffic
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetDdosBurstableDomesticProtection() *string {
	return s.DdosBurstableDomesticProtection
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetDdosBurstableOverseasProtection() *string {
	return s.DdosBurstableOverseasProtection
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetDdosInstanceLevel() *string {
	return s.DdosInstanceLevel
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetDuration() *int32 {
	return s.Duration
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetEdgeRoutineRquest() *string {
	return s.EdgeRoutineRquest
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetEdgeWafRequest() *string {
	return s.EdgeWafRequest
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetLayer4Traffic() *string {
	return s.Layer4Traffic
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetLayer4TrafficIntl() *string {
	return s.Layer4TrafficIntl
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetPlanName() *string {
	return s.PlanName
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetPlanTraffic() *string {
	return s.PlanTraffic
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetPlanType() *string {
	return s.PlanType
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetRenewalDuration() *int64 {
	return s.RenewalDuration
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetSiteQuota() *string {
	return s.SiteQuota
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetSites() []*ListUserRatePlanInstancesResponseBodyInstanceInfoSites {
	return s.Sites
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetSmartRoutingRequest() *string {
	return s.SmartRoutingRequest
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetStaticRequest() *string {
	return s.StaticRequest
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetBillingMode(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.BillingMode = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetBotInstanceLevel(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.BotInstanceLevel = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetBotRequest(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.BotRequest = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetCoverages(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Coverages = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetCreateTime(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetCrossborderTraffic(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.CrossborderTraffic = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetDdosBurstableDomesticProtection(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.DdosBurstableDomesticProtection = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetDdosBurstableOverseasProtection(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.DdosBurstableOverseasProtection = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetDdosInstanceLevel(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.DdosInstanceLevel = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetDuration(v int32) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Duration = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetEdgeRoutineRquest(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.EdgeRoutineRquest = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetEdgeWafRequest(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.EdgeWafRequest = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetExpireTime(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetInstanceId(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetLayer4Traffic(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Layer4Traffic = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetLayer4TrafficIntl(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Layer4TrafficIntl = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetPlanName(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanName = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetPlanTraffic(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanTraffic = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetPlanType(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanType = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetRenewalDuration(v int64) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.RenewalDuration = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetRenewalStatus(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.RenewalStatus = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetSiteQuota(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.SiteQuota = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetSites(v []*ListUserRatePlanInstancesResponseBodyInstanceInfoSites) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Sites = v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetSmartRoutingRequest(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.SmartRoutingRequest = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetStaticRequest(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.StaticRequest = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetStatus(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) SetSubscribeType(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfo {
	s.SubscribeType = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfo) Validate() error {
	if s.Sites != nil {
		for _, item := range s.Sites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserRatePlanInstancesResponseBodyInstanceInfoSites struct {
	// The site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The site status. Valid values:
	//
	// - **pending**: The site is pending configuration.
	//
	// - **active**: The site is activated.
	//
	// - **offline**: The site is offline.
	//
	// - **moved**: The site has been superseded.
	//
	// example:
	//
	// pending
	SiteStatus *string `json:"SiteStatus,omitempty" xml:"SiteStatus,omitempty"`
}

func (s ListUserRatePlanInstancesResponseBodyInstanceInfoSites) String() string {
	return dara.Prettify(s)
}

func (s ListUserRatePlanInstancesResponseBodyInstanceInfoSites) GoString() string {
	return s.String()
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteName() *string {
	return s.SiteName
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteStatus() *string {
	return s.SiteStatus
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteId(v int64) *ListUserRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteId = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteName(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteName = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteStatus(v string) *ListUserRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteStatus = &v
	return s
}

func (s *ListUserRatePlanInstancesResponseBodyInstanceInfoSites) Validate() error {
	return dara.Validate(s)
}
