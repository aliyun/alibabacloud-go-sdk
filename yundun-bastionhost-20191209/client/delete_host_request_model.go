// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *DeleteHostRequest
	GetHostId() *string
	SetInstanceId(v string) *DeleteHostRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteHostRequest
	GetRegionId() *string
}

type DeleteHostRequest struct {
	// The ID of the host that you want to delete.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host on which you want to delete the host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to delete the host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostRequest) GetHostId() *string {
	return s.HostId
}

func (s *DeleteHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHostRequest) SetHostId(v string) *DeleteHostRequest {
	s.HostId = &v
	return s
}

func (s *DeleteHostRequest) SetInstanceId(v string) *DeleteHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostRequest) SetRegionId(v string) *DeleteHostRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostRequest) Validate() error {
	return dara.Validate(s)
}
