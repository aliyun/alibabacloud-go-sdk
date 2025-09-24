// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDefaultWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *UpdateDefaultWorkspaceRequest
	GetWorkspaceId() *string
}

type UpdateDefaultWorkspaceRequest struct {
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDefaultWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDefaultWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDefaultWorkspaceRequest) SetWorkspaceId(v string) *UpdateDefaultWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDefaultWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
