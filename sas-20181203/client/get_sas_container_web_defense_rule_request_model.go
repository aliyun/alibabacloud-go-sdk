// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *GetSasContainerWebDefenseRuleRequest
	GetRuleId() *int64
}

type GetSasContainerWebDefenseRuleRequest struct {
	// Rule ID.
	//
	// > You can call the [ListSasContainerWebDefenseRule](~~ListSasContainerWebDefenseRule~~) API to get this parameter.
	//
	// example:
	//
	// 1600009
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetSasContainerWebDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetSasContainerWebDefenseRuleRequest) SetRuleId(v int64) *GetSasContainerWebDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
