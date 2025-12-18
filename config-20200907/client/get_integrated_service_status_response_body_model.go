// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegratedServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorDeliveryDataType(v string) *GetIntegratedServiceStatusResponseBody
	GetAggregatorDeliveryDataType() *string
	SetData(v bool) *GetIntegratedServiceStatusResponseBody
	GetData() *bool
	SetIntegratedTypes(v string) *GetIntegratedServiceStatusResponseBody
	GetIntegratedTypes() *string
	SetRequestId(v string) *GetIntegratedServiceStatusResponseBody
	GetRequestId() *string
}

type GetIntegratedServiceStatusResponseBody struct {
	// The type of the event that is integrated across accounts. Valid values:
	//
	// 	- NonCompliantNotification: non-compliance event
	//
	// example:
	//
	// NonCompliantNotification
	AggregatorDeliveryDataType *string `json:"AggregatorDeliveryDataType,omitempty" xml:"AggregatorDeliveryDataType,omitempty"`
	// Indicates whether the product has been integrated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 2E396C84-8D50-5F95-97FA-C0367181BA8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIntegratedServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegratedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegratedServiceStatusResponseBody) GetAggregatorDeliveryDataType() *string {
	return s.AggregatorDeliveryDataType
}

func (s *GetIntegratedServiceStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *GetIntegratedServiceStatusResponseBody) GetIntegratedTypes() *string {
	return s.IntegratedTypes
}

func (s *GetIntegratedServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegratedServiceStatusResponseBody) SetAggregatorDeliveryDataType(v string) *GetIntegratedServiceStatusResponseBody {
	s.AggregatorDeliveryDataType = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetData(v bool) *GetIntegratedServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetIntegratedTypes(v string) *GetIntegratedServiceStatusResponseBody {
	s.IntegratedTypes = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) SetRequestId(v string) *GetIntegratedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegratedServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
