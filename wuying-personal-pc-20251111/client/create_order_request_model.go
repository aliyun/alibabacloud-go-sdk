// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateOrderRequest
	GetApiKey() *string
	SetAutoPay(v bool) *CreateOrderRequest
	GetAutoPay() *bool
	SetDynamicProductParams(v []*CreateOrderRequestDynamicProductParams) *CreateOrderRequest
	GetDynamicProductParams() []*CreateOrderRequestDynamicProductParams
	SetOperationType(v string) *CreateOrderRequest
	GetOperationType() *string
	SetOrderFrom(v string) *CreateOrderRequest
	GetOrderFrom() *string
	SetSessionId(v string) *CreateOrderRequest
	GetSessionId() *string
}

type CreateOrderRequest struct {
	// This parameter is required.
	ApiKey  *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	AutoPay *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	DynamicProductParams []*CreateOrderRequestDynamicProductParams `json:"DynamicProductParams,omitempty" xml:"DynamicProductParams,omitempty" type:"Repeated"`
	// This parameter is required.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OrderFrom     *string `json:"OrderFrom,omitempty" xml:"OrderFrom,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateOrderRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderRequest) GetDynamicProductParams() []*CreateOrderRequestDynamicProductParams {
	return s.DynamicProductParams
}

func (s *CreateOrderRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateOrderRequest) GetOrderFrom() *string {
	return s.OrderFrom
}

func (s *CreateOrderRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateOrderRequest) SetApiKey(v string) *CreateOrderRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateOrderRequest) SetAutoPay(v bool) *CreateOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderRequest) SetDynamicProductParams(v []*CreateOrderRequestDynamicProductParams) *CreateOrderRequest {
	s.DynamicProductParams = v
	return s
}

func (s *CreateOrderRequest) SetOperationType(v string) *CreateOrderRequest {
	s.OperationType = &v
	return s
}

func (s *CreateOrderRequest) SetOrderFrom(v string) *CreateOrderRequest {
	s.OrderFrom = &v
	return s
}

func (s *CreateOrderRequest) SetSessionId(v string) *CreateOrderRequest {
	s.SessionId = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	if s.DynamicProductParams != nil {
		for _, item := range s.DynamicProductParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrderRequestDynamicProductParams struct {
	Amount              *int32                                                     `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DeliveryAddress     *string                                                    `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DynamicAttributes   map[string]*string                                         `json:"DynamicAttributes,omitempty" xml:"DynamicAttributes,omitempty"`
	InputActivateConfig *CreateOrderRequestDynamicProductParamsInputActivateConfig `json:"InputActivateConfig,omitempty" xml:"InputActivateConfig,omitempty" type:"Struct"`
	InstanceIdList      []*string                                                  `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	LinkedResourceId    *string                                                    `json:"LinkedResourceId,omitempty" xml:"LinkedResourceId,omitempty"`
	PackageCode         *string                                                    `json:"PackageCode,omitempty" xml:"PackageCode,omitempty"`
	ProductCode         *string                                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductSkuCode      *string                                                    `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ResourceId          *string                                                    `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CreateOrderRequestDynamicProductParams) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestDynamicProductParams) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestDynamicProductParams) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateOrderRequestDynamicProductParams) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *CreateOrderRequestDynamicProductParams) GetDynamicAttributes() map[string]*string {
	return s.DynamicAttributes
}

func (s *CreateOrderRequestDynamicProductParams) GetInputActivateConfig() *CreateOrderRequestDynamicProductParamsInputActivateConfig {
	return s.InputActivateConfig
}

func (s *CreateOrderRequestDynamicProductParams) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *CreateOrderRequestDynamicProductParams) GetLinkedResourceId() *string {
	return s.LinkedResourceId
}

func (s *CreateOrderRequestDynamicProductParams) GetPackageCode() *string {
	return s.PackageCode
}

func (s *CreateOrderRequestDynamicProductParams) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateOrderRequestDynamicProductParams) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *CreateOrderRequestDynamicProductParams) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateOrderRequestDynamicProductParams) SetAmount(v int32) *CreateOrderRequestDynamicProductParams {
	s.Amount = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetDeliveryAddress(v string) *CreateOrderRequestDynamicProductParams {
	s.DeliveryAddress = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetDynamicAttributes(v map[string]*string) *CreateOrderRequestDynamicProductParams {
	s.DynamicAttributes = v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetInputActivateConfig(v *CreateOrderRequestDynamicProductParamsInputActivateConfig) *CreateOrderRequestDynamicProductParams {
	s.InputActivateConfig = v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetInstanceIdList(v []*string) *CreateOrderRequestDynamicProductParams {
	s.InstanceIdList = v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetLinkedResourceId(v string) *CreateOrderRequestDynamicProductParams {
	s.LinkedResourceId = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetPackageCode(v string) *CreateOrderRequestDynamicProductParams {
	s.PackageCode = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetProductCode(v string) *CreateOrderRequestDynamicProductParams {
	s.ProductCode = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetProductSkuCode(v string) *CreateOrderRequestDynamicProductParams {
	s.ProductSkuCode = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) SetResourceId(v string) *CreateOrderRequestDynamicProductParams {
	s.ResourceId = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParams) Validate() error {
	if s.InputActivateConfig != nil {
		if err := s.InputActivateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderRequestDynamicProductParamsInputActivateConfig struct {
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s CreateOrderRequestDynamicProductParamsInputActivateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestDynamicProductParamsInputActivateConfig) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) GetCityName() *string {
	return s.CityName
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) GetDesktopName() *string {
	return s.DesktopName
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) GetImageId() *string {
	return s.ImageId
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) SetCityName(v string) *CreateOrderRequestDynamicProductParamsInputActivateConfig {
	s.CityName = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) SetDesktopName(v string) *CreateOrderRequestDynamicProductParamsInputActivateConfig {
	s.DesktopName = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) SetImageId(v string) *CreateOrderRequestDynamicProductParamsInputActivateConfig {
	s.ImageId = &v
	return s
}

func (s *CreateOrderRequestDynamicProductParamsInputActivateConfig) Validate() error {
	return dara.Validate(s)
}
