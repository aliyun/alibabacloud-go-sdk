// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody
	GetGroup() *UpdateGroupResponseBodyGroup
	SetRequestId(v string) *UpdateGroupResponseBody
	GetRequestId() *string
}

type UpdateGroupResponseBody struct {
	// The information about the RAM user group.
	Group *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CDA656E3-3CE9-4A03-A8A3-B42A0C3C3287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) GetGroup() *UpdateGroupResponseBodyGroup {
	return s.Group
}

func (s *UpdateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupResponseBody) SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody {
	s.Group = v
	return s
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// Test-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-20T03:44:27Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyGroup) GetComments() *string {
	return s.Comments
}

func (s *UpdateGroupResponseBodyGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateGroupResponseBodyGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupResponseBodyGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateGroupResponseBodyGroup) SetComments(v string) *UpdateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetCreateDate(v string) *UpdateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetDisplayName(v string) *UpdateGroupResponseBodyGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupId(v string) *UpdateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupName(v string) *UpdateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetUpdateDate(v string) *UpdateGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
