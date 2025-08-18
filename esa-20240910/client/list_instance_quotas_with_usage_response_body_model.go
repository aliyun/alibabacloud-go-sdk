// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasWithUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceQuotasWithUsageResponseBody
	GetInstanceId() *string
	SetQuotas(v []*ListInstanceQuotasWithUsageResponseBodyQuotas) *ListInstanceQuotasWithUsageResponseBody
	GetQuotas() []*ListInstanceQuotasWithUsageResponseBodyQuotas
	SetRequestId(v string) *ListInstanceQuotasWithUsageResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListInstanceQuotasWithUsageResponseBody
	GetStatus() *string
}

type ListInstanceQuotasWithUsageResponseBody struct {
	// The plan ID.[](~~2850189~~)
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quotas and their actual usage in the plan.
	Quotas []*ListInstanceQuotasWithUsageResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 85H66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The plan status. Valid values:
	//
	// 	- online: The plan is in service.
	//
	// 	- offline: The plan has expired within an allowable period. In this state, the plan is unavailable.
	//
	// 	- disable: The plan is released.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceQuotasWithUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasWithUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasWithUsageResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceQuotasWithUsageResponseBody) GetQuotas() []*ListInstanceQuotasWithUsageResponseBodyQuotas {
	return s.Quotas
}

func (s *ListInstanceQuotasWithUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceQuotasWithUsageResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceQuotasWithUsageResponseBody) SetInstanceId(v string) *ListInstanceQuotasWithUsageResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBody) SetQuotas(v []*ListInstanceQuotasWithUsageResponseBodyQuotas) *ListInstanceQuotasWithUsageResponseBody {
	s.Quotas = v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBody) SetRequestId(v string) *ListInstanceQuotasWithUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBody) SetStatus(v string) *ListInstanceQuotasWithUsageResponseBody {
	s.Status = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceQuotasWithUsageResponseBodyQuotas struct {
	// The quota name.
	//
	// example:
	//
	// redirect_rules|rule_quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 10
	QuotaValue *string `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The usage of the quota in each website associated with the plan.
	SiteUsage []*ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty" type:"Repeated"`
	// The quota usage.
	//
	// example:
	//
	// 3
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListInstanceQuotasWithUsageResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasWithUsageResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) GetQuotaValue() *string {
	return s.QuotaValue
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) GetSiteUsage() []*ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage {
	return s.SiteUsage
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) GetUsage() *string {
	return s.Usage
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) SetQuotaName(v string) *ListInstanceQuotasWithUsageResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) SetQuotaValue(v string) *ListInstanceQuotasWithUsageResponseBodyQuotas {
	s.QuotaValue = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) SetSiteUsage(v []*ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) *ListInstanceQuotasWithUsageResponseBodyQuotas {
	s.SiteUsage = v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) SetUsage(v string) *ListInstanceQuotasWithUsageResponseBodyQuotas {
	s.Usage = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotas) Validate() error {
	return dara.Validate(s)
}

type ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage struct {
	// The website ID.
	//
	// example:
	//
	// 34818329392****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// test.top
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The quota usage of the website.
	//
	// example:
	//
	// 1
	SiteUsage *string `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
}

func (s ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) GetSiteName() *string {
	return s.SiteName
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) GetSiteUsage() *string {
	return s.SiteUsage
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) SetSiteId(v int64) *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage {
	s.SiteId = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) SetSiteName(v string) *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage {
	s.SiteName = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) SetSiteUsage(v string) *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage {
	s.SiteUsage = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage) Validate() error {
	return dara.Validate(s)
}
