// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidRatePlanInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) *ListPostpaidRatePlanInstancesResponseBody
	GetInstanceInfo() []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfo
	SetPageNumber(v int32) *ListPostpaidRatePlanInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPostpaidRatePlanInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListPostpaidRatePlanInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPostpaidRatePlanInstancesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListPostpaidRatePlanInstancesResponseBody
	GetTotalPage() *int32
}

type ListPostpaidRatePlanInstancesResponseBody struct {
	InstanceInfo []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	PageNumber   *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListPostpaidRatePlanInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidRatePlanInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetInstanceInfo() []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPostpaidRatePlanInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetInstanceInfo(v []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) *ListPostpaidRatePlanInstancesResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetPageNumber(v int32) *ListPostpaidRatePlanInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetPageSize(v int32) *ListPostpaidRatePlanInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetRequestId(v string) *ListPostpaidRatePlanInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetTotalCount(v int32) *ListPostpaidRatePlanInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) SetTotalPage(v int32) *ListPostpaidRatePlanInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBody) Validate() error {
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

type ListPostpaidRatePlanInstancesResponseBodyInstanceInfo struct {
	BillingMethod      *string                                                       `json:"BillingMethod,omitempty" xml:"BillingMethod,omitempty"`
	BillingMode        *string                                                       `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	Coverages          *string                                                       `json:"Coverages,omitempty" xml:"Coverages,omitempty"`
	CreateTime         *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpectedUpdateTime *string                                                       `json:"ExpectedUpdateTime,omitempty" xml:"ExpectedUpdateTime,omitempty"`
	InstanceId         *string                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PlanName           *string                                                       `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanNameCn         *string                                                       `json:"PlanNameCn,omitempty" xml:"PlanNameCn,omitempty"`
	PlanType           *string                                                       `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	SiteQuota          *string                                                       `json:"SiteQuota,omitempty" xml:"SiteQuota,omitempty"`
	Sites              []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
	Status             *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetBillingMode() *string {
	return s.BillingMode
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetCoverages() *string {
	return s.Coverages
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetExpectedUpdateTime() *string {
	return s.ExpectedUpdateTime
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetPlanName() *string {
	return s.PlanName
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetPlanNameCn() *string {
	return s.PlanNameCn
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetPlanType() *string {
	return s.PlanType
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetSiteQuota() *string {
	return s.SiteQuota
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetSites() []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites {
	return s.Sites
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetBillingMethod(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.BillingMethod = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetBillingMode(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.BillingMode = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetCoverages(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.Coverages = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetCreateTime(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetExpectedUpdateTime(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.ExpectedUpdateTime = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetInstanceId(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetPlanName(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanName = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetPlanNameCn(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanNameCn = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetPlanType(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.PlanType = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetSiteQuota(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.SiteQuota = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetSites(v []*ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.Sites = v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) SetStatus(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfo) Validate() error {
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

type ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites struct {
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	SiteStatus *string `json:"SiteStatus,omitempty" xml:"SiteStatus,omitempty"`
}

func (s ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) GoString() string {
	return s.String()
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteName() *string {
	return s.SiteName
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) GetSiteStatus() *string {
	return s.SiteStatus
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteId(v int64) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteId = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteName(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteName = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) SetSiteStatus(v string) *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites {
	s.SiteStatus = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponseBodyInstanceInfoSites) Validate() error {
	return dara.Validate(s)
}
