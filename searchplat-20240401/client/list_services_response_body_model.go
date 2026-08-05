// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListServicesResponseBody
	GetRequestId() *string
	SetResult(v []*ListServicesResponseBodyResult) *ListServicesResponseBody
	GetResult() []*ListServicesResponseBodyResult
}

type ListServicesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2423C841-91C4-5E51-B296-590D367967FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task execution result.
	Result []*ListServicesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServicesResponseBody) GetResult() []*ListServicesResponseBodyResult {
	return s.Result
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetResult(v []*ListServicesResponseBodyResult) *ListServicesResponseBody {
	s.Result = v
	return s
}

func (s *ListServicesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServicesResponseBodyResult struct {
	// The billing method.
	ChargeWay []*string `json:"chargeWay,omitempty" xml:"chargeWay,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// 33
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The metadata.
	//
	// example:
	//
	// {
	//
	//             "maxTokens": 512
	//
	//         }
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// The model type.
	//
	// example:
	//
	// deployment
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The service name.
	//
	// example:
	//
	// 文本向量化
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ops-text-embedding-001
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service type.
	//
	// example:
	//
	// text-embedding
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListServicesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyResult) GetChargeWay() []*string {
	return s.ChargeWay
}

func (s *ListServicesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListServicesResponseBodyResult) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListServicesResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *ListServicesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListServicesResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServicesResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesResponseBodyResult) SetChargeWay(v []*string) *ListServicesResponseBodyResult {
	s.ChargeWay = v
	return s
}

func (s *ListServicesResponseBodyResult) SetDescription(v string) *ListServicesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListServicesResponseBodyResult) SetMeta(v map[string]interface{}) *ListServicesResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListServicesResponseBodyResult) SetModelType(v string) *ListServicesResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *ListServicesResponseBodyResult) SetName(v string) *ListServicesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyResult) SetServiceId(v string) *ListServicesResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyResult) SetServiceType(v string) *ListServicesResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
