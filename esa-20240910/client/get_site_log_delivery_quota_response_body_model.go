// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteLogDeliveryQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetSiteLogDeliveryQuotaResponseBody
	GetBusinessType() *string
	SetFreeQuota(v int64) *GetSiteLogDeliveryQuotaResponseBody
	GetFreeQuota() *int64
	SetRequestId(v string) *GetSiteLogDeliveryQuotaResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetSiteLogDeliveryQuotaResponseBody
	GetSiteId() *int64
}

type GetSiteLogDeliveryQuotaResponseBody struct {
	// The log category. Valid values:
	//
	// 1.  dcdn_log_access_l1 (default): access logs.
	//
	// 2.  dcdn_log_er: Edge Routine logs.
	//
	// 3.  dcdn_log_waf: firewall logs.
	//
	// 4.  dcdn_log_ipa: TCP/UDP proxy logs.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The remaining quota.
	//
	// example:
	//
	// 3
	FreeQuota *int64 `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C14840EF0EAAB6D97CDE0C5F6554ACE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteLogDeliveryQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaResponseBody) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetSiteLogDeliveryQuotaResponseBody) GetFreeQuota() *int64 {
	return s.FreeQuota
}

func (s *GetSiteLogDeliveryQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteLogDeliveryQuotaResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetBusinessType(v string) *GetSiteLogDeliveryQuotaResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetFreeQuota(v int64) *GetSiteLogDeliveryQuotaResponseBody {
	s.FreeQuota = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetRequestId(v string) *GetSiteLogDeliveryQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) SetSiteId(v int64) *GetSiteLogDeliveryQuotaResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
