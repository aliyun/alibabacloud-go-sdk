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
	BackupVpcEndpoint *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint `json:"BackupVpcEndpoint,omitempty" xml:"BackupVpcEndpoint,omitempty" type:"Struct"`
	// The time when the network access endpoint was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private egress IP address range of the dedicated network access endpoint. This parameter is returned only when NetworkEndpointType is set to private.
	//
	// example:
	//
	// 172.168.x.x
	EgressPrivateIpAddresses []*string `json:"EgressPrivateIpAddresses,omitempty" xml:"EgressPrivateIpAddresses,omitempty" type:"Repeated"`
	// The public egress IP address range of the shared network access endpoint. This parameter is returned only when NetworkEndpointType is set to shared.
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
	// The ID of the dedicated network access endpoint.
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
	// - shared: Shared network access endpoint.
	//
	// - private: Dedicated network access endpoint.
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
	// - pending: Pending initialization.
	//
	// - creating: Being created.
	//
	// - running: Running.
	//
	// - deleting: Being deleted.
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

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GetBackupVpcEndpoint() *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	return s.BackupVpcEndpoint
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

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetBackupVpcEndpoint(v *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.BackupVpcEndpoint = v
	return s
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
	if s.BackupVpcEndpoint != nil {
		if err := s.BackupVpcEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint struct {
	BackupEgressPrivateIpAddresses []*string `json:"BackupEgressPrivateIpAddresses,omitempty" xml:"BackupEgressPrivateIpAddresses,omitempty" type:"Repeated"`
	BackupEgressPublicIpAddresses  []*string `json:"BackupEgressPublicIpAddresses,omitempty" xml:"BackupEgressPublicIpAddresses,omitempty" type:"Repeated"`
	BackupSecurityGroupId          *string   `json:"BackupSecurityGroupId,omitempty" xml:"BackupSecurityGroupId,omitempty"`
	BackupVSwitchIds               []*string `json:"BackupVSwitchIds,omitempty" xml:"BackupVSwitchIds,omitempty" type:"Repeated"`
	BackupVpcId                    *string   `json:"BackupVpcId,omitempty" xml:"BackupVpcId,omitempty"`
	BackupVpcRegionId              *string   `json:"BackupVpcRegionId,omitempty" xml:"BackupVpcRegionId,omitempty"`
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupEgressPrivateIpAddresses() []*string {
	return s.BackupEgressPrivateIpAddresses
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupEgressPublicIpAddresses() []*string {
	return s.BackupEgressPublicIpAddresses
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupSecurityGroupId() *string {
	return s.BackupSecurityGroupId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupVSwitchIds() []*string {
	return s.BackupVSwitchIds
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupVpcId() *string {
	return s.BackupVpcId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) GetBackupVpcRegionId() *string {
	return s.BackupVpcRegionId
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupEgressPrivateIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupEgressPrivateIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupEgressPublicIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupEgressPublicIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupSecurityGroupId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupSecurityGroupId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupVSwitchIds(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupVSwitchIds = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupVpcId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupVpcId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) SetBackupVpcRegionId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint {
	s.BackupVpcRegionId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpointBackupVpcEndpoint) Validate() error {
	return dara.Validate(s)
}
