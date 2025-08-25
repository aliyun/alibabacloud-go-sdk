// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMpsSchedulerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMpsSchedulerListRequest
	GetAppId() *string
	SetEndTime(v int64) *QueryMpsSchedulerListRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *QueryMpsSchedulerListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryMpsSchedulerListRequest
	GetPageSize() *int32
	SetStartTime(v int64) *QueryMpsSchedulerListRequest
	GetStartTime() *int64
	SetType(v int32) *QueryMpsSchedulerListRequest
	GetType() *int32
	SetUniqueId(v string) *QueryMpsSchedulerListRequest
	GetUniqueId() *string
	SetWorkspaceId(v string) *QueryMpsSchedulerListRequest
	GetWorkspaceId() *string
}

type QueryMpsSchedulerListRequest struct {
	// This parameter is required.
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	UniqueId   *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMpsSchedulerListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListRequest) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMpsSchedulerListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMpsSchedulerListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryMpsSchedulerListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMpsSchedulerListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMpsSchedulerListRequest) GetType() *int32 {
	return s.Type
}

func (s *QueryMpsSchedulerListRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *QueryMpsSchedulerListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMpsSchedulerListRequest) SetAppId(v string) *QueryMpsSchedulerListRequest {
	s.AppId = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetEndTime(v int64) *QueryMpsSchedulerListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetPageNumber(v int32) *QueryMpsSchedulerListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetPageSize(v int32) *QueryMpsSchedulerListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetStartTime(v int64) *QueryMpsSchedulerListRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetType(v int32) *QueryMpsSchedulerListRequest {
	s.Type = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetUniqueId(v string) *QueryMpsSchedulerListRequest {
	s.UniqueId = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) SetWorkspaceId(v string) *QueryMpsSchedulerListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMpsSchedulerListRequest) Validate() error {
	return dara.Validate(s)
}
