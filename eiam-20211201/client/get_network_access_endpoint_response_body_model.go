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
	// The network access endpoint information.
	NetworkAccessEndpoint *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint `json:"NetworkAccessEndpoint,omitempty" xml:"NetworkAccessEndpoint,omitempty" type:"Struct"`
	// The request ID.
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
	if s.NetworkAccessEndpoint != nil {
		if err := s.NetworkAccessEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint struct {
	// The time when the network access endpoint was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private egress IP addresses of the dedicated network access endpoint. This parameter is returned only when NetworkEndpointType is set to private.
	//
	// example:
	//
	// 172.168.x.x
	EgressPrivateIpAddresses []*string `json:"EgressPrivateIpAddresses,omitempty" xml:"EgressPrivateIpAddresses,omitempty" type:"Repeated"`
	// The public egress IP addresses of the shared network access endpoint. This parameter is returned only when NetworkEndpointType is set to shared.
	//
	// example:
	//
	// 203.0.XX.XX/27
	EgressPublicIpAddresses []*string `json:"EgressPublicIpAddresses,omitempty" xml:"EgressPublicIpAddresses,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The dedicated network access endpoint ID.
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The name of the dedicated network access endpoint.
	//
	// example:
	//
	// Xx-business VPC access endpoint
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// The type of the network access endpoint. Valid values:
	//
	// - shared: shared network access endpoint.
	//
	// - private: dedicated network access endpoint.
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// The ID of the security group used by the dedicated network access endpoint.
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the network access endpoint. Valid values:
	//
	// - pending: pending initialization.
	//
	// - creating: being created.
	//
	// - running: running.
	//
	// - deleting: being deleted.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the dedicated network access endpoint was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The list of vSwitches to which the dedicated network access endpoint is connected.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the dedicated network access endpoint is connected.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region of the VPC to which the dedicated network access endpoint is connected.
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
