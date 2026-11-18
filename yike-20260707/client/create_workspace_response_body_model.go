// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateWorkspaceResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *CreateWorkspaceResponseBody
	GetWorkspaceId() *string
}

type CreateWorkspaceResponseBody struct {
	// The ID of the default project created through automatic creation when the studio is created.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the studio.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBody) SetProjectId(v string) *CreateWorkspaceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
