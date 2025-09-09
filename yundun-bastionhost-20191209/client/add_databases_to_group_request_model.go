// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatabasesToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseIds(v []*string) *AddDatabasesToGroupRequest
	GetDatabaseIds() []*string
	SetHostGroupId(v string) *AddDatabasesToGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *AddDatabasesToGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AddDatabasesToGroupRequest
	GetRegionId() *string
}

type AddDatabasesToGroupRequest struct {
	// An array that consists of the database IDs.
	//
	// This parameter is required.
	DatabaseIds []*string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
	// The ID of the asset group to which you want to add the databases.
	//
	// >  You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the asset group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDatabasesToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDatabasesToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDatabasesToGroupRequest) GetDatabaseIds() []*string {
	return s.DatabaseIds
}

func (s *AddDatabasesToGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AddDatabasesToGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddDatabasesToGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDatabasesToGroupRequest) SetDatabaseIds(v []*string) *AddDatabasesToGroupRequest {
	s.DatabaseIds = v
	return s
}

func (s *AddDatabasesToGroupRequest) SetHostGroupId(v string) *AddDatabasesToGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *AddDatabasesToGroupRequest) SetInstanceId(v string) *AddDatabasesToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDatabasesToGroupRequest) SetRegionId(v string) *AddDatabasesToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddDatabasesToGroupRequest) Validate() error {
	return dara.Validate(s)
}
