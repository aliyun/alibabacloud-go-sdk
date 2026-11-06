// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *GetResourceQuotaRequest
	GetBizRegionId() *string
	SetChargeType(v string) *GetResourceQuotaRequest
	GetChargeType() *string
	SetProductType(v string) *GetResourceQuotaRequest
	GetProductType() *string
	SetResourceType(v string) *GetResourceQuotaRequest
	GetResourceType() *string
}

type GetResourceQuotaRequest struct {
	// The business region ID. If this parameter is omitted, the general quota that is not limited by region is queried. The quotas of different regions are not added together. This parameter must be consistent with the business region for What to do next.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method query condition. If this parameter is omitted, the general quota that is not limited by billing method is queried. The quotas of different billing methods are not added together. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The product type. Set this parameter to the following value when you call the operation:
	//
	// - CloudApp: WUYING Cloud App.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The business resource type to query. Valid values:
	//
	// - AppInstanceGroup: the quantity quota for delivery groups. Unit: count.
	//
	// - AppAuthUser: the quantity quota for authorized users. Unit: users.
	//
	// - PoolNodeAmount: the quantity quota for purchased resource nodes. Unit: count.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetResourceQuotaRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetResourceQuotaRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetResourceQuotaRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetResourceQuotaRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceQuotaRequest) SetBizRegionId(v string) *GetResourceQuotaRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetResourceQuotaRequest) SetChargeType(v string) *GetResourceQuotaRequest {
	s.ChargeType = &v
	return s
}

func (s *GetResourceQuotaRequest) SetProductType(v string) *GetResourceQuotaRequest {
	s.ProductType = &v
	return s
}

func (s *GetResourceQuotaRequest) SetResourceType(v string) *GetResourceQuotaRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
