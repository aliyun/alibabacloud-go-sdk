// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceIdName interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *WorkspaceIdName
	GetWorkspaceId() *string
}

type WorkspaceIdName struct {
	// example:
	//
	// ws123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceIdName) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceIdName) GoString() string {
	return s.String()
}

func (s *WorkspaceIdName) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspaceIdName) SetWorkspaceId(v string) *WorkspaceIdName {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspaceIdName) Validate() error {
	return dara.Validate(s)
}
