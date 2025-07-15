// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCategoriesByTaskIdRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetCategoriesByTaskIdRequest
	GetWorkspaceId() *string
}

type GetCategoriesByTaskIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCategoriesByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCategoriesByTaskIdRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCategoriesByTaskIdRequest) SetTaskId(v string) *GetCategoriesByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetCategoriesByTaskIdRequest) SetWorkspaceId(v string) *GetCategoriesByTaskIdRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetCategoriesByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
