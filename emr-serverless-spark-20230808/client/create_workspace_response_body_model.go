// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *CreateWorkspaceResponseBody
	GetOperationId() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *CreateWorkspaceResponseBody
	GetWorkspaceId() *string
}

type CreateWorkspaceResponseBody struct {
	// The operation ID.
	//
	// example:
	//
	// op-******
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBody) SetOperationId(v string) *CreateWorkspaceResponseBody {
	s.OperationId = &v
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
