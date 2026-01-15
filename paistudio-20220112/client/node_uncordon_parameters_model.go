// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeUncordonParameters interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *NodeUncordonParameters
	GetQuotaId() *string
	SetWorkspaceId(v string) *NodeUncordonParameters
	GetWorkspaceId() *string
}

type NodeUncordonParameters struct {
	QuotaId     *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s NodeUncordonParameters) String() string {
	return dara.Prettify(s)
}

func (s NodeUncordonParameters) GoString() string {
	return s.String()
}

func (s *NodeUncordonParameters) GetQuotaId() *string {
	return s.QuotaId
}

func (s *NodeUncordonParameters) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *NodeUncordonParameters) SetQuotaId(v string) *NodeUncordonParameters {
	s.QuotaId = &v
	return s
}

func (s *NodeUncordonParameters) SetWorkspaceId(v string) *NodeUncordonParameters {
	s.WorkspaceId = &v
	return s
}

func (s *NodeUncordonParameters) Validate() error {
	return dara.Validate(s)
}
