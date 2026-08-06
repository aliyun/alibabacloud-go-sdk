// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetPromptRequest
	GetWorkspaceId() *string
}

type GetPromptRequest struct {
	// The workspace ID. To obtain the workspace ID, refer to [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 114243
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPromptRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPromptRequest) GoString() string {
	return s.String()
}

func (s *GetPromptRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPromptRequest) SetWorkspaceId(v string) *GetPromptRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPromptRequest) Validate() error {
	return dara.Validate(s)
}
