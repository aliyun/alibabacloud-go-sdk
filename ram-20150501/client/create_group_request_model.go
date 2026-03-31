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
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
}

type CreateGroupRequest struct {
	// The description.
	//
	// The value can be up to 128 characters in length.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The name of the user group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
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

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) SetComments(v string) *CreateGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
