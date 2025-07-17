// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertRulesShrinkRequest
	GetName() *string
	SetOwner(v string) *ListAlertRulesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListAlertRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int64) *ListAlertRulesShrinkRequest
	GetPageSize() *int64
	SetReceiver(v string) *ListAlertRulesShrinkRequest
	GetReceiver() *string
	SetTaskIdsShrink(v string) *ListAlertRulesShrinkRequest
	GetTaskIdsShrink() *string
	SetTypesShrink(v string) *ListAlertRulesShrinkRequest
	GetTypesShrink() *string
}

type ListAlertRulesShrinkRequest struct {
	// The name of the rule.
	//
	// example:
	//
	// error_rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// example:
	//
	// 1933790683****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Alibaba Cloud account used by the alert recipient.
	//
	// example:
	//
	// 1933790683****
	Receiver *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	// The IDs of the scheduling tasks.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The alert triggering condition.
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListAlertRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertRulesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertRulesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertRulesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertRulesShrinkRequest) GetReceiver() *string {
	return s.Receiver
}

func (s *ListAlertRulesShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *ListAlertRulesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListAlertRulesShrinkRequest) SetName(v string) *ListAlertRulesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetOwner(v string) *ListAlertRulesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetPageNumber(v int32) *ListAlertRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetPageSize(v int64) *ListAlertRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetReceiver(v string) *ListAlertRulesShrinkRequest {
	s.Receiver = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetTaskIdsShrink(v string) *ListAlertRulesShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) SetTypesShrink(v string) *ListAlertRulesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListAlertRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
