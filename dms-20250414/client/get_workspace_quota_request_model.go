// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetWorkspaceQuotaRequest
	GetWorkspaceId() *string
}

type GetWorkspaceQuotaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceQuotaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceQuotaRequest) SetWorkspaceId(v string) *GetWorkspaceQuotaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
