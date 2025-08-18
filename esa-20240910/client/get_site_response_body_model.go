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
	// The queried website information.
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
	return dara.Validate(s)
}

type GetSiteResponseBodySiteModel struct {
	// The DNS setup option for the website. Valid values:
	//
	// 	- **NS**
	//
	// 	- **CNAME**
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The CNAME of the website domain. If you use CNAME setup when you add your website to ESA, the value is the CNAME that you configured then.
	//
	// example:
	//
	// example.cname.com
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// The service location. Valid values:
	//
	// 	- **domestic**: the Chinese mainland.
	//
	// 	- **global**: global.
	//
	// 	- **overseas**: outside the Chinese mainland.
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The time when the WEBsite was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The plan ID.
	//
	// example:
	//
	// cas-merge-q6h0bv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The nameservers assigned to the website domain. They are separated by commas (,).
	//
	// example:
	//
	// male1-1.ialicdn.com,female1-1.ialicdn.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	OfflineReason  *string `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	// The plan name.
	//
	// example:
	//
	// plan-168777532****
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The specification of the plan associated with the website.
	//
	// example:
	//
	// normal
	PlanSpecName *string `json:"PlanSpecName,omitempty" xml:"PlanSpecName,omitempty"`
	// The ID of your Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek26g6i6se****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The website status. Valid values:
	//
	// 	- **pending**: The website is to be configured.
	//
	// 	- **active**: The website is active.
	//
	// 	- **offline**: The website is suspended.
	//
	// 	- **moved**: The website has been added and verified by another Alibaba Cloud account.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the website.
	//
	// example:
	//
	// {"tag1":"value1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the WEBsite was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The information about custom nameservers of the website domain. The key is a custom nameserver name, and the value is the IP address of the custom nameserver. Multiple IP addresses are separated by commas (,).
	VanityNSList map[string]*string `json:"VanityNSList,omitempty" xml:"VanityNSList,omitempty"`
	// The code that is used to verify the website domain ownership. As part of the verification TXT record, this parameter is returned for websites that use CNAME setup.
	//
	// example:
	//
	// verify_d516cb3740f81f0cef77d162edd1****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// The status of version management. If true is returned, version management is enabled for the website.
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
