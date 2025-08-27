// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvRuleSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyNeed(v bool) *IsvRuleSaveShrinkRequest
	GetApplyNeed() *bool
	SetBookType(v string) *IsvRuleSaveShrinkRequest
	GetBookType() *string
	SetBookuserListShrink(v string) *IsvRuleSaveShrinkRequest
	GetBookuserListShrink() *string
	SetRuleNeed(v bool) *IsvRuleSaveShrinkRequest
	GetRuleNeed() *bool
	SetStatus(v int32) *IsvRuleSaveShrinkRequest
	GetStatus() *int32
	SetUserId(v string) *IsvRuleSaveShrinkRequest
	GetUserId() *string
}

type IsvRuleSaveShrinkRequest struct {
	ApplyNeed *bool `json:"apply_need,omitempty" xml:"apply_need,omitempty"`
	// This parameter is required.
	BookType           *string `json:"book_type,omitempty" xml:"book_type,omitempty"`
	BookuserListShrink *string `json:"bookuser_list,omitempty" xml:"bookuser_list,omitempty"`
	RuleNeed           *bool   `json:"rule_need,omitempty" xml:"rule_need,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s IsvRuleSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveShrinkRequest) GetApplyNeed() *bool {
	return s.ApplyNeed
}

func (s *IsvRuleSaveShrinkRequest) GetBookType() *string {
	return s.BookType
}

func (s *IsvRuleSaveShrinkRequest) GetBookuserListShrink() *string {
	return s.BookuserListShrink
}

func (s *IsvRuleSaveShrinkRequest) GetRuleNeed() *bool {
	return s.RuleNeed
}

func (s *IsvRuleSaveShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *IsvRuleSaveShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *IsvRuleSaveShrinkRequest) SetApplyNeed(v bool) *IsvRuleSaveShrinkRequest {
	s.ApplyNeed = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) SetBookType(v string) *IsvRuleSaveShrinkRequest {
	s.BookType = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) SetBookuserListShrink(v string) *IsvRuleSaveShrinkRequest {
	s.BookuserListShrink = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) SetRuleNeed(v bool) *IsvRuleSaveShrinkRequest {
	s.RuleNeed = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) SetStatus(v int32) *IsvRuleSaveShrinkRequest {
	s.Status = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) SetUserId(v string) *IsvRuleSaveShrinkRequest {
	s.UserId = &v
	return s
}

func (s *IsvRuleSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
