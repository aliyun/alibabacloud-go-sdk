// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRenewPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *GetResourceRenewPriceRequest
	GetAppInstanceGroupId() *string
	SetPeriod(v int64) *GetResourceRenewPriceRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *GetResourceRenewPriceRequest
	GetPeriodUnit() *string
	SetProductType(v string) *GetResourceRenewPriceRequest
	GetProductType() *string
}

type GetResourceRenewPriceRequest struct {
	// The delivery group ID. You can call [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
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

func (s GetResourceRenewPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetResourceRenewPriceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *GetResourceRenewPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetResourceRenewPriceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetResourceRenewPriceRequest) SetAppInstanceGroupId(v string) *GetResourceRenewPriceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriod(v int64) *GetResourceRenewPriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriodUnit(v string) *GetResourceRenewPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetProductType(v string) *GetResourceRenewPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetResourceRenewPriceRequest) Validate() error {
	return dara.Validate(s)
}
