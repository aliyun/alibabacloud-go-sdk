// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ModifyAssetGroupRequest
	GetGroupId() *int64
	SetSourceIp(v string) *ModifyAssetGroupRequest
	GetSourceIp() *string
	SetUuids(v string) *ModifyAssetGroupRequest
	GetUuids() *string
}

type ModifyAssetGroupRequest struct {
	// The ID of the new server group to which the servers belong.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of server groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9586199
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 10.12.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUIDs of the servers for which you want to change the server group. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 076a446d-df7d-424c-bdc5-bb5dc7f1****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyAssetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAssetGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyAssetGroupRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyAssetGroupRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyAssetGroupRequest) SetGroupId(v int64) *ModifyAssetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyAssetGroupRequest) SetSourceIp(v string) *ModifyAssetGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAssetGroupRequest) SetUuids(v string) *ModifyAssetGroupRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyAssetGroupRequest) Validate() error {
	return dara.Validate(s)
}
