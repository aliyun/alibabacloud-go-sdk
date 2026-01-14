// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointTrafficPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest
	GetClientToken() *string
	SetEndpointId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest
	GetEndpointId() *string
	SetPolicyConfigurations(v []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *UpdateCustomRoutingEndpointTrafficPoliciesRequest
	GetPolicyConfigurations() []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations
	SetRegionId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest
	GetRegionId() *string
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint for which you want to modify the traffic policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-2zewuzypq5e6r3pfh****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The access policies.
	//
	// You can specify a maximum of 500 traffic policies for each endpoint.
	//
	// This parameter is required.
	PolicyConfigurations []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) GetPolicyConfigurations() []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyConfigurations(v []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) Validate() error {
	if s.PolicyConfigurations != nil {
		for _, item := range s.PolicyConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations struct {
	// The IP address of the destination which to allow traffic.
	//
	// This parameter takes effect only when you set the **TrafficToEndpointPolicy*	- parameter to **AllowCustom**. You can call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation to query the traffic policy of an endpoint.
	//
	// You can specify a maximum of 500 traffic policies for each endpoint.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the traffic policy that you want to modify.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// ply-bptest2****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The port range of the destination to which traffic is forwarded. The value of this parameter must fall within the port range of the endpoint group.
	//
	// If you do not specify object, traffic is forwarded to all ports.
	//
	// This parameter takes effect only when you set the **TrafficToEndpointPolicy*	- parameter to **AllowCustom**. You can call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation to query the traffic policy of an endpoint.
	//
	// You can specify a maximum of 500 port ranges for each endpoint and a maximum of 10 port ranges for each traffic policy.
	PortRanges []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GetPortRanges() []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetAddress(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPolicyId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPortRanges(v []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges struct {
	// The first port of the destination port range. The value of this parameter must fall within the port range of the backend service.
	//
	// This parameter takes effect only when you set the **TrafficToEndpointPolicy*	- parameter to **AllowCustom**. You can call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation to query the traffic policy of an endpoint.
	//
	// If the start port and end port values are empty, traffic on all ports of the destination is allowed.
	//
	// You can specify a maximum of 500 port ranges for each endpoint and a maximum of 10 port ranges for each traffic policy.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the destination port range. The value of this parameter must fall within the port range of the backend service.
	//
	// This parameter takes effect only when you set the **TrafficToEndpointPolicy*	- parameter to **AllowCustom**. You can call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation to query the traffic policy of an endpoint.
	//
	// If the start port and end port values are empty, traffic on all ports of the destination is allowed.
	//
	// You can specify a maximum of 500 port ranges for each endpoint and a maximum of 10 port ranges for each traffic policy.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetFromPort(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetToPort(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}
