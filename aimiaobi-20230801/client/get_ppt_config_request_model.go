// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *GetPptConfigRequest
	GetExternalUserId() *string
	SetWorkspaceId(v string) *GetPptConfigRequest
	GetWorkspaceId() *string
}

type GetPptConfigRequest struct {
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
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

func (s *GetPptConfigRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *GetPptConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptConfigRequest) SetExternalUserId(v string) *GetPptConfigRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetPptConfigRequest) SetWorkspaceId(v string) *GetPptConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptConfigRequest) Validate() error {
	return dara.Validate(s)
}
