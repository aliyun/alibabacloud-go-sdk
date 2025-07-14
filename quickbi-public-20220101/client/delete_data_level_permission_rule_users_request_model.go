// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelPermissionRuleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteUserModel(v string) *DeleteDataLevelPermissionRuleUsersRequest
	GetDeleteUserModel() *string
}

type DeleteDataLevelPermissionRuleUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
	DeleteUserModel *string `json:"DeleteUserModel,omitempty" xml:"DeleteUserModel,omitempty"`
}

func (s DeleteDataLevelPermissionRuleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersRequest) GetDeleteUserModel() *string {
	return s.DeleteUserModel
}

func (s *DeleteDataLevelPermissionRuleUsersRequest) SetDeleteUserModel(v string) *DeleteDataLevelPermissionRuleUsersRequest {
	s.DeleteUserModel = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersRequest) Validate() error {
	return dara.Validate(s)
}
