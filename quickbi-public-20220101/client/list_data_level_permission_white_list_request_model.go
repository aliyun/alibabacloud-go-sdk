// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLevelPermissionWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *ListDataLevelPermissionWhiteListRequest
	GetCubeId() *string
	SetRuleType(v string) *ListDataLevelPermissionWhiteListRequest
	GetRuleType() *string
}

type ListDataLevelPermissionWhiteListRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d5db23c-e4f2-49dd-a883-92285b48e14a
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The type of row and column permission to which the whitelist belongs. Valid values:
	//
	// - ROW_LEVEL: row-level permission
	//
	// - COLUMN_LEVEL: column-level permission
	//
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListDataLevelPermissionWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *ListDataLevelPermissionWhiteListRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListDataLevelPermissionWhiteListRequest) SetCubeId(v string) *ListDataLevelPermissionWhiteListRequest {
	s.CubeId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListRequest) SetRuleType(v string) *ListDataLevelPermissionWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
