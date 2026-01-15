// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeCordonParameters interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *NodeCordonParameters
	GetComment() *string
	SetQuotaId(v string) *NodeCordonParameters
	GetQuotaId() *string
	SetWorkspaceId(v string) *NodeCordonParameters
	GetWorkspaceId() *string
}

type NodeCordonParameters struct {
	Comment     *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	QuotaId     *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s NodeCordonParameters) String() string {
	return dara.Prettify(s)
}

func (s NodeCordonParameters) GoString() string {
	return s.String()
}

func (s *NodeCordonParameters) GetComment() *string {
	return s.Comment
}

func (s *NodeCordonParameters) GetQuotaId() *string {
	return s.QuotaId
}

func (s *NodeCordonParameters) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *NodeCordonParameters) SetComment(v string) *NodeCordonParameters {
	s.Comment = &v
	return s
}

func (s *NodeCordonParameters) SetQuotaId(v string) *NodeCordonParameters {
	s.QuotaId = &v
	return s
}

func (s *NodeCordonParameters) SetWorkspaceId(v string) *NodeCordonParameters {
	s.WorkspaceId = &v
	return s
}

func (s *NodeCordonParameters) Validate() error {
	return dara.Validate(s)
}
