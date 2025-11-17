// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLevelPermissionWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataLevelPermissionWhiteListResponseBody
	GetRequestId() *string
	SetResult(v *ListDataLevelPermissionWhiteListResponseBodyResult) *ListDataLevelPermissionWhiteListResponseBody
	GetResult() *ListDataLevelPermissionWhiteListResponseBodyResult
	SetSuccess(v bool) *ListDataLevelPermissionWhiteListResponseBody
	GetSuccess() *bool
}

type ListDataLevelPermissionWhiteListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whitelist for the specified row-level permission type.
	Result *ListDataLevelPermissionWhiteListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLevelPermissionWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLevelPermissionWhiteListResponseBody) GetResult() *ListDataLevelPermissionWhiteListResponseBodyResult {
	return s.Result
}

func (s *ListDataLevelPermissionWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *ListDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetResult(v *ListDataLevelPermissionWhiteListResponseBodyResult) *ListDataLevelPermissionWhiteListResponseBody {
	s.Result = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *ListDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataLevelPermissionWhiteListResponseBodyResult struct {
	// Dataset ID.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// Type of dataset row and column permissions. Possible values:
	//
	// - ROW_LEVEL: Row-level permission
	//
	// - COLUMN_LEVEL: Column-level permission
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// Whitelist information.
	UsersModel *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel `json:"UsersModel,omitempty" xml:"UsersModel,omitempty" type:"Struct"`
}

func (s ListDataLevelPermissionWhiteListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) GetCubeId() *string {
	return s.CubeId
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) GetRuleType() *string {
	return s.RuleType
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) GetUsersModel() *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel {
	return s.UsersModel
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetCubeId(v string) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetRuleType(v string) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.RuleType = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetUsersModel(v *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.UsersModel = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) Validate() error {
	if s.UsersModel != nil {
		if err := s.UsersModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataLevelPermissionWhiteListResponseBodyResultUsersModel struct {
	// UserGroups.
	UserGroups []*string `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
	// Users.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) String() string {
	return dara.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) GetUserGroups() []*string {
	return s.UserGroups
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) GetUsers() []*string {
	return s.Users
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) SetUserGroups(v []*string) *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel {
	s.UserGroups = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) SetUsers(v []*string) *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel {
	s.Users = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) Validate() error {
	return dara.Validate(s)
}
