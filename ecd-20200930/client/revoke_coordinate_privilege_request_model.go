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
	// The coordination stream ID. This value is the `Coid` returned by the [ApplyCoordinationForMonitoring](~~ApplyCoordinationForMonitoring~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// co-fqsm6e8ee75w61fp9
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The username of the end user.
	//
	// example:
	//
	// zhangsan
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the coordination user.
	//
	// This parameter is required.
	//
	// example:
	//
	// TENANT_ADMIN
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// The UUID (unique identifier) of the device.
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
