// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceSubscriptionEstimateCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetClientToken() *string
	SetOrderType(v string) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetOrderType() *string
	SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetRegionId() *string
	SetResourcePeriod(v []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetResourcePeriod() []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod
	SetServiceInstanceId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest
	GetServiceInstanceId() *string
}

type GetServiceInstanceSubscriptionEstimateCostRequest struct {
	// Ensures idempotence of the request. Generate a parameter value from your client to ensure its uniqueness across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Order type. Possible value: Renewal.
	//
	// This parameter is required.
	//
	// example:
	//
	// Renewal
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The renewal duration for all prepaid resources of the service instance. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the renewal duration of all prepaid resources of the service instance, which is the unit of the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource renewal configuration.
	ResourcePeriod []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod `json:"ResourcePeriod,omitempty" xml:"ResourcePeriod,omitempty" type:"Repeated"`
	// Service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetResourcePeriod() []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	return s.ResourcePeriod
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetClientToken(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetOrderType(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.OrderType = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetRegionId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetResourcePeriod(v []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ResourcePeriod = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod struct {
	// Renewal duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the resource renewal duration, which is the unit of the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Resource ARN (Aliyun Resource Name).
	//
	// example:
	//
	// acs:ecs:cn-guangzhou:1361753504587228:instance/i-7xv9pgeqvhxg10jji3vd
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) GetPeriod() *int32 {
	return s.Period
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetResourceArn(v string) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.ResourceArn = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) Validate() error {
	return dara.Validate(s)
}
