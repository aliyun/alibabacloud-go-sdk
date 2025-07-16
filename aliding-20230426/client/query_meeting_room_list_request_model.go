// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryMeetingRoomListRequest
	GetMaxResults() *int32
	SetNextToken(v int64) *QueryMeetingRoomListRequest
	GetNextToken() *int64
	SetTenantContext(v *QueryMeetingRoomListRequestTenantContext) *QueryMeetingRoomListRequest
	GetTenantContext() *QueryMeetingRoomListRequestTenantContext
}

type QueryMeetingRoomListRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 124L
	NextToken     *int64                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContext *QueryMeetingRoomListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMeetingRoomListRequest) GetNextToken() *int64 {
	return s.NextToken
}

func (s *QueryMeetingRoomListRequest) GetTenantContext() *QueryMeetingRoomListRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMeetingRoomListRequest) SetMaxResults(v int32) *QueryMeetingRoomListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMeetingRoomListRequest) SetNextToken(v int64) *QueryMeetingRoomListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomListRequest) SetTenantContext(v *QueryMeetingRoomListRequestTenantContext) *QueryMeetingRoomListRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMeetingRoomListRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMeetingRoomListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMeetingRoomListRequestTenantContext) SetTenantId(v string) *QueryMeetingRoomListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMeetingRoomListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
