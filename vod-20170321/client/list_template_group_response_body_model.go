// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTemplateGroupResponseBody
	GetRequestId() *string
	SetTemplateGroups(v *ListTemplateGroupResponseBodyTemplateGroups) *ListTemplateGroupResponseBody
	GetTemplateGroups() *ListTemplateGroupResponseBodyTemplateGroups
}

type ListTemplateGroupResponseBody struct {
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateGroups *ListTemplateGroupResponseBodyTemplateGroups `json:"TemplateGroups,omitempty" xml:"TemplateGroups,omitempty" type:"Struct"`
}

func (s ListTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateGroupResponseBody) GetTemplateGroups() *ListTemplateGroupResponseBodyTemplateGroups {
	return s.TemplateGroups
}

func (s *ListTemplateGroupResponseBody) SetRequestId(v string) *ListTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateGroupResponseBody) SetTemplateGroups(v *ListTemplateGroupResponseBodyTemplateGroups) *ListTemplateGroupResponseBody {
	s.TemplateGroups = v
	return s
}

func (s *ListTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplateGroupResponseBodyTemplateGroups struct {
	TemplateGroup []*ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup `json:"TemplateGroup,omitempty" xml:"TemplateGroup,omitempty" type:"Repeated"`
}

func (s ListTemplateGroupResponseBodyTemplateGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupResponseBodyTemplateGroups) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupResponseBodyTemplateGroups) GetTemplateGroup() []*ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	return s.TemplateGroup
}

func (s *ListTemplateGroupResponseBodyTemplateGroups) SetTemplateGroup(v []*ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) *ListTemplateGroupResponseBodyTemplateGroups {
	s.TemplateGroup = v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroups) Validate() error {
	return dara.Validate(s)
}

type ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup struct {
	DefaultGroup  *string `json:"DefaultGroup,omitempty" xml:"DefaultGroup,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupSymbol   *string `json:"GroupSymbol,omitempty" xml:"GroupSymbol,omitempty"`
	GroupType     *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IsLocked      *string `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TranscodeMode *string `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
}

func (s ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetDefaultGroup() *string {
	return s.DefaultGroup
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetGroupSymbol() *string {
	return s.GroupSymbol
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetIsLocked() *string {
	return s.IsLocked
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetName() *string {
	return s.Name
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetStatus() *string {
	return s.Status
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) GetTranscodeMode() *string {
	return s.TranscodeMode
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetDefaultGroup(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.DefaultGroup = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetGroupId(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.GroupId = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetGroupSymbol(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.GroupSymbol = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetGroupType(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.GroupType = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetIsLocked(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.IsLocked = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetName(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.Name = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetStatus(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.Status = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) SetTranscodeMode(v string) *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup {
	s.TranscodeMode = &v
	return s
}

func (s *ListTemplateGroupResponseBodyTemplateGroupsTemplateGroup) Validate() error {
	return dara.Validate(s)
}
