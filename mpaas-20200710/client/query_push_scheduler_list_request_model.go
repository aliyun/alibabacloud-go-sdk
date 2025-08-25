// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushSchedulerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryPushSchedulerListRequest
	GetAppId() *string
	SetEndTime(v int64) *QueryPushSchedulerListRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *QueryPushSchedulerListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryPushSchedulerListRequest
	GetPageSize() *int32
	SetStartTime(v int64) *QueryPushSchedulerListRequest
	GetStartTime() *int64
	SetTenantId(v string) *QueryPushSchedulerListRequest
	GetTenantId() *string
	SetType(v int32) *QueryPushSchedulerListRequest
	GetType() *int32
	SetUniqueId(v string) *QueryPushSchedulerListRequest
	GetUniqueId() *string
	SetWorkspaceId(v string) *QueryPushSchedulerListRequest
	GetWorkspaceId() *string
}

type QueryPushSchedulerListRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	EndTime    *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	UniqueId  *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryPushSchedulerListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListRequest) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryPushSchedulerListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryPushSchedulerListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryPushSchedulerListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPushSchedulerListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryPushSchedulerListRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPushSchedulerListRequest) GetType() *int32 {
	return s.Type
}

func (s *QueryPushSchedulerListRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *QueryPushSchedulerListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryPushSchedulerListRequest) SetAppId(v string) *QueryPushSchedulerListRequest {
	s.AppId = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetEndTime(v int64) *QueryPushSchedulerListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetPageNumber(v int32) *QueryPushSchedulerListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetPageSize(v int32) *QueryPushSchedulerListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetStartTime(v int64) *QueryPushSchedulerListRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetTenantId(v string) *QueryPushSchedulerListRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetType(v int32) *QueryPushSchedulerListRequest {
	s.Type = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetUniqueId(v string) *QueryPushSchedulerListRequest {
	s.UniqueId = &v
	return s
}

func (s *QueryPushSchedulerListRequest) SetWorkspaceId(v string) *QueryPushSchedulerListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryPushSchedulerListRequest) Validate() error {
	return dara.Validate(s)
}
