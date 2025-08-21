// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComments(v string) *CreateGroupRequest
	GetComments() *string
	SetDisplayName(v string) *CreateGroupRequest
	GetDisplayName() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
}

type CreateGroupRequest struct {
	// The description.
	//
	// The description can be up to 128 characters in length.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user group.
	//
	// The name can be up to 24 characters in length.
	//
	// example:
	//
	// Dev-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the RAM user group. You must specify this parameter.
	//
	// The name can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetComments() *string {
	return s.Comments
}

func (s *CreateGroupRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) SetComments(v string) *CreateGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateGroupRequest) SetDisplayName(v string) *CreateGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
