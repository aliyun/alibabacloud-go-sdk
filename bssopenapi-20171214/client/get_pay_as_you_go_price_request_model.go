// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayAsYouGoPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModuleList(v []*GetPayAsYouGoPriceRequestModuleList) *GetPayAsYouGoPriceRequest
	GetModuleList() []*GetPayAsYouGoPriceRequestModuleList
	SetOwnerId(v int64) *GetPayAsYouGoPriceRequest
	GetOwnerId() *int64
	SetProductCode(v string) *GetPayAsYouGoPriceRequest
	GetProductCode() *string
	SetProductType(v string) *GetPayAsYouGoPriceRequest
	GetProductType() *string
	SetRegion(v string) *GetPayAsYouGoPriceRequest
	GetRegion() *string
	SetSubscriptionType(v string) *GetPayAsYouGoPriceRequest
	GetSubscriptionType() *string
}

type GetPayAsYouGoPriceRequest struct {
	// The details of pricing modules.
	//
	// This parameter is required.
	ModuleList []*GetPayAsYouGoPriceRequestModuleList `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Repeated"`
	OwnerId    *int64                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The billing method. Set the value to PayAsYouGo.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s GetPayAsYouGoPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceRequest) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceRequest) GetModuleList() []*GetPayAsYouGoPriceRequestModuleList {
	return s.ModuleList
}

func (s *GetPayAsYouGoPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPayAsYouGoPriceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetPayAsYouGoPriceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetPayAsYouGoPriceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetPayAsYouGoPriceRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *GetPayAsYouGoPriceRequest) SetModuleList(v []*GetPayAsYouGoPriceRequestModuleList) *GetPayAsYouGoPriceRequest {
	s.ModuleList = v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetOwnerId(v int64) *GetPayAsYouGoPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductCode(v string) *GetPayAsYouGoPriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductType(v string) *GetPayAsYouGoPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetRegion(v string) *GetPayAsYouGoPriceRequest {
	s.Region = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetSubscriptionType(v string) *GetPayAsYouGoPriceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceRequestModuleList struct {
	// The configuration of the Nth pricing module. Valid values of N: 1 to 50. Format: AA:aa,BB:bb. The values of AA and BB are the property IDs of the pricing module. The values of aa and bb are the property values of the pricing module.
	//
	// >  You can call the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html) operation to obtain the configuration parameters of the pricing module.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType:ecs.g5.xlarge,IoOptimized:IoOptimized,ImageOs:linux
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The code of the Nth pricing module.
	//
	// >  You can call the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html) operation to obtain the module code.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The price type of the Nth pricing module. Valid values:
	//
	// 	- Hour: hourly price
	//
	// 	- Usage: usage price
	//
	// 	- Month: monthly price
	//
	// 	- Year: annual price
	//
	// >  You can call the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html) operation to obtain the configuration parameters of the pricing module.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
}

func (s GetPayAsYouGoPriceRequestModuleList) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceRequestModuleList) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceRequestModuleList) GetConfig() *string {
	return s.Config
}

func (s *GetPayAsYouGoPriceRequestModuleList) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetPayAsYouGoPriceRequestModuleList) GetPriceType() *string {
	return s.PriceType
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetConfig(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.Config = &v
	return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetModuleCode(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.ModuleCode = &v
	return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetPriceType(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.PriceType = &v
	return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) Validate() error {
	return dara.Validate(s)
}
