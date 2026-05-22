// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSitesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSitesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSitesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSitesResponseBody
	GetRequestId() *string
	SetSites(v []*ListSitesResponseBodySites) *ListSitesResponseBody
	GetSites() []*ListSitesResponseBodySites
	SetTotalCount(v int32) *ListSitesResponseBody
	GetTotalCount() *int32
}

type ListSitesResponseBody struct {
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sites      []*ListSitesResponseBodySites `json:"Sites,omitempty" xml:"Sites,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSitesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSitesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSitesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSitesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSitesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSitesResponseBody) GetSites() []*ListSitesResponseBodySites {
	return s.Sites
}

func (s *ListSitesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSitesResponseBody) SetPageNumber(v int32) *ListSitesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSitesResponseBody) SetPageSize(v int32) *ListSitesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSitesResponseBody) SetRequestId(v string) *ListSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSitesResponseBody) SetSites(v []*ListSitesResponseBodySites) *ListSitesResponseBody {
	s.Sites = v
	return s
}

func (s *ListSitesResponseBody) SetTotalCount(v int32) *ListSitesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSitesResponseBody) Validate() error {
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

type ListSitesResponseBodySites struct {
	AccessType      *string                `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CnameZone       *string                `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	Coverage        *string                `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	CreateTime      *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId      *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NameServerList  *string                `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	OfflineReason   *string                `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	PlanName        *string                `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanSpecName    *string                `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SiteId          *int64                 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName        *string                `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status          *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UpdateTime      *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VerifyCode      *string                `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	VisitTime       *string                `json:"VisitTime,omitempty" xml:"VisitTime,omitempty"`
}

func (s ListSitesResponseBodySites) String() string {
	return dara.Prettify(s)
}

func (s ListSitesResponseBodySites) GoString() string {
	return s.String()
}

func (s *ListSitesResponseBodySites) GetAccessType() *string {
	return s.AccessType
}

func (s *ListSitesResponseBodySites) GetCnameZone() *string {
	return s.CnameZone
}

func (s *ListSitesResponseBodySites) GetCoverage() *string {
	return s.Coverage
}

func (s *ListSitesResponseBodySites) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSitesResponseBodySites) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSitesResponseBodySites) GetNameServerList() *string {
	return s.NameServerList
}

func (s *ListSitesResponseBodySites) GetOfflineReason() *string {
	return s.OfflineReason
}

func (s *ListSitesResponseBodySites) GetPlanName() *string {
	return s.PlanName
}

func (s *ListSitesResponseBodySites) GetPlanSpecName() *string {
	return s.PlanSpecName
}

func (s *ListSitesResponseBodySites) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSitesResponseBodySites) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSitesResponseBodySites) GetSiteName() *string {
	return s.SiteName
}

func (s *ListSitesResponseBodySites) GetStatus() *string {
	return s.Status
}

func (s *ListSitesResponseBodySites) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListSitesResponseBodySites) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSitesResponseBodySites) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *ListSitesResponseBodySites) GetVisitTime() *string {
	return s.VisitTime
}

func (s *ListSitesResponseBodySites) SetAccessType(v string) *ListSitesResponseBodySites {
	s.AccessType = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCnameZone(v string) *ListSitesResponseBodySites {
	s.CnameZone = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCoverage(v string) *ListSitesResponseBodySites {
	s.Coverage = &v
	return s
}

func (s *ListSitesResponseBodySites) SetCreateTime(v string) *ListSitesResponseBodySites {
	s.CreateTime = &v
	return s
}

func (s *ListSitesResponseBodySites) SetInstanceId(v string) *ListSitesResponseBodySites {
	s.InstanceId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetNameServerList(v string) *ListSitesResponseBodySites {
	s.NameServerList = &v
	return s
}

func (s *ListSitesResponseBodySites) SetOfflineReason(v string) *ListSitesResponseBodySites {
	s.OfflineReason = &v
	return s
}

func (s *ListSitesResponseBodySites) SetPlanName(v string) *ListSitesResponseBodySites {
	s.PlanName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetPlanSpecName(v string) *ListSitesResponseBodySites {
	s.PlanSpecName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetResourceGroupId(v string) *ListSitesResponseBodySites {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetSiteId(v int64) *ListSitesResponseBodySites {
	s.SiteId = &v
	return s
}

func (s *ListSitesResponseBodySites) SetSiteName(v string) *ListSitesResponseBodySites {
	s.SiteName = &v
	return s
}

func (s *ListSitesResponseBodySites) SetStatus(v string) *ListSitesResponseBodySites {
	s.Status = &v
	return s
}

func (s *ListSitesResponseBodySites) SetTags(v map[string]interface{}) *ListSitesResponseBodySites {
	s.Tags = v
	return s
}

func (s *ListSitesResponseBodySites) SetUpdateTime(v string) *ListSitesResponseBodySites {
	s.UpdateTime = &v
	return s
}

func (s *ListSitesResponseBodySites) SetVerifyCode(v string) *ListSitesResponseBodySites {
	s.VerifyCode = &v
	return s
}

func (s *ListSitesResponseBodySites) SetVisitTime(v string) *ListSitesResponseBodySites {
	s.VisitTime = &v
	return s
}

func (s *ListSitesResponseBodySites) Validate() error {
	return dara.Validate(s)
}
