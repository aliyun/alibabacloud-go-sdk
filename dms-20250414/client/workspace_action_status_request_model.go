// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *WorkspaceActionStatusRequest
	GetKey() *string
	SetWorkspaceId(v string) *WorkspaceActionStatusRequest
	GetWorkspaceId() *string
}

type WorkspaceActionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ws-xxx-xxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceActionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionStatusRequest) GoString() string {
	return s.String()
}

func (s *WorkspaceActionStatusRequest) GetKey() *string {
	return s.Key
}

func (s *WorkspaceActionStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspaceActionStatusRequest) SetKey(v string) *WorkspaceActionStatusRequest {
	s.Key = &v
	return s
}

func (s *WorkspaceActionStatusRequest) SetWorkspaceId(v string) *WorkspaceActionStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspaceActionStatusRequest) Validate() error {
	return dara.Validate(s)
}
