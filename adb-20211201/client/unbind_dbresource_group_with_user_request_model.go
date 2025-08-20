// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest
	GetDBClusterId() *string
	SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest
	GetGroupName() *string
	SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest
	GetGroupUser() *string
}

type UnbindDBResourceGroupWithUserRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// user1
	GroupUser *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
}

func (s UnbindDBResourceGroupWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnbindDBResourceGroupWithUserRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UnbindDBResourceGroupWithUserRequest) GetGroupUser() *string {
	return s.GroupUser
}

func (s *UnbindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) Validate() error {
	return dara.Validate(s)
}
