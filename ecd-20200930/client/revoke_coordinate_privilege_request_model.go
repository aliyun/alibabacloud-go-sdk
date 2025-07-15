// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCoordinatePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoId(v string) *RevokeCoordinatePrivilegeRequest
	GetCoId() *string
	SetEndUserId(v string) *RevokeCoordinatePrivilegeRequest
	GetEndUserId() *string
	SetRegionId(v string) *RevokeCoordinatePrivilegeRequest
	GetRegionId() *string
	SetUserType(v string) *RevokeCoordinatePrivilegeRequest
	GetUserType() *string
	SetUuid(v string) *RevokeCoordinatePrivilegeRequest
	GetUuid() *string
}

type RevokeCoordinatePrivilegeRequest struct {
	// The ID of the stream coordination task.
	//
	// This parameter is required.
	//
	// example:
	//
	// co-fqsm6e8ee75w61fp9
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// zhangsan
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of user who requires the coordinate permissions.
	//
	// Set the value to TENANT_ADMIN. Only tenant administrators can be granted with the coordinate permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// TENANT_ADMIN
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// The unique identifier of the client. If you use an Alibaba Cloud Workspace client, click **About*	- on the client logon page to view the identifier of the client.
	//
	// example:
	//
	// C78CA9E99315687575DD2844C1F3****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RevokeCoordinatePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeCoordinatePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *RevokeCoordinatePrivilegeRequest) GetCoId() *string {
	return s.CoId
}

func (s *RevokeCoordinatePrivilegeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RevokeCoordinatePrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeCoordinatePrivilegeRequest) GetUserType() *string {
	return s.UserType
}

func (s *RevokeCoordinatePrivilegeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RevokeCoordinatePrivilegeRequest) SetCoId(v string) *RevokeCoordinatePrivilegeRequest {
	s.CoId = &v
	return s
}

func (s *RevokeCoordinatePrivilegeRequest) SetEndUserId(v string) *RevokeCoordinatePrivilegeRequest {
	s.EndUserId = &v
	return s
}

func (s *RevokeCoordinatePrivilegeRequest) SetRegionId(v string) *RevokeCoordinatePrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeCoordinatePrivilegeRequest) SetUserType(v string) *RevokeCoordinatePrivilegeRequest {
	s.UserType = &v
	return s
}

func (s *RevokeCoordinatePrivilegeRequest) SetUuid(v string) *RevokeCoordinatePrivilegeRequest {
	s.Uuid = &v
	return s
}

func (s *RevokeCoordinatePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
