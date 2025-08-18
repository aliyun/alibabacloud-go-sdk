// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetExpressionFieldsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpressionsShrink(v string) *BatchGetExpressionFieldsShrinkRequest
	GetExpressionsShrink() *string
	SetPhase(v string) *BatchGetExpressionFieldsShrinkRequest
	GetPhase() *string
	SetSiteId(v int64) *BatchGetExpressionFieldsShrinkRequest
	GetSiteId() *int64
}

type BatchGetExpressionFieldsShrinkRequest struct {
	// List of expressions.
	//
	// example:
	//
	// http_bot
	ExpressionsShrink *string `json:"Expressions,omitempty" xml:"Expressions,omitempty"`
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

func (s BatchGetExpressionFieldsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetExpressionsShrink() *string {
	return s.ExpressionsShrink
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetExpressionsShrink(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.ExpressionsShrink = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetPhase(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetSiteId(v int64) *BatchGetExpressionFieldsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
