// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *ListRulesRequest
	GetEventBusName() *string
	SetLimit(v int32) *ListRulesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListRulesRequest
	GetNextToken() *string
	SetRuleNamePrefix(v string) *ListRulesRequest
	GetRuleNamePrefix() *string
}

type ListRulesRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of entries to be returned in a single call. You can use this parameter and the NextToken parameter to implement paging. A maximum of 100 entries can be returned in a single call.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// If you set the Limit parameter and excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix of the rule name.
	//
	// example:
	//
	// test
	RuleNamePrefix *string `json:"RuleNamePrefix,omitempty" xml:"RuleNamePrefix,omitempty"`
}

func (s ListRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListRulesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRulesRequest) GetRuleNamePrefix() *string {
	return s.RuleNamePrefix
}

func (s *ListRulesRequest) SetEventBusName(v string) *ListRulesRequest {
	s.EventBusName = &v
	return s
}

func (s *ListRulesRequest) SetLimit(v int32) *ListRulesRequest {
	s.Limit = &v
	return s
}

func (s *ListRulesRequest) SetNextToken(v string) *ListRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRulesRequest) SetRuleNamePrefix(v string) *ListRulesRequest {
	s.RuleNamePrefix = &v
	return s
}

func (s *ListRulesRequest) Validate() error {
	return dara.Validate(s)
}
