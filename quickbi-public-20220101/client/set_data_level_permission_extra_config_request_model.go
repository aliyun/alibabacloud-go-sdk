// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionExtraConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *SetDataLevelPermissionExtraConfigRequest
	GetCubeId() *string
	SetMissHitPolicy(v string) *SetDataLevelPermissionExtraConfigRequest
	GetMissHitPolicy() *string
	SetRuleType(v string) *SetDataLevelPermissionExtraConfigRequest
	GetRuleType() *string
}

type SetDataLevelPermissionExtraConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-******-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NONE
	MissHitPolicy *string `json:"MissHitPolicy,omitempty" xml:"MissHitPolicy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s SetDataLevelPermissionExtraConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *SetDataLevelPermissionExtraConfigRequest) GetMissHitPolicy() *string {
	return s.MissHitPolicy
}

func (s *SetDataLevelPermissionExtraConfigRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetCubeId(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.CubeId = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetMissHitPolicy(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.MissHitPolicy = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetRuleType(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.RuleType = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigRequest) Validate() error {
	return dara.Validate(s)
}
