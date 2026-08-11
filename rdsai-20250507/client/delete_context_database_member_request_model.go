// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *DeleteContextDatabaseMemberRequest
	GetMemberId() *string
	SetWorkspaceId(v string) *DeleteContextDatabaseMemberRequest
	GetWorkspaceId() *string
}

type DeleteContextDatabaseMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteContextDatabaseMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *DeleteContextDatabaseMemberRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteContextDatabaseMemberRequest) SetMemberId(v string) *DeleteContextDatabaseMemberRequest {
	s.MemberId = &v
	return s
}

func (s *DeleteContextDatabaseMemberRequest) SetWorkspaceId(v string) *DeleteContextDatabaseMemberRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteContextDatabaseMemberRequest) Validate() error {
	return dara.Validate(s)
}
