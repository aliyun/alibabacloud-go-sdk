// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessEndpoints(v []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) *ListNetworkAccessEndpointsResponseBody
	GetNetworkAccessEndpoints() []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints
	SetNextToken(v string) *ListNetworkAccessEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNetworkAccessEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListNetworkAccessEndpointsResponseBody
	GetTotalCount() *int64
}

type ListNetworkAccessEndpointsResponseBody struct {
	// The list of network access endpoints.
	NetworkAccessEndpoints []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints `json:"NetworkAccessEndpoints,omitempty" xml:"NetworkAccessEndpoints,omitempty" type:"Repeated"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBody) GetNetworkAccessEndpoints() []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	return s.NetworkAccessEndpoints
}

func (s *ListNetworkAccessEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkAccessEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkAccessEndpointsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNetworkAccessEndpoints(v []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) *ListNetworkAccessEndpointsResponseBody {
	s.NetworkAccessEndpoints = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNextToken(v string) *ListNetworkAccessEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetTotalCount(v int64) *ListNetworkAccessEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) Validate() error {
	if s.NetworkAccessEndpoints != nil {
		for _, item := range s.NetworkAccessEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints struct {
	BackupVpcEndpoint *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint `json:"BackupVpcEndpoint,omitempty" xml:"BackupVpcEndpoint,omitempty" type:"Struct"`
	// The creation time of the network access endpoint. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network access endpoint ID.
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The network access endpoint name.
	//
	// example:
	//
	// VPC access endpoint for xx service
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
	// The security group ID used by the dedicated network access endpoint.
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the network access endpoint. Valid values:
	//
	//
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
	// The last update time of the network access endpoint. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The list of vSwitches for the dedicated network access endpoint.
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID of the dedicated network access endpoint.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region of the VPC for the dedicated network access endpoint.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetBackupVpcEndpoint() *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	return s.BackupVpcEndpoint
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetNetworkAccessEndpointType() *string {
	return s.NetworkAccessEndpointType
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetStatus() *string {
	return s.Status
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetBackupVpcEndpoint(v *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.BackupVpcEndpoint = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetCreateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetInstanceId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointName(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetSecurityGroupId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.SecurityGroupId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetStatus(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.Status = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetUpdateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.UpdateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVSwitchIds(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VSwitchIds = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcRegionId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcRegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) Validate() error {
	if s.BackupVpcEndpoint != nil {
		if err := s.BackupVpcEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint struct {
	BackupEgressPrivateIpAddresses []*string `json:"BackupEgressPrivateIpAddresses,omitempty" xml:"BackupEgressPrivateIpAddresses,omitempty" type:"Repeated"`
	BackupEgressPublicIpAddresses  []*string `json:"BackupEgressPublicIpAddresses,omitempty" xml:"BackupEgressPublicIpAddresses,omitempty" type:"Repeated"`
	BackupSecurityGroupId          *string   `json:"BackupSecurityGroupId,omitempty" xml:"BackupSecurityGroupId,omitempty"`
	BackupVSwitchIds               []*string `json:"BackupVSwitchIds,omitempty" xml:"BackupVSwitchIds,omitempty" type:"Repeated"`
	BackupVpcId                    *string   `json:"BackupVpcId,omitempty" xml:"BackupVpcId,omitempty"`
	BackupVpcRegionId              *string   `json:"BackupVpcRegionId,omitempty" xml:"BackupVpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupEgressPrivateIpAddresses() []*string {
	return s.BackupEgressPrivateIpAddresses
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupEgressPublicIpAddresses() []*string {
	return s.BackupEgressPublicIpAddresses
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupSecurityGroupId() *string {
	return s.BackupSecurityGroupId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupVSwitchIds() []*string {
	return s.BackupVSwitchIds
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupVpcId() *string {
	return s.BackupVpcId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) GetBackupVpcRegionId() *string {
	return s.BackupVpcRegionId
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupEgressPrivateIpAddresses(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupEgressPrivateIpAddresses = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupEgressPublicIpAddresses(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupEgressPublicIpAddresses = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupSecurityGroupId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupSecurityGroupId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupVSwitchIds(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupVSwitchIds = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupVpcId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupVpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) SetBackupVpcRegionId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint {
	s.BackupVpcRegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpointsBackupVpcEndpoint) Validate() error {
	return dara.Validate(s)
}
