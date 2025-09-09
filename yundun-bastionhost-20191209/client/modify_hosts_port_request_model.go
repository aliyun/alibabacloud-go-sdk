// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsPortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostIds(v string) *ModifyHostsPortRequest
	GetHostIds() *string
	SetInstanceId(v string) *ModifyHostsPortRequest
	GetInstanceId() *string
	SetPort(v string) *ModifyHostsPortRequest
	GetPort() *string
	SetProtocolName(v string) *ModifyHostsPortRequest
	GetProtocolName() *string
	SetRegionId(v string) *ModifyHostsPortRequest
	GetRegionId() *string
}

type ModifyHostsPortRequest struct {
	// The ID of the host for which you want to change the port. The value is a JSON string. You can add up to 100 host IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the IDs of hosts.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to change the port of the host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new port of the host. The port number must be an integer. Valid values: 22 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used to connect to the host. Valid values:
	//
	// 	- **SSH**
	//
	// 	- **RDP**
	//
	// This parameter is required.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The region ID of the bastion host for which you want to change the port of the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostsPortRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsPortRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *ModifyHostsPortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostsPortRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyHostsPortRequest) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ModifyHostsPortRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostsPortRequest) SetHostIds(v string) *ModifyHostsPortRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsPortRequest) SetInstanceId(v string) *ModifyHostsPortRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsPortRequest) SetPort(v string) *ModifyHostsPortRequest {
	s.Port = &v
	return s
}

func (s *ModifyHostsPortRequest) SetProtocolName(v string) *ModifyHostsPortRequest {
	s.ProtocolName = &v
	return s
}

func (s *ModifyHostsPortRequest) SetRegionId(v string) *ModifyHostsPortRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsPortRequest) Validate() error {
	return dara.Validate(s)
}
