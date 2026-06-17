// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleNames(v []*string) *DeleteEventRulesRequest
	GetRuleNames() []*string
}

type DeleteEventRulesRequest struct {
	// The names of the event-triggered alert rules to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule1
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteEventRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesRequest) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *DeleteEventRulesRequest) SetRuleNames(v []*string) *DeleteEventRulesRequest {
	s.RuleNames = v
	return s
}

func (s *DeleteEventRulesRequest) Validate() error {
	return dara.Validate(s)
}
