// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsActiveAddressTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddressType(v string) *ModifyHostsActiveAddressTypeRequest
	GetActiveAddressType() *string
	SetHostIds(v string) *ModifyHostsActiveAddressTypeRequest
	GetHostIds() *string
	SetInstanceId(v string) *ModifyHostsActiveAddressTypeRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHostsActiveAddressTypeRequest
	GetRegionId() *string
}

type ModifyHostsActiveAddressTypeRequest struct {
	// The new portal type of the host. Valid values:
	//
	// 	- **Public**: public portal
	//
	// 	- **Private**: internal portal
	//
	// This parameter is required.
	//
	// example:
	//
	// Private
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The ID of the host for which you want to change the portal type. The value is a JSON string. You can add up to 100 host IDs.
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to change the portal type of the host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to change the portal type of the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostsActiveAddressTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeRequest) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ModifyHostsActiveAddressTypeRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *ModifyHostsActiveAddressTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostsActiveAddressTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostsActiveAddressTypeRequest) SetActiveAddressType(v string) *ModifyHostsActiveAddressTypeRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetHostIds(v string) *ModifyHostsActiveAddressTypeRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetInstanceId(v string) *ModifyHostsActiveAddressTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetRegionId(v string) *ModifyHostsActiveAddressTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) Validate() error {
	return dara.Validate(s)
}
