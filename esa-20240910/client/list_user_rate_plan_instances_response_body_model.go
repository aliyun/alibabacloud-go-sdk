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
	InstanceInfo []*ListUserRatePlanInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	PageNumber   *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage    *int32                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	BillingMode                     *string                                                   `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	BotInstanceLevel                *string                                                   `json:"BotInstanceLevel,omitempty" xml:"BotInstanceLevel,omitempty"`
	BotRequest                      *string                                                   `json:"BotRequest,omitempty" xml:"BotRequest,omitempty"`
	Coverages                       *string                                                   `json:"Coverages,omitempty" xml:"Coverages,omitempty"`
	CreateTime                      *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossborderTraffic              *string                                                   `json:"CrossborderTraffic,omitempty" xml:"CrossborderTraffic,omitempty"`
	DdosBurstableDomesticProtection *string                                                   `json:"DdosBurstableDomesticProtection,omitempty" xml:"DdosBurstableDomesticProtection,omitempty"`
	DdosBurstableOverseasProtection *string                                                   `json:"DdosBurstableOverseasProtection,omitempty" xml:"DdosBurstableOverseasProtection,omitempty"`
	DdosInstanceLevel               *string                                                   `json:"DdosInstanceLevel,omitempty" xml:"DdosInstanceLevel,omitempty"`
	Duration                        *int32                                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EdgeRoutineRquest               *string                                                   `json:"EdgeRoutineRquest,omitempty" xml:"EdgeRoutineRquest,omitempty"`
	EdgeWafRequest                  *string                                                   `json:"EdgeWafRequest,omitempty" xml:"EdgeWafRequest,omitempty"`
	ExpireTime                      *string                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId                      *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Layer4Traffic                   *string                                                   `json:"Layer4Traffic,omitempty" xml:"Layer4Traffic,omitempty"`
	Layer4TrafficIntl               *string                                                   `json:"Layer4TrafficIntl,omitempty" xml:"Layer4TrafficIntl,omitempty"`
	PlanName                        *string                                                   `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanTraffic                     *string                                                   `json:"PlanTraffic,omitempty" xml:"PlanTraffic,omitempty"`
	PlanType                        *string                                                   `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	SiteQuota                       *string                                                   `json:"SiteQuota,omitempty" xml:"SiteQuota,omitempty"`
	Sites                           []*ListUserRatePlanInstancesResponseBodyInstanceInfoSites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
	SmartRoutingRequest             *string                                                   `json:"SmartRoutingRequest,omitempty" xml:"SmartRoutingRequest,omitempty"`
	StaticRequest                   *string                                                   `json:"StaticRequest,omitempty" xml:"StaticRequest,omitempty"`
	Status                          *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscribeType                   *string                                                   `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
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
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
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
