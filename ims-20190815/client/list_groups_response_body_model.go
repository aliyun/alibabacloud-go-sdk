// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v *ListGroupsResponseBodyGroups) *ListGroupsResponseBody
	GetGroups() *ListGroupsResponseBodyGroups
	SetIsTruncated(v bool) *ListGroupsResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListGroupsResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
}

type ListGroupsResponseBody struct {
	// The information about the RAM user groups.
	Groups *ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 065527AA-2F2E-AD7C-7484-F2626CFE4934
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetGroups() *ListGroupsResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListGroupsResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) SetGroups(v *ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetIsTruncated(v bool) *ListGroupsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponseBody) SetMarker(v string) *ListGroupsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGroupsResponseBodyGroups struct {
	Group []*ListGroupsResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) GetGroup() []*ListGroupsResponseBodyGroupsGroup {
	return s.Group
}

func (s *ListGroupsResponseBodyGroups) SetGroup(v []*ListGroupsResponseBodyGroupsGroup) *ListGroupsResponseBodyGroups {
	s.Group = v
	return s
}

func (s *ListGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}

type ListGroupsResponseBodyGroupsGroup struct {
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
	// 2020-10-19T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
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
	// dev-team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-19T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListGroupsResponseBodyGroupsGroup) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroupsGroup) GetComments() *string {
	return s.Comments
}

func (s *ListGroupsResponseBodyGroupsGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListGroupsResponseBodyGroupsGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGroupsResponseBodyGroupsGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsResponseBodyGroupsGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyGroupsGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListGroupsResponseBodyGroupsGroup) SetComments(v string) *ListGroupsResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetCreateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.CreateDate = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetDisplayName(v string) *ListGroupsResponseBodyGroupsGroup {
	s.DisplayName = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetUpdateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.UpdateDate = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) Validate() error {
	return dara.Validate(s)
}
