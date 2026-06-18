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
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the site.
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
	// The access type of the site. Valid values:
	//
	// - **NS**: Access via NS.
	//
	// - **CNAME**: Access via CNAME.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// For sites onboarded via CNAME, use this suffix to configure the CNAME record.
	//
	// example:
	//
	// example.cname.com
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// The acceleration region. Valid values:
	//
	// - **domestic**: Chinese mainland only
	//
	// - **global**: Global
	//
	// - **overseas**: Global (excluding the Chinese mainland)
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The time (in UTC) when the site was created, formatted in ISO 8601 (`yyyy-MM-ddTHH:mm:ssZ`).
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the plan instance.
	//
	// example:
	//
	// cas-merge-q6h0bv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A comma-separated list of name servers assigned to the site.
	//
	// example:
	//
	// male1-1.ialicdn.com,female1-1.ialicdn.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// The reason the site is offline. This parameter appears only when `Status` is `offline`. Valid values:
	//
	// - **expiration_arrears**: The subscription plan has expired or the account has overdue payments.
	//
	// - **internally_disabled**: The site was disabled by the system.
	//
	// - **missing_icp**: The domain is missing an ICP license.
	//
	// - **content_violation**: The site violated content policies.
	//
	// - **proactively_disabled**: The site was disabled either by you or by a usage limit that you configured.
	//
	// example:
	//
	// expiration_ arrears
	OfflineReason *string `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	// The name of the plan.
	//
	// example:
	//
	// plan-168777532****
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The name of the plan specification.
	//
	// example:
	//
	// normal
	PlanSpecName *string `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek26g6i6se****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the site.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The name of the site.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The status of the site. Valid values:
	//
	// - **pending**: The site is pending configuration.
	//
	// - **active**: The site is active.
	//
	// - **offline**: The site is offline.
	//
	// - **moved**: The site has been superseded.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the site.
	//
	// example:
	//
	// {"tag1":"value1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time (in UTC) when the site was last updated, formatted in ISO 8601 (`yyyy-MM-ddTHH:mm:ssZ`).
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Each key is a custom name server, and its value is a comma-separated list of the server\\"s IP addresses.
	VanityNSList map[string]*string `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
	// For sites onboarded via CNAME, you must configure this code as a TXT record.
	//
	// example:
	//
	// verify_d516cb3740f81f0cef77d162edd1****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// If `true`, version management is enabled for the site.
	//
	// example:
	//
	// true
	VersionManagement *bool `json:"VersionManagement,omitempty" xml:"VersionManagement,omitempty"`
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
