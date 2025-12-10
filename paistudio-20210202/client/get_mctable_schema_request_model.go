// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMCTableSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetMCTableSchemaRequest
	GetWorkspaceId() *string
}

type GetMCTableSchemaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMCTableSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMCTableSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMCTableSchemaRequest) SetWorkspaceId(v string) *GetMCTableSchemaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMCTableSchemaRequest) Validate() error {
	return dara.Validate(s)
}
