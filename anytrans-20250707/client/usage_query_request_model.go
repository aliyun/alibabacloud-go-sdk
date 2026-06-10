// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *UsageQueryRequest
	GetEndTime() *int64
	SetStartTime(v int64) *UsageQueryRequest
	GetStartTime() *int64
	SetWorkspaceId(v string) *UsageQueryRequest
	GetWorkspaceId() *string
}

type UsageQueryRequest struct {
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UsageQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s UsageQueryRequest) GoString() string {
	return s.String()
}

func (s *UsageQueryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UsageQueryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UsageQueryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UsageQueryRequest) SetEndTime(v int64) *UsageQueryRequest {
	s.EndTime = &v
	return s
}

func (s *UsageQueryRequest) SetStartTime(v int64) *UsageQueryRequest {
	s.StartTime = &v
	return s
}

func (s *UsageQueryRequest) SetWorkspaceId(v string) *UsageQueryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UsageQueryRequest) Validate() error {
	return dara.Validate(s)
}
