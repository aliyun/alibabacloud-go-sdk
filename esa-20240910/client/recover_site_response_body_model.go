// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *RecoverSiteResponseBody
	GetAccessType() *string
	SetCnameZone(v string) *RecoverSiteResponseBody
	GetCnameZone() *string
	SetCoverage(v string) *RecoverSiteResponseBody
	GetCoverage() *string
	SetCreateTime(v string) *RecoverSiteResponseBody
	GetCreateTime() *string
	SetInstanceId(v string) *RecoverSiteResponseBody
	GetInstanceId() *string
	SetNameServerList(v string) *RecoverSiteResponseBody
	GetNameServerList() *string
	SetOfflineReason(v string) *RecoverSiteResponseBody
	GetOfflineReason() *string
	SetPlanName(v string) *RecoverSiteResponseBody
	GetPlanName() *string
	SetRequestId(v string) *RecoverSiteResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *RecoverSiteResponseBody
	GetResourceGroupId() *string
	SetSiteId(v int64) *RecoverSiteResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *RecoverSiteResponseBody
	GetSiteName() *string
	SetStatus(v string) *RecoverSiteResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *RecoverSiteResponseBody
	GetUpdateTime() *string
	SetVerifyCode(v string) *RecoverSiteResponseBody
	GetVerifyCode() *string
}

type RecoverSiteResponseBody struct {
	// The access type. Valid values:
	//
	// - **NS**: access through NS hosting.
	//
	// - **CNAME**: access through CNAME.
	//
	// example:
	//
	// NS
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The CNAME suffix of the site. For sites that are accessed through CNAME, this field indicates the CNAME suffix that needs to be configured for records.
	//
	// example:
	//
	// gf-test.hkrt.cn
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// The acceleration region of the site. Valid values:
	//
	// - **domestic**: the Chinese mainland only.
	//
	// - **global**: global.
	//
	// - **overseas**: global (excluding the Chinese mainland).
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The plan instance ID.
	//
	// example:
	//
	// esa-site-9vjienwn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of name servers assigned to the site, separated by commas (,). When the site uses NS access, this field contains values. You need to change the DNS servers of the site to these name servers. Then you can verify the site ownership and activate the site.
	//
	// example:
	//
	// ns1.example.com,ns2.example.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// The reason why the site was deactivated. Valid values:
	//
	// - **expiration_arrears**: The subscription plan expired or the account has an overdue payment.
	//
	// - **internally_disabled**: The site was disabled by the system.
	//
	// - **missing_icp**: The domain name does not have an ICP filing.
	//
	// - **content_violation**: Content violation.
	//
	// - **proactively_disabled**: You proactively disabled the site or the site was disabled because the usage cap you configured was reached.
	//
	// example:
	//
	// expiration_ arrears
	OfflineReason *string `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	// The plan name.
	//
	// example:
	//
	// basic
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-axxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// - **moved**: The site has been replaced.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-03-13T02:13:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The site ownership verification code. When the site is accessed through CNAME, this is the TXT verification code that needs to be configured.
	//
	// example:
	//
	// verify_d516cb3740f81f0cef77d162edd1****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s RecoverSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverSiteResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverSiteResponseBody) GetAccessType() *string {
	return s.AccessType
}

func (s *RecoverSiteResponseBody) GetCnameZone() *string {
	return s.CnameZone
}

func (s *RecoverSiteResponseBody) GetCoverage() *string {
	return s.Coverage
}

func (s *RecoverSiteResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RecoverSiteResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RecoverSiteResponseBody) GetNameServerList() *string {
	return s.NameServerList
}

func (s *RecoverSiteResponseBody) GetOfflineReason() *string {
	return s.OfflineReason
}

func (s *RecoverSiteResponseBody) GetPlanName() *string {
	return s.PlanName
}

func (s *RecoverSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverSiteResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RecoverSiteResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *RecoverSiteResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *RecoverSiteResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RecoverSiteResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *RecoverSiteResponseBody) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *RecoverSiteResponseBody) SetAccessType(v string) *RecoverSiteResponseBody {
	s.AccessType = &v
	return s
}

func (s *RecoverSiteResponseBody) SetCnameZone(v string) *RecoverSiteResponseBody {
	s.CnameZone = &v
	return s
}

func (s *RecoverSiteResponseBody) SetCoverage(v string) *RecoverSiteResponseBody {
	s.Coverage = &v
	return s
}

func (s *RecoverSiteResponseBody) SetCreateTime(v string) *RecoverSiteResponseBody {
	s.CreateTime = &v
	return s
}

func (s *RecoverSiteResponseBody) SetInstanceId(v string) *RecoverSiteResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RecoverSiteResponseBody) SetNameServerList(v string) *RecoverSiteResponseBody {
	s.NameServerList = &v
	return s
}

func (s *RecoverSiteResponseBody) SetOfflineReason(v string) *RecoverSiteResponseBody {
	s.OfflineReason = &v
	return s
}

func (s *RecoverSiteResponseBody) SetPlanName(v string) *RecoverSiteResponseBody {
	s.PlanName = &v
	return s
}

func (s *RecoverSiteResponseBody) SetRequestId(v string) *RecoverSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverSiteResponseBody) SetResourceGroupId(v string) *RecoverSiteResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *RecoverSiteResponseBody) SetSiteId(v int64) *RecoverSiteResponseBody {
	s.SiteId = &v
	return s
}

func (s *RecoverSiteResponseBody) SetSiteName(v string) *RecoverSiteResponseBody {
	s.SiteName = &v
	return s
}

func (s *RecoverSiteResponseBody) SetStatus(v string) *RecoverSiteResponseBody {
	s.Status = &v
	return s
}

func (s *RecoverSiteResponseBody) SetUpdateTime(v string) *RecoverSiteResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *RecoverSiteResponseBody) SetVerifyCode(v string) *RecoverSiteResponseBody {
	s.VerifyCode = &v
	return s
}

func (s *RecoverSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
