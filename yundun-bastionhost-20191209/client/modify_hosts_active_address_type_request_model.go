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
	SetProjectId(v int64) *ModifyHostsActiveAddressTypeRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyHostsActiveAddressTypeRequest
	GetRegionId() *string
}

type ModifyHostsActiveAddressTypeRequest struct {
	// The new endpoint type of the host. Valid values:
	//
	// - **Public**: Internet
	//
	// - **Private**: private network
	//
	// This parameter is required.
	//
	// example:
	//
	// Private
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The IDs of the hosts whose endpoint types you want to modify. The value is a character string in JSON format. You can specify up to 100 host IDs.
	//
	// You can invoke the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The instance ID of the Bastionhost instance to which the host whose endpoint type you want to modify belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The area ID of the Bastionhost instance to which the host whose endpoint type you want to modify belongs.
	//
	// > For more information about the mapping between area IDs and area names, see [Regions and active zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ModifyHostsActiveAddressTypeRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ModifyHostsActiveAddressTypeRequest) SetProjectId(v int64) *ModifyHostsActiveAddressTypeRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetRegionId(v string) *ModifyHostsActiveAddressTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) Validate() error {
	return dara.Validate(s)
}
