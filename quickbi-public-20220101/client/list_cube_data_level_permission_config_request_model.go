// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCubeDataLevelPermissionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *ListCubeDataLevelPermissionConfigRequest
	GetCubeId() *string
	SetRuleType(v string) *ListCubeDataLevelPermissionConfigRequest
	GetRuleType() *string
}

type ListCubeDataLevelPermissionConfigRequest struct {
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The type of the dataset row and column permission. Valid values:
	//
	// 	- ROW_LEVEL: row-level permissions
	//
	// 	- COLUMN_LEVEL: column-level permissions
	//
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListCubeDataLevelPermissionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCubeDataLevelPermissionConfigRequest) GoString() string {
	return s.String()
}

func (s *ListCubeDataLevelPermissionConfigRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *ListCubeDataLevelPermissionConfigRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListCubeDataLevelPermissionConfigRequest) SetCubeId(v string) *ListCubeDataLevelPermissionConfigRequest {
	s.CubeId = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigRequest) SetRuleType(v string) *ListCubeDataLevelPermissionConfigRequest {
	s.RuleType = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigRequest) Validate() error {
	return dara.Validate(s)
}
