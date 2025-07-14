// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataLevelPermissionRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleModel(v string) *SetDataLevelPermissionRuleConfigRequest
	GetRuleModel() *string
}

type SetDataLevelPermissionRuleConfigRequest struct {
	// This parameter is required.
	RuleModel *string `json:"RuleModel,omitempty" xml:"RuleModel,omitempty"`
}

func (s SetDataLevelPermissionRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigRequest) GetRuleModel() *string {
	return s.RuleModel
}

func (s *SetDataLevelPermissionRuleConfigRequest) SetRuleModel(v string) *SetDataLevelPermissionRuleConfigRequest {
	s.RuleModel = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}
