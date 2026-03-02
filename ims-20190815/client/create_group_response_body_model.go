// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody
	GetGroup() *CreateGroupResponseBodyGroup
	SetRequestId(v string) *CreateGroupResponseBody
	GetRequestId() *string
}

type CreateGroupResponseBody struct {
	Group     *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetGroup() *CreateGroupResponseBodyGroup {
	return s.Group
}

func (s *CreateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupResponseBody) SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody {
	s.Group = v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupResponseBodyGroup struct {
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyGroup) GetComments() *string {
	return s.Comments
}

func (s *CreateGroupResponseBodyGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateGroupResponseBodyGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupResponseBodyGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateGroupResponseBodyGroup) SetComments(v string) *CreateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetCreateDate(v string) *CreateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetDisplayName(v string) *CreateGroupResponseBodyGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupId(v string) *CreateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupName(v string) *CreateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetUpdateDate(v string) *CreateGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
