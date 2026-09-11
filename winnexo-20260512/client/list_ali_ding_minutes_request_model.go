// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *ListAliDingMinutesRequest
	GetCursor() *string
	SetEndTime(v string) *ListAliDingMinutesRequest
	GetEndTime() *string
	SetPageSize(v int32) *ListAliDingMinutesRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListAliDingMinutesRequest
	GetStartTime() *string
	SetTenantId(v string) *ListAliDingMinutesRequest
	GetTenantId() *string
}

type ListAliDingMinutesRequest struct {
	// The cursor for the paged query. Set this parameter to 0 for the first request. For subsequent requests, set this parameter to the **nextCursor*	- value returned in the previous response. For more information about paging, see the response parameters.
	//
	// example:
	//
	// opaque-next-cursor
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The actual end timestamp of the live session, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T23:59:59+08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query start time. This value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-01T00:00:00+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through the winnexo-cli --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAliDingMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingMinutesRequest) GoString() string {
	return s.String()
}

func (s *ListAliDingMinutesRequest) GetCursor() *string {
	return s.Cursor
}

func (s *ListAliDingMinutesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAliDingMinutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliDingMinutesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAliDingMinutesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAliDingMinutesRequest) SetCursor(v string) *ListAliDingMinutesRequest {
	s.Cursor = &v
	return s
}

func (s *ListAliDingMinutesRequest) SetEndTime(v string) *ListAliDingMinutesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAliDingMinutesRequest) SetPageSize(v int32) *ListAliDingMinutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAliDingMinutesRequest) SetStartTime(v string) *ListAliDingMinutesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAliDingMinutesRequest) SetTenantId(v string) *ListAliDingMinutesRequest {
	s.TenantId = &v
	return s
}

func (s *ListAliDingMinutesRequest) Validate() error {
	return dara.Validate(s)
}
