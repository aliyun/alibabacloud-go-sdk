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
	// The information about the group.
	Group *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D3F0679E-9757-95DB-AF2D-04D5188C69C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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

func (s *CreateGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupResponseBodyGroup) SetComments(v string) *CreateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetCreateDate(v string) *CreateGroupResponseBodyGroup {
	s.CreateDate = &v
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

func (s *CreateGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
