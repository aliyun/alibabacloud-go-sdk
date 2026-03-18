// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateQuotaRequest
	GetChargeType() *string
	SetCommodityCode(v string) *CreateQuotaRequest
	GetCommodityCode() *string
	SetCommodityData(v string) *CreateQuotaRequest
	GetCommodityData() *string
	SetPartNickName(v string) *CreateQuotaRequest
	GetPartNickName() *string
}

type CreateQuotaRequest struct {
	// Billing type. Use payasyougo for pay-as-you-go or subscription for subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// payasyougo
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// Product code. For example:
	//
	// Alibaba Cloud China Website (www\\.aliyun.com): use odps for pay-as-you-go and odpsplus for subscription.
	//
	// Alibaba Cloud International Website (www\\.alibabacloud.com): use odps_intl for pay-as-you-go and odpsplus_intl for subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// Quota specification.
	//
	// 	Notice: Required only for subscription quotas.
	//
	// The minimum CU value is 50.
	//
	// ord_time supports month and year.
	//
	// example:
	//
	// {"CU":50,"ord_time":"1:Month","autoRenew":false}
	CommodityData *string `json:"commodityData,omitempty" xml:"commodityData,omitempty"`
	// 	Notice:
	//
	// Required only for subscription quotas.
	//
	// example:
	//
	// MyQuota
	PartNickName *string `json:"partNickName,omitempty" xml:"partNickName,omitempty"`
}

func (s CreateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateQuotaRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateQuotaRequest) GetCommodityData() *string {
	return s.CommodityData
}

func (s *CreateQuotaRequest) GetPartNickName() *string {
	return s.PartNickName
}

func (s *CreateQuotaRequest) SetChargeType(v string) *CreateQuotaRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateQuotaRequest) SetCommodityCode(v string) *CreateQuotaRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateQuotaRequest) SetCommodityData(v string) *CreateQuotaRequest {
	s.CommodityData = &v
	return s
}

func (s *CreateQuotaRequest) SetPartNickName(v string) *CreateQuotaRequest {
	s.PartNickName = &v
	return s
}

func (s *CreateQuotaRequest) Validate() error {
	return dara.Validate(s)
}
