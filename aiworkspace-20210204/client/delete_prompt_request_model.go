// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeletePromptRequest
	GetWorkspaceId() *string
}

type DeletePromptRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeletePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptRequest) GoString() string {
	return s.String()
}

func (s *DeletePromptRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeletePromptRequest) SetWorkspaceId(v string) *DeletePromptRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeletePromptRequest) Validate() error {
	return dara.Validate(s)
}
