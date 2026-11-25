// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetContextDatabaseUsageRequest
	GetEndTime() *string
	SetStartTime(v string) *GetContextDatabaseUsageRequest
	GetStartTime() *string
	SetWorkspaceId(v string) *GetContextDatabaseUsageRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseUsageRequest struct {
	// The end time of the query range, exclusive. It must be later than `StartTime`. Use the ISO 8601 time format, such as `2026-08-01T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-01T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the query range, inclusive. Use the ISO 8601 time format, such as `2026-07-01T00:00:00Z` (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the workspace to query. If specified, only the data of this workspace under the caller account is queried. Locked workspaces can also be queried. If omitted, all workspaces under the caller account are queried.
	//
	// example:
	//
	// ws-x9z4ctzed9gfu2w
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseUsageRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetContextDatabaseUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextDatabaseUsageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseUsageRequest) SetEndTime(v string) *GetContextDatabaseUsageRequest {
	s.EndTime = &v
	return s
}

func (s *GetContextDatabaseUsageRequest) SetStartTime(v string) *GetContextDatabaseUsageRequest {
	s.StartTime = &v
	return s
}

func (s *GetContextDatabaseUsageRequest) SetWorkspaceId(v string) *GetContextDatabaseUsageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseUsageRequest) Validate() error {
	return dara.Validate(s)
}
