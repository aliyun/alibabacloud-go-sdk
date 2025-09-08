// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFactAuditUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetFactAuditUrlRequest
	GetWorkspaceId() *string
}

type GetFactAuditUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetFactAuditUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFactAuditUrlRequest) GoString() string {
	return s.String()
}

func (s *GetFactAuditUrlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetFactAuditUrlRequest) SetWorkspaceId(v string) *GetFactAuditUrlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetFactAuditUrlRequest) Validate() error {
	return dara.Validate(s)
}
