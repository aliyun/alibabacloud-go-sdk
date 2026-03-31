// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody
	GetGroup() *GetGroupResponseBodyGroup
	SetRequestId(v string) *GetGroupResponseBody
	GetRequestId() *string
}

type GetGroupResponseBody struct {
	// The information about the RAM user group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4065824-E422-3ED6-68B1-1AF7D5C7804C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) GetGroup() *GetGroupResponseBodyGroup {
	return s.Group
}

func (s *GetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGroupResponseBodyGroup struct {
	// The description of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user group was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// g-FpMEHiMysofp****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2015-02-11T03:15:21Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) GetComments() *string {
	return s.Comments
}

func (s *GetGroupResponseBodyGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupResponseBodyGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetGroupResponseBodyGroup) SetComments(v string) *GetGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetCreateDate(v string) *GetGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupName(v string) *GetGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateDate(v string) *GetGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

func (s *GetGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
