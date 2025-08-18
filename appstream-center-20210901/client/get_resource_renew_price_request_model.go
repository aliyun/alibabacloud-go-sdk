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
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
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
