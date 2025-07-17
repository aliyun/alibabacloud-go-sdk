// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertRulesRequest
	GetName() *string
	SetOwner(v string) *ListAlertRulesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListAlertRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int64) *ListAlertRulesRequest
	GetPageSize() *int64
	SetReceiver(v string) *ListAlertRulesRequest
	GetReceiver() *string
	SetTaskIds(v []*int64) *ListAlertRulesRequest
	GetTaskIds() []*int64
	SetTypes(v []*string) *ListAlertRulesRequest
	GetTypes() []*string
}

type ListAlertRulesRequest struct {
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The alert triggering condition.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertRulesRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertRulesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertRulesRequest) GetReceiver() *string {
	return s.Receiver
}

func (s *ListAlertRulesRequest) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *ListAlertRulesRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListAlertRulesRequest) SetName(v string) *ListAlertRulesRequest {
	s.Name = &v
	return s
}

func (s *ListAlertRulesRequest) SetOwner(v string) *ListAlertRulesRequest {
	s.Owner = &v
	return s
}

func (s *ListAlertRulesRequest) SetPageNumber(v int32) *ListAlertRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRulesRequest) SetPageSize(v int64) *ListAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertRulesRequest) SetReceiver(v string) *ListAlertRulesRequest {
	s.Receiver = &v
	return s
}

func (s *ListAlertRulesRequest) SetTaskIds(v []*int64) *ListAlertRulesRequest {
	s.TaskIds = v
	return s
}

func (s *ListAlertRulesRequest) SetTypes(v []*string) *ListAlertRulesRequest {
	s.Types = v
	return s
}

func (s *ListAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
