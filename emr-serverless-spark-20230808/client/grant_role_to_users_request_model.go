// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleToUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleArn(v string) *GrantRoleToUsersRequest
	GetRoleArn() *string
	SetUserArns(v []*string) *GrantRoleToUsersRequest
	GetUserArns() []*string
	SetRegionId(v string) *GrantRoleToUsersRequest
	GetRegionId() *string
}

type GrantRoleToUsersRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:emr::w-975bcfda9625****:role/Owner
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The user ARNs.
	UserArns []*string `json:"userArns,omitempty" xml:"userArns,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GrantRoleToUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleToUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GrantRoleToUsersRequest) GetUserArns() []*string {
	return s.UserArns
}

func (s *GrantRoleToUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantRoleToUsersRequest) SetRoleArn(v string) *GrantRoleToUsersRequest {
	s.RoleArn = &v
	return s
}

func (s *GrantRoleToUsersRequest) SetUserArns(v []*string) *GrantRoleToUsersRequest {
	s.UserArns = v
	return s
}

func (s *GrantRoleToUsersRequest) SetRegionId(v string) *GrantRoleToUsersRequest {
	s.RegionId = &v
	return s
}

func (s *GrantRoleToUsersRequest) Validate() error {
	return dara.Validate(s)
}
