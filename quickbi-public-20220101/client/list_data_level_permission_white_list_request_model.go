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
	// Dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d5db23c-e4f2-49dd-a883-92285b48e14a
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// Type of row and column permission that the whitelist belongs to:
	//
	// - ROW_LEVEL: Row-level permission
	//
	// - COLUMN_LEVEL: Column-level permission
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
