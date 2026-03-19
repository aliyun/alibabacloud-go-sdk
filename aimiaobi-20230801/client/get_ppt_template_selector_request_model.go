// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptTemplateSelectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetPptTemplateSelectorRequest
	GetWorkspaceId() *string
}

type GetPptTemplateSelectorRequest struct {
	// example:
	//
	// lm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPptTemplateSelectorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPptTemplateSelectorRequest) GoString() string {
	return s.String()
}

func (s *GetPptTemplateSelectorRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptTemplateSelectorRequest) SetWorkspaceId(v string) *GetPptTemplateSelectorRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptTemplateSelectorRequest) Validate() error {
	return dara.Validate(s)
}
