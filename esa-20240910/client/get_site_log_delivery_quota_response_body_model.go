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
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	FreeQuota    *int64  `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId       *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
