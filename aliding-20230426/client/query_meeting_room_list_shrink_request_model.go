// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryMeetingRoomListShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v int64) *QueryMeetingRoomListShrinkRequest
	GetNextToken() *int64
	SetTenantContextShrink(v string) *QueryMeetingRoomListShrinkRequest
	GetTenantContextShrink() *string
}

type QueryMeetingRoomListShrinkRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 124L
	NextToken           *int64  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryMeetingRoomListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMeetingRoomListShrinkRequest) GetNextToken() *int64 {
	return s.NextToken
}

func (s *QueryMeetingRoomListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMeetingRoomListShrinkRequest) SetMaxResults(v int32) *QueryMeetingRoomListShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMeetingRoomListShrinkRequest) SetNextToken(v int64) *QueryMeetingRoomListShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomListShrinkRequest) SetTenantContextShrink(v string) *QueryMeetingRoomListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMeetingRoomListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
