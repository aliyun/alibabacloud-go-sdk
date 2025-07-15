// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorSessionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTrafficMirrorSessionAttributeRequest
	GetDryRun() *bool
	SetEnabled(v bool) *UpdateTrafficMirrorSessionAttributeRequest
	GetEnabled() *bool
	SetOwnerAccount(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTrafficMirrorSessionAttributeRequest
	GetOwnerId() *int64
	SetPacketLength(v int32) *UpdateTrafficMirrorSessionAttributeRequest
	GetPacketLength() *int32
	SetPriority(v int32) *UpdateTrafficMirrorSessionAttributeRequest
	GetPriority() *int32
	SetRegionId(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTrafficMirrorSessionAttributeRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorFilterId(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorFilterId() *string
	SetTrafficMirrorSessionDescription(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorSessionDescription() *string
	SetTrafficMirrorSessionId(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorSessionId() *string
	SetTrafficMirrorSessionName(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorSessionName() *string
	SetTrafficMirrorTargetId(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorTargetId() *string
	SetTrafficMirrorTargetType(v string) *UpdateTrafficMirrorSessionAttributeRequest
	GetTrafficMirrorTargetType() *string
	SetVirtualNetworkId(v int32) *UpdateTrafficMirrorSessionAttributeRequest
	GetVirtualNetworkId() *int32
}

type UpdateTrafficMirrorSessionAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run, without performing the actual request. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable the traffic mirror session. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The maximum transmission unit (MTU).
	//
	// Valid values: **64 to 9600**. Default value: **1500**.
	//
	// example:
	//
	// 1500
	PacketLength *int32 `json:"PacketLength,omitempty" xml:"PacketLength,omitempty"`
	// The new priority of the traffic mirror session. Valid values: **1*	- to **32766**.
	//
	// A smaller value indicates a higher priority. You cannot specify identical priorities for traffic mirror sessions that are created in the same region by using the same account.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the traffic mirror session. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about the regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic mirror filter.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The new description of the traffic mirror session.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a new session.
	TrafficMirrorSessionDescription *string `json:"TrafficMirrorSessionDescription,omitempty" xml:"TrafficMirrorSessionDescription,omitempty"`
	// The ID of the traffic mirror session.
	//
	// This parameter is required.
	//
	// example:
	//
	// tms-j6cla50buc44ap8tu****
	TrafficMirrorSessionId *string `json:"TrafficMirrorSessionId,omitempty" xml:"TrafficMirrorSessionId,omitempty"`
	// The new name of the traffic mirror session.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	TrafficMirrorSessionName *string `json:"TrafficMirrorSessionName,omitempty" xml:"TrafficMirrorSessionName,omitempty"`
	// The ID of the traffic mirror destination.
	//
	// example:
	//
	// eni-j6c2fp57q8rr47rp*****
	TrafficMirrorTargetId *string `json:"TrafficMirrorTargetId,omitempty" xml:"TrafficMirrorTargetId,omitempty"`
	// The new type of the traffic mirror destination. Valid values:
	//
	// 	- **NetworkInterface**: an elastic network interface (ENI)
	//
	// 	- **SLB**: an internal-facing Server Load Balancer (SLB) instance
	//
	// example:
	//
	// NetworkInterface
	TrafficMirrorTargetType *string `json:"TrafficMirrorTargetType,omitempty" xml:"TrafficMirrorTargetType,omitempty"`
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: **0*	- to **16777215**.
	//
	// You can use VNIs to identify mirrored traffic from different sessions at the traffic mirror destination. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
	//
	// example:
	//
	// 10
	VirtualNetworkId *int32 `json:"VirtualNetworkId,omitempty" xml:"VirtualNetworkId,omitempty"`
}

func (s UpdateTrafficMirrorSessionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorSessionAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetPacketLength() *int32 {
	return s.PacketLength
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorSessionDescription() *string {
	return s.TrafficMirrorSessionDescription
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorSessionName() *string {
	return s.TrafficMirrorSessionName
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorTargetId() *string {
	return s.TrafficMirrorTargetId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetTrafficMirrorTargetType() *string {
	return s.TrafficMirrorTargetType
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) GetVirtualNetworkId() *int32 {
	return s.VirtualNetworkId
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetClientToken(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetDryRun(v bool) *UpdateTrafficMirrorSessionAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetEnabled(v bool) *UpdateTrafficMirrorSessionAttributeRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetOwnerAccount(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetOwnerId(v int64) *UpdateTrafficMirrorSessionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetPacketLength(v int32) *UpdateTrafficMirrorSessionAttributeRequest {
	s.PacketLength = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetPriority(v int32) *UpdateTrafficMirrorSessionAttributeRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetRegionId(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetResourceOwnerId(v int64) *UpdateTrafficMirrorSessionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorFilterId(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorSessionDescription(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorSessionDescription = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorSessionId(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorSessionName(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorSessionName = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorTargetId(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorTargetId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetTrafficMirrorTargetType(v string) *UpdateTrafficMirrorSessionAttributeRequest {
	s.TrafficMirrorTargetType = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) SetVirtualNetworkId(v int32) *UpdateTrafficMirrorSessionAttributeRequest {
	s.VirtualNetworkId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
