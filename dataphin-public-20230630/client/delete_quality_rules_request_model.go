// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteQualityRulesRequestDeleteCommand) *DeleteQualityRulesRequest
	GetDeleteCommand() *DeleteQualityRulesRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteQualityRulesRequest
	GetOpTenantId() *int64
}

type DeleteQualityRulesRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteQualityRulesRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualityRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityRulesRequest) GetDeleteCommand() *DeleteQualityRulesRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteQualityRulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityRulesRequest) SetDeleteCommand(v *DeleteQualityRulesRequestDeleteCommand) *DeleteQualityRulesRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteQualityRulesRequest) SetOpTenantId(v int64) *DeleteQualityRulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityRulesRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteQualityRulesRequestDeleteCommand struct {
	// This parameter is required.
	RuleIdList []*int64 `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
}

func (s DeleteQualityRulesRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRulesRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteQualityRulesRequestDeleteCommand) GetRuleIdList() []*int64 {
	return s.RuleIdList
}

func (s *DeleteQualityRulesRequestDeleteCommand) SetRuleIdList(v []*int64) *DeleteQualityRulesRequestDeleteCommand {
	s.RuleIdList = v
	return s
}

func (s *DeleteQualityRulesRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
