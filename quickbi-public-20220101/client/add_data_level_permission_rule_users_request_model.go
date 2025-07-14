// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionRuleUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddUserModel(v string) *AddDataLevelPermissionRuleUsersRequest
	GetAddUserModel() *string
}

type AddDataLevelPermissionRuleUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","addModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
	AddUserModel *string `json:"AddUserModel,omitempty" xml:"AddUserModel,omitempty"`
}

func (s AddDataLevelPermissionRuleUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersRequest) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersRequest) GetAddUserModel() *string {
	return s.AddUserModel
}

func (s *AddDataLevelPermissionRuleUsersRequest) SetAddUserModel(v string) *AddDataLevelPermissionRuleUsersRequest {
	s.AddUserModel = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersRequest) Validate() error {
	return dara.Validate(s)
}
