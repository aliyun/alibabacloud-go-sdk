// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *WorkspaceActionLogRequest
	GetKey() *string
	SetWorkspaceId(v string) *WorkspaceActionLogRequest
	GetWorkspaceId() *string
}

type WorkspaceActionLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceActionLogRequest) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionLogRequest) GoString() string {
	return s.String()
}

func (s *WorkspaceActionLogRequest) GetKey() *string {
	return s.Key
}

func (s *WorkspaceActionLogRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspaceActionLogRequest) SetKey(v string) *WorkspaceActionLogRequest {
	s.Key = &v
	return s
}

func (s *WorkspaceActionLogRequest) SetWorkspaceId(v string) *WorkspaceActionLogRequest {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspaceActionLogRequest) Validate() error {
	return dara.Validate(s)
}
