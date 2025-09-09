// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *GetHostRequest
	GetHostId() *string
	SetInstanceId(v string) *GetHostRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetHostRequest
	GetRegionId() *string
}

type GetHostRequest struct {
	// The ID of the host that you want to query. You can specify only one host ID.
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host in which you want to query the host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to query the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostRequest) GoString() string {
	return s.String()
}

func (s *GetHostRequest) GetHostId() *string {
	return s.HostId
}

func (s *GetHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostRequest) SetHostId(v string) *GetHostRequest {
	s.HostId = &v
	return s
}

func (s *GetHostRequest) SetInstanceId(v string) *GetHostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostRequest) SetRegionId(v string) *GetHostRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostRequest) Validate() error {
	return dara.Validate(s)
}
