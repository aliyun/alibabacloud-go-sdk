// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinatePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoId(v string) *ApplyCoordinatePrivilegeRequest
	GetCoId() *string
	SetEndUserId(v string) *ApplyCoordinatePrivilegeRequest
	GetEndUserId() *string
	SetRegionId(v string) *ApplyCoordinatePrivilegeRequest
	GetRegionId() *string
	SetUserType(v string) *ApplyCoordinatePrivilegeRequest
	GetUserType() *string
	SetUuid(v string) *ApplyCoordinatePrivilegeRequest
	GetUuid() *string
}

type ApplyCoordinatePrivilegeRequest struct {
	// The ID of the application for the coordinate permissions.
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
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of user who requires the coordinate permissions.
	//
	// Valid value: TENANT_ADMIN.
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
	// 8b241d415da244a6936d6d6fa4f20f4d
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyCoordinatePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinatePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *ApplyCoordinatePrivilegeRequest) GetCoId() *string {
	return s.CoId
}

func (s *ApplyCoordinatePrivilegeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ApplyCoordinatePrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyCoordinatePrivilegeRequest) GetUserType() *string {
	return s.UserType
}

func (s *ApplyCoordinatePrivilegeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ApplyCoordinatePrivilegeRequest) SetCoId(v string) *ApplyCoordinatePrivilegeRequest {
	s.CoId = &v
	return s
}

func (s *ApplyCoordinatePrivilegeRequest) SetEndUserId(v string) *ApplyCoordinatePrivilegeRequest {
	s.EndUserId = &v
	return s
}

func (s *ApplyCoordinatePrivilegeRequest) SetRegionId(v string) *ApplyCoordinatePrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyCoordinatePrivilegeRequest) SetUserType(v string) *ApplyCoordinatePrivilegeRequest {
	s.UserType = &v
	return s
}

func (s *ApplyCoordinatePrivilegeRequest) SetUuid(v string) *ApplyCoordinatePrivilegeRequest {
	s.Uuid = &v
	return s
}

func (s *ApplyCoordinatePrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
