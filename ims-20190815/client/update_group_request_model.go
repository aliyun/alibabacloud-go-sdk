// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *UpdateGroupRequest
	GetGroupName() *string
	SetNewComments(v string) *UpdateGroupRequest
	GetNewComments() *string
	SetNewDisplayName(v string) *UpdateGroupRequest
	GetNewDisplayName() *string
	SetNewGroupName(v string) *UpdateGroupRequest
	GetNewGroupName() *string
}

type UpdateGroupRequest struct {
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NewComments    *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewGroupName   *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupRequest) GetNewComments() *string {
	return s.NewComments
}

func (s *UpdateGroupRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateGroupRequest) GetNewGroupName() *string {
	return s.NewGroupName
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewComments(v string) *UpdateGroupRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateGroupRequest) SetNewDisplayName(v string) *UpdateGroupRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	return dara.Validate(s)
}
