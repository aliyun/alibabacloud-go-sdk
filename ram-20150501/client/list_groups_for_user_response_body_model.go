// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v *ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody
	GetGroups() *ListGroupsForUserResponseBodyGroups
	SetRequestId(v string) *ListGroupsForUserResponseBody
	GetRequestId() *string
}

type ListGroupsForUserResponseBody struct {
	Groups *ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DA772B52-BF9F-54CA-AC77-AA7A2DA89D46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) GetGroups() *ListGroupsForUserResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForUserResponseBody) SetGroups(v *ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetRequestId(v string) *ListGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForUserResponseBody) Validate() error {
	if s.Groups != nil {
		if err := s.Groups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupsForUserResponseBodyGroups struct {
	Group []*ListGroupsForUserResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsForUserResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroups) GetGroup() []*ListGroupsForUserResponseBodyGroupsGroup {
	return s.Group
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroup(v []*ListGroupsForUserResponseBodyGroupsGroup) *ListGroupsForUserResponseBodyGroups {
	s.Group = v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) Validate() error {
	if s.Group != nil {
		for _, item := range s.Group {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForUserResponseBodyGroupsGroup struct {
	Comments  *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	JoinDate  *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
}

func (s ListGroupsForUserResponseBodyGroupsGroup) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) GetComments() *string {
	return s.Comments
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) GetJoinDate() *string {
	return s.JoinDate
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetComments(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetJoinDate(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.JoinDate = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) Validate() error {
	return dara.Validate(s)
}
