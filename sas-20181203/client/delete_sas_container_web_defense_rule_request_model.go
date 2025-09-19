// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSasContainerWebDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *DeleteSasContainerWebDefenseRuleRequest
	GetRuleId() *int64
}

type DeleteSasContainerWebDefenseRuleRequest struct {
	// The rule ID.
	//
	// >  You can call the ListContainerWebDefenseRule operation to query the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400597
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteSasContainerWebDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSasContainerWebDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSasContainerWebDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteSasContainerWebDefenseRuleRequest) SetRuleId(v int64) *DeleteSasContainerWebDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteSasContainerWebDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
