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
	SetInstanceId(v string) *BatchGetExpressionFieldsShrinkRequest
	GetInstanceId() *string
	SetKind(v string) *BatchGetExpressionFieldsShrinkRequest
	GetKind() *string
	SetPhase(v string) *BatchGetExpressionFieldsShrinkRequest
	GetPhase() *string
	SetPlanNameEn(v string) *BatchGetExpressionFieldsShrinkRequest
	GetPlanNameEn() *string
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
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Kind              *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// WAF Phase
	//
	// example:
	//
	// http_bot
	Phase      *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	PlanNameEn *string `json:"PlanNameEn,omitempty" xml:"PlanNameEn,omitempty"`
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

func (s *BatchGetExpressionFieldsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetKind() *string {
	return s.Kind
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetPlanNameEn() *string {
	return s.PlanNameEn
}

func (s *BatchGetExpressionFieldsShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetExpressionsShrink(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.ExpressionsShrink = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetInstanceId(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetKind(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.Kind = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetPhase(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetPlanNameEn(v string) *BatchGetExpressionFieldsShrinkRequest {
	s.PlanNameEn = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) SetSiteId(v int64) *BatchGetExpressionFieldsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchGetExpressionFieldsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
