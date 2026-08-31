// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallEndTimeBegin(v int64) *ReadOutboundTaskCallListShrinkRequest
	GetCallEndTimeBegin() *int64
	SetCallEndTimeEnd(v int64) *ReadOutboundTaskCallListShrinkRequest
	GetCallEndTimeEnd() *int64
	SetCallStartTimeBegin(v int64) *ReadOutboundTaskCallListShrinkRequest
	GetCallStartTimeBegin() *int64
	SetCallStartTimeEnd(v int64) *ReadOutboundTaskCallListShrinkRequest
	GetCallStartTimeEnd() *int64
	SetCallerUacAccountId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetCallerUacAccountId() *string
	SetCurrent(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetCurrent() *int32
	SetCurrentWorkspaceId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetCurrentWorkspaceId() *string
	SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListShrinkRequest
	GetCustomerNameOrPhone() *string
	SetDisplayStatusListShrink(v string) *ReadOutboundTaskCallListShrinkRequest
	GetDisplayStatusListShrink() *string
	SetDurationRangeListShrink(v string) *ReadOutboundTaskCallListShrinkRequest
	GetDurationRangeListShrink() *string
	SetLabelTagsShrink(v string) *ReadOutboundTaskCallListShrinkRequest
	GetLabelTagsShrink() *string
	SetMaxResults(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadOutboundTaskCallListShrinkRequest
	GetNextToken() *string
	SetSize(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetSize() *int32
	SetTaskId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetTaskId() *string
	SetUserId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetUserId() *string
}

type ReadOutboundTaskCallListShrinkRequest struct {
	// example:
	//
	// 1735689600000
	CallEndTimeBegin *int64 `json:"CallEndTimeBegin,omitempty" xml:"CallEndTimeBegin,omitempty"`
	// example:
	//
	// 1767225600000
	CallEndTimeEnd *int64 `json:"CallEndTimeEnd,omitempty" xml:"CallEndTimeEnd,omitempty"`
	// example:
	//
	// 1735689600000
	CallStartTimeBegin *int64 `json:"CallStartTimeBegin,omitempty" xml:"CallStartTimeBegin,omitempty"`
	// example:
	//
	// 1767225600000
	CallStartTimeEnd *int64 `json:"CallStartTimeEnd,omitempty" xml:"CallStartTimeEnd,omitempty"`
	// example:
	//
	// abc123***
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// abc123***
	CurrentWorkspaceId *string `json:"CurrentWorkspaceId,omitempty" xml:"CurrentWorkspaceId,omitempty"`
	// example:
	//
	// 张先生
	CustomerNameOrPhone *string `json:"CustomerNameOrPhone,omitempty" xml:"CustomerNameOrPhone,omitempty"`
	// example:
	//
	// ["1", "2"]
	DisplayStatusListShrink *string `json:"DisplayStatusList,omitempty" xml:"DisplayStatusList,omitempty"`
	// example:
	//
	// ["WITHIN_15_SECONDS", "FROM_1_TO_3_MINUTES"]
	DurationRangeListShrink *string `json:"DurationRangeList,omitempty" xml:"DurationRangeList,omitempty"`
	// example:
	//
	// ["有意向", "高净值"]
	LabelTagsShrink *string `json:"LabelTags,omitempty" xml:"LabelTags,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 51CC272E-D879-1B23-B98E-FCFB072D362B
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReadOutboundTaskCallListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCallEndTimeBegin() *int64 {
	return s.CallEndTimeBegin
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCallEndTimeEnd() *int64 {
	return s.CallEndTimeEnd
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCallStartTimeBegin() *int64 {
	return s.CallStartTimeBegin
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCallStartTimeEnd() *int64 {
	return s.CallStartTimeEnd
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCurrentWorkspaceId() *string {
	return s.CurrentWorkspaceId
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCustomerNameOrPhone() *string {
	return s.CustomerNameOrPhone
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetDisplayStatusListShrink() *string {
	return s.DisplayStatusListShrink
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetDurationRangeListShrink() *string {
	return s.DurationRangeListShrink
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetLabelTagsShrink() *string {
	return s.LabelTagsShrink
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCallEndTimeBegin(v int64) *ReadOutboundTaskCallListShrinkRequest {
	s.CallEndTimeBegin = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCallEndTimeEnd(v int64) *ReadOutboundTaskCallListShrinkRequest {
	s.CallEndTimeEnd = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCallStartTimeBegin(v int64) *ReadOutboundTaskCallListShrinkRequest {
	s.CallStartTimeBegin = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCallStartTimeEnd(v int64) *ReadOutboundTaskCallListShrinkRequest {
	s.CallStartTimeEnd = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCallerUacAccountId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCurrent(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.Current = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCurrentWorkspaceId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.CurrentWorkspaceId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.CustomerNameOrPhone = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetDisplayStatusListShrink(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.DisplayStatusListShrink = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetDurationRangeListShrink(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.DurationRangeListShrink = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetLabelTagsShrink(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.LabelTagsShrink = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetMaxResults(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetNextToken(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetSize(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.Size = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetTaskId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetUserId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
