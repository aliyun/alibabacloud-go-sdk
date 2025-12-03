// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupExternalId(v string) *CreateGroupRequest
	GetGroupExternalId() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
}

type CreateGroupRequest struct {
	// The external ID.
	//
	// example:
	//
	// group_2bo6lefcewdausyyxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// The organization name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name001
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) SetGroupExternalId(v string) *CreateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
