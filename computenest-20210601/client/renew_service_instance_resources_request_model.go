// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RenewServiceInstanceResourcesRequest
	GetClientToken() *string
	SetPeriod(v int32) *RenewServiceInstanceResourcesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewServiceInstanceResourcesRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *RenewServiceInstanceResourcesRequest
	GetRegionId() *string
	SetResourcePeriod(v []*RenewServiceInstanceResourcesRequestResourcePeriod) *RenewServiceInstanceResourcesRequest
	GetResourcePeriod() []*RenewServiceInstanceResourcesRequestResourcePeriod
	SetServiceInstanceId(v string) *RenewServiceInstanceResourcesRequest
	GetServiceInstanceId() *string
}

type RenewServiceInstanceResourcesRequest struct {
	// A client token to ensure that the request is idempotent. Generate a unique value for this parameter from your client for each request. The token can contain only ASCII characters and must not exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The renewal duration for all subscription resources in the service instance. The unit is specified by \\`PeriodUnit\\`.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit of the renewal duration for all subscription resources in the service instance. This parameter specifies the unit for \\`Period\\`. Valid values: \\`Month\\` and \\`Year\\`. Default value: \\`Month\\`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource renewals.
	ResourcePeriod []*RenewServiceInstanceResourcesRequestResourcePeriod `json:"ResourcePeriod,omitempty" xml:"ResourcePeriod,omitempty" type:"Repeated"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RenewServiceInstanceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewServiceInstanceResourcesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewServiceInstanceResourcesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewServiceInstanceResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewServiceInstanceResourcesRequest) GetResourcePeriod() []*RenewServiceInstanceResourcesRequestResourcePeriod {
	return s.ResourcePeriod
}

func (s *RenewServiceInstanceResourcesRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *RenewServiceInstanceResourcesRequest) SetClientToken(v string) *RenewServiceInstanceResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetPeriod(v int32) *RenewServiceInstanceResourcesRequest {
	s.Period = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetPeriodUnit(v string) *RenewServiceInstanceResourcesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetRegionId(v string) *RenewServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetResourcePeriod(v []*RenewServiceInstanceResourcesRequestResourcePeriod) *RenewServiceInstanceResourcesRequest {
	s.ResourcePeriod = v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *RenewServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) Validate() error {
	if s.ResourcePeriod != nil {
		for _, item := range s.ResourcePeriod {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RenewServiceInstanceResourcesRequestResourcePeriod struct {
	// The renewal duration for the resource. The unit is specified by \\`PeriodUnit\\`.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit of the renewal duration for the resource. This parameter specifies the unit for \\`Period\\`. Valid values: \\`Month\\` and \\`Year\\`. Default value: \\`Month\\`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The Aliyun Resource Name (ARN) of the resource.
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s RenewServiceInstanceResourcesRequestResourcePeriod) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResourcesRequestResourcePeriod) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetPeriod(v int32) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.Period = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetPeriodUnit(v string) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetResourceArn(v string) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.ResourceArn = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) Validate() error {
	return dara.Validate(s)
}
