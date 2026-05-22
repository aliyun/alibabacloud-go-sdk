// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSiteResponseBody
	GetRequestId() *string
	SetSiteModel(v *GetSiteResponseBodySiteModel) *GetSiteResponseBody
	GetSiteModel() *GetSiteResponseBodySiteModel
}

type GetSiteResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteModel *GetSiteResponseBodySiteModel `json:"SiteModel,omitempty" xml:"SiteModel,omitempty" type:"Struct"`
}

func (s GetSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteResponseBody) GetSiteModel() *GetSiteResponseBodySiteModel {
	return s.SiteModel
}

func (s *GetSiteResponseBody) SetRequestId(v string) *GetSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteResponseBody) SetSiteModel(v *GetSiteResponseBodySiteModel) *GetSiteResponseBody {
	s.SiteModel = v
	return s
}

func (s *GetSiteResponseBody) Validate() error {
	if s.SiteModel != nil {
		if err := s.SiteModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSiteResponseBodySiteModel struct {
	AccessType        *string                `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CnameZone         *string                `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	Coverage          *string                `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	CreateTime        *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId        *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NameServerList    *string                `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	OfflineReason     *string                `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	PlanName          *string                `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanSpecName      *string                `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	ResourceGroupId   *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SiteId            *int64                 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName          *string                `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status            *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags              map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UpdateTime        *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VanityNSList      map[string]*string     `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
	VerifyCode        *string                `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	VersionManagement *bool                  `json:"VersionManagement,omitempty" xml:"VersionManagement,omitempty"`
}

func (s GetSiteResponseBodySiteModel) String() string {
	return dara.Prettify(s)
}

func (s GetSiteResponseBodySiteModel) GoString() string {
	return s.String()
}

func (s *GetSiteResponseBodySiteModel) GetAccessType() *string {
	return s.AccessType
}

func (s *GetSiteResponseBodySiteModel) GetCnameZone() *string {
	return s.CnameZone
}

func (s *GetSiteResponseBodySiteModel) GetCoverage() *string {
	return s.Coverage
}

func (s *GetSiteResponseBodySiteModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSiteResponseBodySiteModel) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSiteResponseBodySiteModel) GetNameServerList() *string {
	return s.NameServerList
}

func (s *GetSiteResponseBodySiteModel) GetOfflineReason() *string {
	return s.OfflineReason
}

func (s *GetSiteResponseBodySiteModel) GetPlanName() *string {
	return s.PlanName
}

func (s *GetSiteResponseBodySiteModel) GetPlanSpecName() *string {
	return s.PlanSpecName
}

func (s *GetSiteResponseBodySiteModel) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSiteResponseBodySiteModel) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteResponseBodySiteModel) GetSiteName() *string {
	return s.SiteName
}

func (s *GetSiteResponseBodySiteModel) GetStatus() *string {
	return s.Status
}

func (s *GetSiteResponseBodySiteModel) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetSiteResponseBodySiteModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSiteResponseBodySiteModel) GetVanityNSList() map[string]*string {
	return s.VanityNSList
}

func (s *GetSiteResponseBodySiteModel) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *GetSiteResponseBodySiteModel) GetVersionManagement() *bool {
	return s.VersionManagement
}

func (s *GetSiteResponseBodySiteModel) SetAccessType(v string) *GetSiteResponseBodySiteModel {
	s.AccessType = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCnameZone(v string) *GetSiteResponseBodySiteModel {
	s.CnameZone = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCoverage(v string) *GetSiteResponseBodySiteModel {
	s.Coverage = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetCreateTime(v string) *GetSiteResponseBodySiteModel {
	s.CreateTime = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetInstanceId(v string) *GetSiteResponseBodySiteModel {
	s.InstanceId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetNameServerList(v string) *GetSiteResponseBodySiteModel {
	s.NameServerList = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetOfflineReason(v string) *GetSiteResponseBodySiteModel {
	s.OfflineReason = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetPlanName(v string) *GetSiteResponseBodySiteModel {
	s.PlanName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetPlanSpecName(v string) *GetSiteResponseBodySiteModel {
	s.PlanSpecName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetResourceGroupId(v string) *GetSiteResponseBodySiteModel {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetSiteId(v int64) *GetSiteResponseBodySiteModel {
	s.SiteId = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetSiteName(v string) *GetSiteResponseBodySiteModel {
	s.SiteName = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetStatus(v string) *GetSiteResponseBodySiteModel {
	s.Status = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetTags(v map[string]interface{}) *GetSiteResponseBodySiteModel {
	s.Tags = v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetUpdateTime(v string) *GetSiteResponseBodySiteModel {
	s.UpdateTime = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVanityNSList(v map[string]*string) *GetSiteResponseBodySiteModel {
	s.VanityNSList = v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVerifyCode(v string) *GetSiteResponseBodySiteModel {
	s.VerifyCode = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) SetVersionManagement(v bool) *GetSiteResponseBodySiteModel {
	s.VersionManagement = &v
	return s
}

func (s *GetSiteResponseBodySiteModel) Validate() error {
	return dara.Validate(s)
}
