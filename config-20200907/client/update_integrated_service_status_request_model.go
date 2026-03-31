// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegratedServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorDeliveryDataType(v string) *UpdateIntegratedServiceStatusRequest
	GetAggregatorDeliveryDataType() *string
	SetIntegratedTypes(v string) *UpdateIntegratedServiceStatusRequest
	GetIntegratedTypes() *string
	SetServiceCode(v string) *UpdateIntegratedServiceStatusRequest
	GetServiceCode() *string
	SetStatus(v bool) *UpdateIntegratedServiceStatusRequest
	GetStatus() *bool
}

type UpdateIntegratedServiceStatusRequest struct {
	// The type of the event that is integrated across accounts. Valid values:
	//
	// 	- NonCompliantNotification: non-compliance event
	//
	// example:
	//
	// NonCompliantNotification
	AggregatorDeliveryDataType *string `json:"AggregatorDeliveryDataType,omitempty" xml:"AggregatorDeliveryDataType,omitempty"`
	// The types of the integrated events. Separate multiple event types with commas (,). Valid values:
	//
	// 	- ConfigurationItemChangeNotification: resource change event
	//
	// 	- NonCompliantNotification: non-compliance event
	//
	// example:
	//
	// NonCompliantNotification
	IntegratedTypes *string `json:"IntegratedTypes,omitempty" xml:"IntegratedTypes,omitempty"`
	// The identity of the cloud service that is integrated with Cloud Config. Valid values:
	//
	// 	- eventbridge: EventBridge
	//
	// 	- cms: CloudMonitor
	//
	// 	- bpstudio: Cloud Architect Design Tools
	//
	// This parameter is required.
	//
	// example:
	//
	// cadt
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Specifies whether you want the product to be integrated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIntegratedServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegratedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedServiceStatusRequest) GetAggregatorDeliveryDataType() *string {
	return s.AggregatorDeliveryDataType
}

func (s *UpdateIntegratedServiceStatusRequest) GetIntegratedTypes() *string {
	return s.IntegratedTypes
}

func (s *UpdateIntegratedServiceStatusRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *UpdateIntegratedServiceStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateIntegratedServiceStatusRequest) SetAggregatorDeliveryDataType(v string) *UpdateIntegratedServiceStatusRequest {
	s.AggregatorDeliveryDataType = &v
	return s
}

func (s *UpdateIntegratedServiceStatusRequest) SetIntegratedTypes(v string) *UpdateIntegratedServiceStatusRequest {
	s.IntegratedTypes = &v
	return s
}

func (s *UpdateIntegratedServiceStatusRequest) SetServiceCode(v string) *UpdateIntegratedServiceStatusRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateIntegratedServiceStatusRequest) SetStatus(v bool) *UpdateIntegratedServiceStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateIntegratedServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
