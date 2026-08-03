// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentPattern(v string) *ListDataAgentMemoryRequest
	GetContentPattern() *string
	SetDMSUnit(v string) *ListDataAgentMemoryRequest
	GetDMSUnit() *string
	SetFromId(v string) *ListDataAgentMemoryRequest
	GetFromId() *string
	SetMemFrom(v string) *ListDataAgentMemoryRequest
	GetMemFrom() *string
	SetOrder(v string) *ListDataAgentMemoryRequest
	GetOrder() *string
	SetOrderBy(v string) *ListDataAgentMemoryRequest
	GetOrderBy() *string
	SetPageNum(v int64) *ListDataAgentMemoryRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListDataAgentMemoryRequest
	GetPageSize() *int64
	SetQueryAll(v bool) *ListDataAgentMemoryRequest
	GetQueryAll() *bool
}

type ListDataAgentMemoryRequest struct {
	// The content pattern used for fuzzy match search.
	//
	// example:
	//
	// user preference
	ContentPattern *string `json:"ContentPattern,omitempty" xml:"ContentPattern,omitempty"`
	// The current Data Management unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The source ID.
	//
	// - If MemFrom is set to session, FromId indicates the session ID.
	//
	// - If MemFrom is set to user, FromId indicates the RAM user ID.
	//
	// example:
	//
	// 8zm3**********g3yxa1
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// The memory source. Valid values:
	//
	// - session: Generated from a session.
	//
	// - user: Edited by a user.
	//
	// example:
	//
	// session
	MemFrom *string `json:"MemFrom,omitempty" xml:"MemFrom,omitempty"`
	// The sort order for the specified sort field. Default value: desc. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort field. Default value: hitTimes. Valid values:
	//
	// - hitTimes: The number of hits.
	//
	// - created: The creation time.
	//
	// example:
	//
	// hitTimes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The maximum number of entries per page. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query memories in all statuses. Default value: true.
	//
	// example:
	//
	// true
	QueryAll *bool `json:"QueryAll,omitempty" xml:"QueryAll,omitempty"`
}

func (s ListDataAgentMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMemoryRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentMemoryRequest) GetContentPattern() *string {
	return s.ContentPattern
}

func (s *ListDataAgentMemoryRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ListDataAgentMemoryRequest) GetFromId() *string {
	return s.FromId
}

func (s *ListDataAgentMemoryRequest) GetMemFrom() *string {
	return s.MemFrom
}

func (s *ListDataAgentMemoryRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataAgentMemoryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDataAgentMemoryRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListDataAgentMemoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentMemoryRequest) GetQueryAll() *bool {
	return s.QueryAll
}

func (s *ListDataAgentMemoryRequest) SetContentPattern(v string) *ListDataAgentMemoryRequest {
	s.ContentPattern = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetDMSUnit(v string) *ListDataAgentMemoryRequest {
	s.DMSUnit = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetFromId(v string) *ListDataAgentMemoryRequest {
	s.FromId = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetMemFrom(v string) *ListDataAgentMemoryRequest {
	s.MemFrom = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetOrder(v string) *ListDataAgentMemoryRequest {
	s.Order = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetOrderBy(v string) *ListDataAgentMemoryRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetPageNum(v int64) *ListDataAgentMemoryRequest {
	s.PageNum = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetPageSize(v int64) *ListDataAgentMemoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentMemoryRequest) SetQueryAll(v bool) *ListDataAgentMemoryRequest {
	s.QueryAll = &v
	return s
}

func (s *ListDataAgentMemoryRequest) Validate() error {
	return dara.Validate(s)
}
