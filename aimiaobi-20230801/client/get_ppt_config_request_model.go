// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetPptConfigRequest
	GetWorkspaceId() *string
}

type GetPptConfigRequest struct {
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPptConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPptConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPptConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptConfigRequest) SetWorkspaceId(v string) *GetPptConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptConfigRequest) Validate() error {
	return dara.Validate(s)
}
