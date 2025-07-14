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
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-***-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// Operation Type: You can set this parameter to one of the following values.
	//
	// 	- ADD: Add a whitelist
	//
	// 	- DELETE: deletes a whitelist.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of row-level permissions.
	//
	// 	- ROW_LEVEL: row-level permissions,
	//
	// 	- COLUMN_LEVEL: column-level permissions
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 43342***435,1553a****41231
	TargetIds *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	// Modify the type of the whitelist:
	//
	// 	- 1: user
	//
	// 	- 2: user group
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
