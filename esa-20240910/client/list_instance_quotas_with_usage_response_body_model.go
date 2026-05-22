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
	InstanceId *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Quotas     []*ListInstanceQuotasWithUsageResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
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
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceQuotasWithUsageResponseBodyQuotas struct {
	QuotaName  *string                                                   `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	QuotaValue *string                                                   `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	SiteUsage  []*ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty" type:"Repeated"`
	Usage      *string                                                   `json:"Usage,omitempty" xml:"Usage,omitempty"`
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
	if s.SiteUsage != nil {
		for _, item := range s.SiteUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceQuotasWithUsageResponseBodyQuotasSiteUsage struct {
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
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
