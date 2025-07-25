// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWhiteListModel(v string) *SetDataLevelPermissionWhiteListRequest
	GetWhiteListModel() *string
}

type SetDataLevelPermissionWhiteListRequest struct {
	// { "ruleType": "ROW_LEVEL", // The row-level permission type. "usersModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // The ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Quick BI the UserID of the user. "Huang***3fa822" ] }, "cubeId": "7c7223ae-31d1-4d2f-b11f-3c744528014b" }
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ruleType":"ROW_LEVEL","usersModel":{"userGroups":["26edcb76-****-bdbab78267cb","187e6dd5-1611-4cf7-a034-1a93bd5fecf9"],"users":["4334***358","Huang***3fa822"]},"cubeId":"7c7223ae-****44528014b"}
	WhiteListModel *string `json:"WhiteListModel,omitempty" xml:"WhiteListModel,omitempty"`
}

func (s SetDataLevelPermissionWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListRequest) GetWhiteListModel() *string {
	return s.WhiteListModel
}

func (s *SetDataLevelPermissionWhiteListRequest) SetWhiteListModel(v string) *SetDataLevelPermissionWhiteListRequest {
	s.WhiteListModel = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
