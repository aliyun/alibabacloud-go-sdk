// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *GetInstanceLogRequest
	GetLimit() *int64
	SetSkipLineNum(v int64) *GetInstanceLogRequest
	GetSkipLineNum() *int64
	SetWorkspaceId(v string) *GetInstanceLogRequest
	GetWorkspaceId() *string
}

type GetInstanceLogRequest struct {
	// example:
	//
	// 1
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// 1
	SkipLineNum *int64 `json:"skipLineNum,omitempty" xml:"skipLineNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// w-111
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceLogRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetInstanceLogRequest) GetSkipLineNum() *int64 {
	return s.SkipLineNum
}

func (s *GetInstanceLogRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetInstanceLogRequest) SetLimit(v int64) *GetInstanceLogRequest {
	s.Limit = &v
	return s
}

func (s *GetInstanceLogRequest) SetSkipLineNum(v int64) *GetInstanceLogRequest {
	s.SkipLineNum = &v
	return s
}

func (s *GetInstanceLogRequest) SetWorkspaceId(v string) *GetInstanceLogRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}
