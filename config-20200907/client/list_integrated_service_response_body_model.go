// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegratedServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListIntegratedServiceResponseBodyData) *ListIntegratedServiceResponseBody
	GetData() []*ListIntegratedServiceResponseBodyData
	SetRequestId(v string) *ListIntegratedServiceResponseBody
	GetRequestId() *string
}

type ListIntegratedServiceResponseBody struct {
	// The information about the cloud service that can be integrated.
	Data []*ListIntegratedServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 86DB52A5-0C25-505A-96D5-9BAE1EFA00B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntegratedServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegratedServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegratedServiceResponseBody) GetData() []*ListIntegratedServiceResponseBodyData {
	return s.Data
}

func (s *ListIntegratedServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegratedServiceResponseBody) SetData(v []*ListIntegratedServiceResponseBodyData) *ListIntegratedServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListIntegratedServiceResponseBody) SetRequestId(v string) *ListIntegratedServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegratedServiceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegratedServiceResponseBodyData struct {
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
	// The identifier of the cloud service. Valid values:
	//
	// 	- eventbridge: EventBridge
	//
	// 	- cms: CloudMonitor
	//
	// 	- bpstudio: Cloud Architect Design Tools
	//
	// example:
	//
	// cms
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The name of the cloud service.
	//
	// example:
	//
	// cms
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The integration status of the cloud service. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIntegratedServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntegratedServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntegratedServiceResponseBodyData) GetAggregatorDeliveryDataType() *string {
	return s.AggregatorDeliveryDataType
}

func (s *ListIntegratedServiceResponseBodyData) GetIntegratedTypes() *string {
	return s.IntegratedTypes
}

func (s *ListIntegratedServiceResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListIntegratedServiceResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListIntegratedServiceResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *ListIntegratedServiceResponseBodyData) SetAggregatorDeliveryDataType(v string) *ListIntegratedServiceResponseBodyData {
	s.AggregatorDeliveryDataType = &v
	return s
}

func (s *ListIntegratedServiceResponseBodyData) SetIntegratedTypes(v string) *ListIntegratedServiceResponseBodyData {
	s.IntegratedTypes = &v
	return s
}

func (s *ListIntegratedServiceResponseBodyData) SetServiceCode(v string) *ListIntegratedServiceResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *ListIntegratedServiceResponseBodyData) SetServiceName(v string) *ListIntegratedServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListIntegratedServiceResponseBodyData) SetStatus(v bool) *ListIntegratedServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListIntegratedServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
