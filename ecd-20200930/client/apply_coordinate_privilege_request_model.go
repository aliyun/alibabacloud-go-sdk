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
	// The coordination stream ID. This value is the `Coid` returned by the [ApplyCoordinationForMonitoring](~~ApplyCoordinationForMonitoring~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// co-fqsm6e8ee75w6****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The username of the end user.
	//
	// example:
	//
	// alice
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
	// 3E14A18BD4D088504B9F8A8751AB****
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
