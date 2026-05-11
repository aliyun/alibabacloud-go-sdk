// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTimeLeftRange(v int64) *ListConversationsRequest
	GetBeginTimeLeftRange() *int64
	SetBeginTimeRightRange(v int64) *ListConversationsRequest
	GetBeginTimeRightRange() *int64
	SetCallingNumber(v string) *ListConversationsRequest
	GetCallingNumber() *string
	SetDebugConversation(v int32) *ListConversationsRequest
	GetDebugConversation() *int32
	SetInstanceId(v string) *ListConversationsRequest
	GetInstanceId() *string
	SetIsSandBox(v string) *ListConversationsRequest
	GetIsSandBox() *string
	SetPageNumber(v int32) *ListConversationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConversationsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListConversationsRequest
	GetQuery() *string
	SetResult(v int64) *ListConversationsRequest
	GetResult() *int64
	SetRoundsLeftRange(v int32) *ListConversationsRequest
	GetRoundsLeftRange() *int32
	SetRoundsRightRange(v int32) *ListConversationsRequest
	GetRoundsRightRange() *int32
}

type ListConversationsRequest struct {
	// example:
	//
	// 1638288000000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1637547875311
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// example:
	//
	// 138106*****
	CallingNumber     *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	DebugConversation *int32  `json:"DebugConversation,omitempty" xml:"DebugConversation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IsSandBox *string `json:"IsSandBox,omitempty" xml:"IsSandBox,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 13788914724
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 0
	Result           *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	RoundsLeftRange  *int32 `json:"RoundsLeftRange,omitempty" xml:"RoundsLeftRange,omitempty"`
	RoundsRightRange *int32 `json:"RoundsRightRange,omitempty" xml:"RoundsRightRange,omitempty"`
}

func (s ListConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConversationsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationsRequest) GetBeginTimeLeftRange() *int64 {
	return s.BeginTimeLeftRange
}

func (s *ListConversationsRequest) GetBeginTimeRightRange() *int64 {
	return s.BeginTimeRightRange
}

func (s *ListConversationsRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListConversationsRequest) GetDebugConversation() *int32 {
	return s.DebugConversation
}

func (s *ListConversationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConversationsRequest) GetIsSandBox() *string {
	return s.IsSandBox
}

func (s *ListConversationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConversationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConversationsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListConversationsRequest) GetResult() *int64 {
	return s.Result
}

func (s *ListConversationsRequest) GetRoundsLeftRange() *int32 {
	return s.RoundsLeftRange
}

func (s *ListConversationsRequest) GetRoundsRightRange() *int32 {
	return s.RoundsRightRange
}

func (s *ListConversationsRequest) SetBeginTimeLeftRange(v int64) *ListConversationsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *ListConversationsRequest) SetBeginTimeRightRange(v int64) *ListConversationsRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *ListConversationsRequest) SetCallingNumber(v string) *ListConversationsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListConversationsRequest) SetDebugConversation(v int32) *ListConversationsRequest {
	s.DebugConversation = &v
	return s
}

func (s *ListConversationsRequest) SetInstanceId(v string) *ListConversationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConversationsRequest) SetIsSandBox(v string) *ListConversationsRequest {
	s.IsSandBox = &v
	return s
}

func (s *ListConversationsRequest) SetPageNumber(v int32) *ListConversationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConversationsRequest) SetPageSize(v int32) *ListConversationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConversationsRequest) SetQuery(v string) *ListConversationsRequest {
	s.Query = &v
	return s
}

func (s *ListConversationsRequest) SetResult(v int64) *ListConversationsRequest {
	s.Result = &v
	return s
}

func (s *ListConversationsRequest) SetRoundsLeftRange(v int32) *ListConversationsRequest {
	s.RoundsLeftRange = &v
	return s
}

func (s *ListConversationsRequest) SetRoundsRightRange(v int32) *ListConversationsRequest {
	s.RoundsRightRange = &v
	return s
}

func (s *ListConversationsRequest) Validate() error {
	return dara.Validate(s)
}
