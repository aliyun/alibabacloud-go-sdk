// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDatabasesFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseIds(v []*string) *RemoveDatabasesFromGroupRequest
	GetDatabaseIds() []*string
	SetHostGroupId(v string) *RemoveDatabasesFromGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *RemoveDatabasesFromGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *RemoveDatabasesFromGroupRequest
	GetRegionId() *string
}

type RemoveDatabasesFromGroupRequest struct {
	// The IDs of the databases that you want to remove.
	//
	// This parameter is required.
	DatabaseIds []*string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
	// The ID of the asset group from which you want to remove databases.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the asset group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host whose asset group you want to manage.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to manage.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveDatabasesFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDatabasesFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveDatabasesFromGroupRequest) GetDatabaseIds() []*string {
	return s.DatabaseIds
}

func (s *RemoveDatabasesFromGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *RemoveDatabasesFromGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveDatabasesFromGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveDatabasesFromGroupRequest) SetDatabaseIds(v []*string) *RemoveDatabasesFromGroupRequest {
	s.DatabaseIds = v
	return s
}

func (s *RemoveDatabasesFromGroupRequest) SetHostGroupId(v string) *RemoveDatabasesFromGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *RemoveDatabasesFromGroupRequest) SetInstanceId(v string) *RemoveDatabasesFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveDatabasesFromGroupRequest) SetRegionId(v string) *RemoveDatabasesFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveDatabasesFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
