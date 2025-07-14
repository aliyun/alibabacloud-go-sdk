// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLevelRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *DeleteDataLevelRuleConfigRequest
	GetCubeId() *string
	SetRuleId(v string) *DeleteDataLevelRuleConfigRequest
	GetRuleId() *string
}

type DeleteDataLevelRuleConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a5bb24da-****-a891683e14da
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteDataLevelRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLevelRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *DeleteDataLevelRuleConfigRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteDataLevelRuleConfigRequest) SetCubeId(v string) *DeleteDataLevelRuleConfigRequest {
	s.CubeId = &v
	return s
}

func (s *DeleteDataLevelRuleConfigRequest) SetRuleId(v string) *DeleteDataLevelRuleConfigRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteDataLevelRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}
