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
	// The billing method. Valid values: payasyougo (pay-as-you-go) and subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// payasyougo
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The commodity code. Example: On Alibaba Cloud China Website (www.aliyun.com), the commodity code for pay-as-you-go is odps and the commodity code for subscription is odpsplus. On Alibaba Cloud International Website (www.alibabacloud.com), the commodity code for pay-as-you-go is odps_intl and the commodity code for subscription is odpsplus_intl.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The commodity specifications of the quota.
	//
	// 	Notice:
	//
	// - This parameter is required only for subscription quotas.
	//
	// The minimum unit for cu is 50.
	//
	// ord_time supports monthly (month) and yearly (year) billing.
	//
	// - After creation, commodityData cannot be modified. To modify it, go to the MaxCompute console.
	//
	// example:
	//
	// {"CU":50,"ord_time":"1:Month","autoRenew":false}
	CommodityData *string `json:"commodityData,omitempty" xml:"commodityData,omitempty"`
	// 	Notice: This parameter is required only for subscription quotas.
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
