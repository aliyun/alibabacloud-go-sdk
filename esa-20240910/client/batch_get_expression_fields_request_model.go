// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetExpressionFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpressions(v []*BatchGetExpressionFieldsRequestExpressions) *BatchGetExpressionFieldsRequest
	GetExpressions() []*BatchGetExpressionFieldsRequestExpressions
	SetPhase(v string) *BatchGetExpressionFieldsRequest
	GetPhase() *string
	SetSiteId(v int64) *BatchGetExpressionFieldsRequest
	GetSiteId() *int64
}

type BatchGetExpressionFieldsRequest struct {
	// List of expressions.
	//
	// example:
	//
	// http_bot
	Expressions []*BatchGetExpressionFieldsRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// WAF Phase
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Site ID
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s BatchGetExpressionFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsRequest) GetExpressions() []*BatchGetExpressionFieldsRequestExpressions {
	return s.Expressions
}

func (s *BatchGetExpressionFieldsRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchGetExpressionFieldsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchGetExpressionFieldsRequest) SetExpressions(v []*BatchGetExpressionFieldsRequestExpressions) *BatchGetExpressionFieldsRequest {
	s.Expressions = v
	return s
}

func (s *BatchGetExpressionFieldsRequest) SetPhase(v string) *BatchGetExpressionFieldsRequest {
	s.Phase = &v
	return s
}

func (s *BatchGetExpressionFieldsRequest) SetSiteId(v int64) *BatchGetExpressionFieldsRequest {
	s.SiteId = &v
	return s
}

func (s *BatchGetExpressionFieldsRequest) Validate() error {
	return dara.Validate(s)
}

type BatchGetExpressionFieldsRequestExpressions struct {
	// Content of the expression.
	//
	// example:
	//
	// ip.src eq 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The sequence number of the expression.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchGetExpressionFieldsRequestExpressions) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsRequestExpressions) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsRequestExpressions) GetExpression() *string {
	return s.Expression
}

func (s *BatchGetExpressionFieldsRequestExpressions) GetId() *int64 {
	return s.Id
}

func (s *BatchGetExpressionFieldsRequestExpressions) SetExpression(v string) *BatchGetExpressionFieldsRequestExpressions {
	s.Expression = &v
	return s
}

func (s *BatchGetExpressionFieldsRequestExpressions) SetId(v int64) *BatchGetExpressionFieldsRequestExpressions {
	s.Id = &v
	return s
}

func (s *BatchGetExpressionFieldsRequestExpressions) Validate() error {
	return dara.Validate(s)
}
