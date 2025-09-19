// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAssetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *CreateOrUpdateAssetGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *CreateOrUpdateAssetGroupRequest
	GetGroupName() *string
	SetUuids(v string) *CreateOrUpdateAssetGroupRequest
	GetUuids() *string
}

type CreateOrUpdateAssetGroupRequest struct {
	// The ID of the server group for which you want to add to or remove servers.
	//
	// >  To modify the mapping between an asset and an asset group, you must provide the ID of the asset group. You can call the [DescribeAllGroups](~~DescribeAllGroups~~) to query the IDs of asset groups. If you do not configure this parameter when you call this operation, an asset group is created.
	//
	// example:
	//
	// 55426
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server group that you want to create or the server group for which you want to add or remove a server.
	//
	// >  To modify the mapping between a server and a server group, you must provide the name of the server group. You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the names of server groups. If you do not configure GroupID when you call this operation, a server group is created. In this case, you must configure GroupName.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The UUID of the server in the server group that you want to create or the server group for which you want to add or remove servers. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// aq12-***,s23***
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s CreateOrUpdateAssetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateAssetGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOrUpdateAssetGroupRequest) GetUuids() *string {
	return s.Uuids
}

func (s *CreateOrUpdateAssetGroupRequest) SetGroupId(v int64) *CreateOrUpdateAssetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) SetGroupName(v string) *CreateOrUpdateAssetGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) SetUuids(v string) *CreateOrUpdateAssetGroupRequest {
	s.Uuids = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) Validate() error {
	return dara.Validate(s)
}
