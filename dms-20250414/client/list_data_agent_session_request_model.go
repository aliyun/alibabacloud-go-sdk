// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateEndTime(v int64) *ListDataAgentSessionRequest
	GetCreateEndTime() *int64
	SetCreateStartTime(v int64) *ListDataAgentSessionRequest
	GetCreateStartTime() *int64
	SetCustomAgentId(v string) *ListDataAgentSessionRequest
	GetCustomAgentId() *string
	SetDMSUnit(v string) *ListDataAgentSessionRequest
	GetDMSUnit() *string
	SetIsSaved(v bool) *ListDataAgentSessionRequest
	GetIsSaved() *bool
	SetPageNumber(v int32) *ListDataAgentSessionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentSessionRequest
	GetPageSize() *int32
	SetQueryType(v string) *ListDataAgentSessionRequest
	GetQueryType() *string
	SetTitle(v string) *ListDataAgentSessionRequest
	GetTitle() *string
	SetWorkspaceId(v string) *ListDataAgentSessionRequest
	GetWorkspaceId() *string
}

type ListDataAgentSessionRequest struct {
	// example:
	//
	// 1770912000000
	CreateEndTime *int64 `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	// example:
	//
	// 1770825600000
	CreateStartTime *int64 `json:"CreateStartTime,omitempty" xml:"CreateStartTime,omitempty"`
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// true
	IsSaved *bool `json:"IsSaved,omitempty" xml:"IsSaved,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// myFavorite
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentSessionRequest) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListDataAgentSessionRequest) GetCreateStartTime() *int64 {
	return s.CreateStartTime
}

func (s *ListDataAgentSessionRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListDataAgentSessionRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ListDataAgentSessionRequest) GetIsSaved() *bool {
	return s.IsSaved
}

func (s *ListDataAgentSessionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentSessionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentSessionRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListDataAgentSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *ListDataAgentSessionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentSessionRequest) SetCreateEndTime(v int64) *ListDataAgentSessionRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetCreateStartTime(v int64) *ListDataAgentSessionRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetCustomAgentId(v string) *ListDataAgentSessionRequest {
	s.CustomAgentId = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetDMSUnit(v string) *ListDataAgentSessionRequest {
	s.DMSUnit = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetIsSaved(v bool) *ListDataAgentSessionRequest {
	s.IsSaved = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetPageNumber(v int32) *ListDataAgentSessionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetPageSize(v int32) *ListDataAgentSessionRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetQueryType(v string) *ListDataAgentSessionRequest {
	s.QueryType = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetTitle(v string) *ListDataAgentSessionRequest {
	s.Title = &v
	return s
}

func (s *ListDataAgentSessionRequest) SetWorkspaceId(v string) *ListDataAgentSessionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
