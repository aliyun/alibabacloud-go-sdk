// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *GetResourcePriceRequest
	GetAmount() *int64
	SetAppInstanceType(v string) *GetResourcePriceRequest
	GetAppInstanceType() *string
	SetBizRegionId(v string) *GetResourcePriceRequest
	GetBizRegionId() *string
	SetChargeType(v string) *GetResourcePriceRequest
	GetChargeType() *string
	SetNodeInstanceType(v string) *GetResourcePriceRequest
	GetNodeInstanceType() *string
	SetPeriod(v int64) *GetResourcePriceRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *GetResourcePriceRequest
	GetPeriodUnit() *string
	SetProductType(v string) *GetResourcePriceRequest
	GetProductType() *string
}

type GetResourcePriceRequest struct {
	// The quantity of resources to purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the session specification type to purchase. You can obtain this value by calling the `ListAppInstanceType` operation.
	//
	// Either AppInstanceType or NodeInstanceType must have a value. If both have values, NodeInstanceType is used.
	//
	// example:
	//
	// appstreaming.general
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The region ID of the delivery group. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the resource specification type to purchase. You can obtain this value by calling the [ListNodeInstanceType](https://help.aliyun.com/document_detail/428502.html) operation.
	//
	// Either AppInstanceType or NodeInstanceType must have a value. If both have values, NodeInstanceType is used.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The numeric part of the resource purchase duration. This parameter is used together with PeriodUnit to specify the complete purchase duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit part of the resource purchase duration. This parameter is used together with Period to specify the complete purchase duration. Valid combinations of Period and PeriodUnit:
	//
	// - 1 Week (1 week)
	//
	// - 1 Month (1 month)
	//
	// - 2 Month (2 months)
	//
	// - 3 Month (3 months)
	//
	// - 6 Month (6 months)
	//
	// - 1 Year (1 year)
	//
	// - 2 Year (2 years)
	//
	// - 3 Year (3 years)
	//
	// > This parameter is case-sensitive. For example, `Week` is valid, but `week` is invalid. If the request parameters do not match the combinations listed above, such as `2 Week`, the call to this operation succeeds, but an error occurs during the order placement phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourcePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePriceRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *GetResourcePriceRequest) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *GetResourcePriceRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetResourcePriceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetResourcePriceRequest) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *GetResourcePriceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *GetResourcePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetResourcePriceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetResourcePriceRequest) SetAmount(v int64) *GetResourcePriceRequest {
	s.Amount = &v
	return s
}

func (s *GetResourcePriceRequest) SetAppInstanceType(v string) *GetResourcePriceRequest {
	s.AppInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetBizRegionId(v string) *GetResourcePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetResourcePriceRequest) SetChargeType(v string) *GetResourcePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *GetResourcePriceRequest) SetNodeInstanceType(v string) *GetResourcePriceRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriod(v int64) *GetResourcePriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriodUnit(v string) *GetResourcePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourcePriceRequest) SetProductType(v string) *GetResourcePriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetResourcePriceRequest) Validate() error {
	return dara.Validate(s)
}
