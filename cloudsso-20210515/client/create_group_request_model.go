// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupRequest
	GetDirectoryId() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
}

type CreateGroupRequest struct {
	// The description of the group.
	//
	// The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// This is a group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the group.
	//
	// The name can contain letters, digits, underscores (_), hyphens (-), and periods (.).\\`\\`
	//
	// The name can be up to 128 characters in length.
	//
	// example:
	//
	// TestGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetDirectoryId(v string) *CreateGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
