// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTarget(v string) *ListRemindsRequest
	GetAlertTarget() *string
	SetFounder(v string) *ListRemindsRequest
	GetFounder() *string
	SetNodeId(v int64) *ListRemindsRequest
	GetNodeId() *int64
	SetPageNumber(v int32) *ListRemindsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRemindsRequest
	GetPageSize() *int32
	SetRemindTypes(v string) *ListRemindsRequest
	GetRemindTypes() *string
	SetSearchText(v string) *ListRemindsRequest
	GetSearchText() *string
}

type ListRemindsRequest struct {
	// The ID of the Alibaba Cloud account that is used to receive alert notifications.
	//
	// example:
	//
	// 9527952795****
	AlertTarget *string `json:"AlertTarget,omitempty" xml:"AlertTarget,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the custom alert rules.
	//
	// example:
	//
	// 9527952795****
	Founder *string `json:"Founder,omitempty" xml:"Founder,omitempty"`
	// The ID of the node to which the custom alert rules are applied. You can use the ID to search for the custom alert rules that are applied to the node.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of the page to return. Valid values: 1 to 30. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The conditions that trigger an alert for the node. Valid values: FINISHED, UNFINISHED, ERROR, CYCLE_UNFINISHED, and TIMEOUT. The value FINISHED indicates that the node finishes running. The value UNFINISHED indicates that the node is still running at the specified point in time. The value ERROR indicates that an error occurs when the node is running. The value CYCLE_UNFINISHED indicates that the node does not finish running in the specified scheduling cycle. The value TIMEOUT indicates that the node times out. You can specify multiple conditions for a custom alert rule. If you specify multiple condition, separate them with commas (,).
	//
	// example:
	//
	// FINISHED,ERROR
	RemindTypes *string `json:"RemindTypes,omitempty" xml:"RemindTypes,omitempty"`
	// The keyword in a rule name that is used to search for the rule. Fuzzy search is supported.
	//
	// example:
	//
	// Keyword
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListRemindsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemindsRequest) GoString() string {
	return s.String()
}

func (s *ListRemindsRequest) GetAlertTarget() *string {
	return s.AlertTarget
}

func (s *ListRemindsRequest) GetFounder() *string {
	return s.Founder
}

func (s *ListRemindsRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListRemindsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRemindsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRemindsRequest) GetRemindTypes() *string {
	return s.RemindTypes
}

func (s *ListRemindsRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListRemindsRequest) SetAlertTarget(v string) *ListRemindsRequest {
	s.AlertTarget = &v
	return s
}

func (s *ListRemindsRequest) SetFounder(v string) *ListRemindsRequest {
	s.Founder = &v
	return s
}

func (s *ListRemindsRequest) SetNodeId(v int64) *ListRemindsRequest {
	s.NodeId = &v
	return s
}

func (s *ListRemindsRequest) SetPageNumber(v int32) *ListRemindsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRemindsRequest) SetPageSize(v int32) *ListRemindsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRemindsRequest) SetRemindTypes(v string) *ListRemindsRequest {
	s.RemindTypes = &v
	return s
}

func (s *ListRemindsRequest) SetSearchText(v string) *ListRemindsRequest {
	s.SearchText = &v
	return s
}

func (s *ListRemindsRequest) Validate() error {
	return dara.Validate(s)
}
