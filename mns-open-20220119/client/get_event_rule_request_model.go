// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductName(v string) *GetEventRuleRequest
	GetProductName() *string
	SetRuleName(v string) *GetEventRuleRequest
	GetRuleName() *string
}

type GetEventRuleRequest struct {
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

func (s GetEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleRequest) GoString() string {
	return s.String()
}

func (s *GetEventRuleRequest) GetProductName() *string {
	return s.ProductName
}

func (s *GetEventRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *GetEventRuleRequest) SetProductName(v string) *GetEventRuleRequest {
	s.ProductName = &v
	return s
}

func (s *GetEventRuleRequest) SetRuleName(v string) *GetEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *GetEventRuleRequest) Validate() error {
	return dara.Validate(s)
}
