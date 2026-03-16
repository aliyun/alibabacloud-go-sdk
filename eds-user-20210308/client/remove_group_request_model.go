// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *RemoveGroupRequest
	GetBusinessChannel() *string
	SetGroupId(v string) *RemoveGroupRequest
	GetGroupId() *string
	SetGroupIds(v []*string) *RemoveGroupRequest
	GetGroupIds() []*string
}

type RemoveGroupRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The ID of the user group to be deleted.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The IDs of the user groups to be deleted.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s RemoveGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *RemoveGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveGroupRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *RemoveGroupRequest) SetBusinessChannel(v string) *RemoveGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *RemoveGroupRequest) SetGroupId(v string) *RemoveGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveGroupRequest) SetGroupIds(v []*string) *RemoveGroupRequest {
	s.GroupIds = v
	return s
}

func (s *RemoveGroupRequest) Validate() error {
	return dara.Validate(s)
}
