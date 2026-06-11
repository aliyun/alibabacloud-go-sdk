// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataLevelPermissionWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *AddDataLevelPermissionWhiteListRequest
	GetCubeId() *string
	SetOperateType(v string) *AddDataLevelPermissionWhiteListRequest
	GetOperateType() *string
	SetRuleType(v string) *AddDataLevelPermissionWhiteListRequest
	GetRuleType() *string
	SetTargetIds(v string) *AddDataLevelPermissionWhiteListRequest
	GetTargetIds() *string
	SetTargetType(v string) *AddDataLevelPermissionWhiteListRequest
	GetTargetType() *string
}

type AddDataLevelPermissionWhiteListRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-***-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The operation to perform. Valid values:
	//
	// - ADD: adds users or user groups to the whitelist.
	//
	// - DELETE: removes users or user groups from the whitelist.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of permission. Valid values:
	//
	// - ROW_LEVEL: row-level permission
	//
	// - COLUMN_LEVEL: column-level permission
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The IDs of the users or user groups to add to the whitelist.
	//
	// - If you set TargetType to 1 (user), specify the user IDs.
	//
	// - When `TargetType=2` (user group), the value is the user group ID.
	//
	// example:
	//
	// 43342***435,1553a****41231
	TargetIds *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	// The type of object to add to the whitelist. Valid values:
	//
	// - 1: user
	//
	// - 2: user group
	//
	// example:
	//
	// 1
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AddDataLevelPermissionWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *AddDataLevelPermissionWhiteListRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *AddDataLevelPermissionWhiteListRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *AddDataLevelPermissionWhiteListRequest) GetTargetIds() *string {
	return s.TargetIds
}

func (s *AddDataLevelPermissionWhiteListRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *AddDataLevelPermissionWhiteListRequest) SetCubeId(v string) *AddDataLevelPermissionWhiteListRequest {
	s.CubeId = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetOperateType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.OperateType = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetRuleType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetTargetIds(v string) *AddDataLevelPermissionWhiteListRequest {
	s.TargetIds = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetTargetType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.TargetType = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
