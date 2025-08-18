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
	// The number of resources to purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The type ID of the sessions that you purchase. You can call the `ListAppInstanceType` operation to obtain the ID.
	//
	// You must specify one of AppInstanceType and NodeInstanceType. If you specify both of the parameters, the value of NodeInstanceType takes effect.
	//
	// example:
	//
	// appstreaming.general
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the resource type that you purchase. You can call the [ListNodeInstanceType](https://help.aliyun.com/document_detail/428502.html) to obtain the ID.
	//
	// You must specify one of AppInstanceType and NodeInstanceType. If you specify both of the parameters, the value of NodeInstanceType takes effect.
	//
	// Valid values:
	//
	// 	- appstreaming.vgpu.8c16g.4g: WUYING - Graphics - 8 vCPUs, 16 GiB Memory, 4 GiB GPU Memory
	//
	// 	- appstreaming.general.8c16g: WUYING - General - 8 vCPUs, 16 GiB Memory
	//
	// 	- appstreaming.general.4c8g: WUYING - General - 4 vCPUs, 8 GiB Memory
	//
	// 	- appstreaming.vgpu.14c93g.12g: WUYING - Graphics - 14 vCPUs, 93 GiB Memory, 12 GiB GPU Memory.
	//
	// 	- appstreaming.vgpu.8c31g.16g: WUYING - Graphics - 8 vCPUs, 31 GiB Memory, 16 GiB GPU Memory
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The subscription duration of resources. This parameter must be configured together with `PeriodUnit`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter must be configured together with `Period`. The following items describe valid values for the combinations of `Period` and `PeriodUnit`:
	//
	// 	- 1 Week
	//
	// 	- 1 Month
	//
	// 	- 2 Month
	//
	// 	- 3 Month
	//
	// 	- 6 Month
	//
	// 	- 1 Year
	//
	// 	- 2 Year
	//
	// 	- 3 Year
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify a value combination other than the preceding combinations, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
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
