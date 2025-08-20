// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest
	GetDBClusterId() *string
	SetGroupName(v string) *BindDBResourceGroupWithUserRequest
	GetGroupName() *string
	SetGroupUser(v string) *BindDBResourceGroupWithUserRequest
	GetGroupUser() *string
}

type BindDBResourceGroupWithUserRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the database account. It can be a standard account or a privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// accout
	GroupUser *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
}

func (s BindDBResourceGroupWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindDBResourceGroupWithUserRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *BindDBResourceGroupWithUserRequest) GetGroupUser() *string {
	return s.GroupUser
}

func (s *BindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupName(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupUser(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) Validate() error {
	return dara.Validate(s)
}
