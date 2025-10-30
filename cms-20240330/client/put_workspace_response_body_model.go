// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspaceName(v string) *PutWorkspaceResponseBody
	GetWorkspaceName() *string
}

type PutWorkspaceResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Name of the workspace.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s PutWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *PutWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *PutWorkspaceResponseBody) SetRequestId(v string) *PutWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutWorkspaceResponseBody) SetWorkspaceName(v string) *PutWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *PutWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
