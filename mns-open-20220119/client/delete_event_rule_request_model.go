// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductName(v string) *DeleteEventRuleRequest
	GetProductName() *string
	SetRuleName(v string) *DeleteEventRuleRequest
	GetRuleName() *string
}

type DeleteEventRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleRequest) GetProductName() *string {
	return s.ProductName
}

func (s *DeleteEventRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteEventRuleRequest) SetProductName(v string) *DeleteEventRuleRequest {
	s.ProductName = &v
	return s
}

func (s *DeleteEventRuleRequest) SetRuleName(v string) *DeleteEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteEventRuleRequest) Validate() error {
	return dara.Validate(s)
}
