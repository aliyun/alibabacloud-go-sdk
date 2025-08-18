// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteLogDeliveryQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetSiteLogDeliveryQuotaRequest
	GetBusinessType() *string
	SetSiteId(v int64) *GetSiteLogDeliveryQuotaRequest
	GetSiteId() *int64
}

type GetSiteLogDeliveryQuotaRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteLogDeliveryQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetSiteLogDeliveryQuotaRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteLogDeliveryQuotaRequest) SetBusinessType(v string) *GetSiteLogDeliveryQuotaRequest {
	s.BusinessType = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaRequest) SetSiteId(v int64) *GetSiteLogDeliveryQuotaRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaRequest) Validate() error {
	return dara.Validate(s)
}
