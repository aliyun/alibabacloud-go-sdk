// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvRuleSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyNeed(v bool) *IsvRuleSaveRequest
	GetApplyNeed() *bool
	SetBookType(v string) *IsvRuleSaveRequest
	GetBookType() *string
	SetBookuserList(v []*IsvRuleSaveRequestBookuserList) *IsvRuleSaveRequest
	GetBookuserList() []*IsvRuleSaveRequestBookuserList
	SetRuleNeed(v bool) *IsvRuleSaveRequest
	GetRuleNeed() *bool
	SetStatus(v int32) *IsvRuleSaveRequest
	GetStatus() *int32
	SetUserId(v string) *IsvRuleSaveRequest
	GetUserId() *string
}

type IsvRuleSaveRequest struct {
	ApplyNeed *bool `json:"apply_need,omitempty" xml:"apply_need,omitempty"`
	// This parameter is required.
	BookType     *string                           `json:"book_type,omitempty" xml:"book_type,omitempty"`
	BookuserList []*IsvRuleSaveRequestBookuserList `json:"bookuser_list,omitempty" xml:"bookuser_list,omitempty" type:"Repeated"`
	RuleNeed     *bool                             `json:"rule_need,omitempty" xml:"rule_need,omitempty"`
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

func (s IsvRuleSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveRequest) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveRequest) GetApplyNeed() *bool {
	return s.ApplyNeed
}

func (s *IsvRuleSaveRequest) GetBookType() *string {
	return s.BookType
}

func (s *IsvRuleSaveRequest) GetBookuserList() []*IsvRuleSaveRequestBookuserList {
	return s.BookuserList
}

func (s *IsvRuleSaveRequest) GetRuleNeed() *bool {
	return s.RuleNeed
}

func (s *IsvRuleSaveRequest) GetStatus() *int32 {
	return s.Status
}

func (s *IsvRuleSaveRequest) GetUserId() *string {
	return s.UserId
}

func (s *IsvRuleSaveRequest) SetApplyNeed(v bool) *IsvRuleSaveRequest {
	s.ApplyNeed = &v
	return s
}

func (s *IsvRuleSaveRequest) SetBookType(v string) *IsvRuleSaveRequest {
	s.BookType = &v
	return s
}

func (s *IsvRuleSaveRequest) SetBookuserList(v []*IsvRuleSaveRequestBookuserList) *IsvRuleSaveRequest {
	s.BookuserList = v
	return s
}

func (s *IsvRuleSaveRequest) SetRuleNeed(v bool) *IsvRuleSaveRequest {
	s.RuleNeed = &v
	return s
}

func (s *IsvRuleSaveRequest) SetStatus(v int32) *IsvRuleSaveRequest {
	s.Status = &v
	return s
}

func (s *IsvRuleSaveRequest) SetUserId(v string) *IsvRuleSaveRequest {
	s.UserId = &v
	return s
}

func (s *IsvRuleSaveRequest) Validate() error {
	return dara.Validate(s)
}

type IsvRuleSaveRequestBookuserList struct {
	// This parameter is required.
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// This parameter is required.
	EntityType *int32 `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s IsvRuleSaveRequestBookuserList) String() string {
	return dara.Prettify(s)
}

func (s IsvRuleSaveRequestBookuserList) GoString() string {
	return s.String()
}

func (s *IsvRuleSaveRequestBookuserList) GetEntityId() *string {
	return s.EntityId
}

func (s *IsvRuleSaveRequestBookuserList) GetEntityType() *int32 {
	return s.EntityType
}

func (s *IsvRuleSaveRequestBookuserList) SetEntityId(v string) *IsvRuleSaveRequestBookuserList {
	s.EntityId = &v
	return s
}

func (s *IsvRuleSaveRequestBookuserList) SetEntityType(v int32) *IsvRuleSaveRequestBookuserList {
	s.EntityType = &v
	return s
}

func (s *IsvRuleSaveRequestBookuserList) Validate() error {
	return dara.Validate(s)
}
