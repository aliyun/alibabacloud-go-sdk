// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAccessEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessEndpoint(v *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) *GetNetworkAccessEndpointResponseBody
	GetNetworkAccessEndpoint() *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint
	SetRequestId(v string) *GetNetworkAccessEndpointResponseBody
	GetRequestId() *string
}

type GetNetworkAccessEndpointResponseBody struct {
	// Network endpoint information.
	NetworkAccessEndpoint *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint `json:"NetworkAccessEndpoint,omitempty" xml:"NetworkAccessEndpoint,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkAccessEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponseBody) GetNetworkAccessEndpoint() *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	return s.NetworkAccessEndpoint
}

func (s *GetNetworkAccessEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkAccessEndpointResponseBody) SetNetworkAccessEndpoint(v *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) *GetNetworkAccessEndpointResponseBody {
	s.NetworkAccessEndpoint = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBody) SetRequestId(v string) *GetNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint struct {
	// The time when the baseline was created.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Public egress ip address range of the dedicated network endpoint
	//
	// This field is returned only when NetworkEndpointType is set to private.
	//
	// example:
	//
	// 172.168.x.x
	EgressPrivateIpAddresses []*string `json:"EgressPrivateIpAddresses,omitempty" xml:"EgressPrivateIpAddresses,omitempty" type:"Repeated"`
	// Public egress ip address range of the shared network endpoint
	//
	// This field is returned only when networkEndpointType is set to shared.
	//
	// example:
	//
	// 8.xx.xx.xxx/27
	EgressPublicIpAddresses []*string `json:"EgressPublicIpAddresses,omitempty" xml:"EgressPublicIpAddresses,omitempty" type:"Repeated"`
	// Instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique identifier of the network access endpoint.
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// Private network endpoint name.
	//
	// example:
	//
	// xx business VPC access endpoint
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// Type of the Network Endpoint
	//
	// Possible values:
	//
	// shared: Shared network endpoint
	//
	// private: Dedicated network endpoint
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// The ID of the destination security group.
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Status of the Network Endpoint
	//
	// Possible values:
	//
	// pending: Pending initialization
	//
	// creating: Being created
	//
	// running: Running
	//
	// deleting: Being deleted
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the endpoint was updated.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// List of specified vSwitches associated with the dedicated network endpoint connection.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the outbound virtual private cloud (VPC).
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetEgressPrivateIpAddresses() []*string {
	return s.EgressPrivateIpAddresses
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetEgressPublicIpAddresses() []*string {
	return s.EgressPublicIpAddresses
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetNetworkAccessEndpointType() *string {
	return s.NetworkAccessEndpointType
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetStatus() *string {
	return s.Status
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetCreateTime(v int64) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetEgressPrivateIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.EgressPrivateIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetEgressPublicIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.EgressPublicIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetInstanceId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointName(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointType(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetSecurityGroupId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.SecurityGroupId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetStatus(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.Status = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetUpdateTime(v int64) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.UpdateTime = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVSwitchIds(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VSwitchIds = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVpcId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VpcId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVpcRegionId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VpcRegionId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) Validate() error {
	return dara.Validate(s)
}
