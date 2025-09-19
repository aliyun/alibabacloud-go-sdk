// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *GetSasContainerWebDefenseRuleApplicationRequest
	GetRuleId() *int64
}

type GetSasContainerWebDefenseRuleApplicationRequest struct {
	// The ID of the rule.
	//
	// >  You can call the ListSasContainerWebDefenseRule operation to query the IDs of rules.
	//
	// example:
	//
	// 400599
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetSasContainerWebDefenseRuleApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleApplicationRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetSasContainerWebDefenseRuleApplicationRequest) SetRuleId(v int64) *GetSasContainerWebDefenseRuleApplicationRequest {
	s.RuleId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationRequest) Validate() error {
	return dara.Validate(s)
}
