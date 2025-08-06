// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTemplateGroupConsoleResponseBody
	GetRequestId() *string
	SetTemplateGroups(v *ListTemplateGroupConsoleResponseBodyTemplateGroups) *ListTemplateGroupConsoleResponseBody
	GetTemplateGroups() *ListTemplateGroupConsoleResponseBodyTemplateGroups
}

type ListTemplateGroupConsoleResponseBody struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateGroups *ListTemplateGroupConsoleResponseBodyTemplateGroups `json:"TemplateGroups,omitempty" xml:"TemplateGroups,omitempty" type:"Struct"`
}

func (s ListTemplateGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateGroupConsoleResponseBody) GetTemplateGroups() *ListTemplateGroupConsoleResponseBodyTemplateGroups {
	return s.TemplateGroups
}

func (s *ListTemplateGroupConsoleResponseBody) SetRequestId(v string) *ListTemplateGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBody) SetTemplateGroups(v *ListTemplateGroupConsoleResponseBodyTemplateGroups) *ListTemplateGroupConsoleResponseBody {
	s.TemplateGroups = v
	return s
}

func (s *ListTemplateGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplateGroupConsoleResponseBodyTemplateGroups struct {
	TemplateGroup []*ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup `json:"TemplateGroup,omitempty" xml:"TemplateGroup,omitempty" type:"Repeated"`
}

func (s ListTemplateGroupConsoleResponseBodyTemplateGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupConsoleResponseBodyTemplateGroups) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroups) GetTemplateGroup() []*ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	return s.TemplateGroup
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroups) SetTemplateGroup(v []*ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) *ListTemplateGroupConsoleResponseBodyTemplateGroups {
	s.TemplateGroup = v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroups) Validate() error {
	return dara.Validate(s)
}

type ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup struct {
	DefaultGroup  *string `json:"DefaultGroup,omitempty" xml:"DefaultGroup,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupSymbol   *string `json:"GroupSymbol,omitempty" xml:"GroupSymbol,omitempty"`
	GroupType     *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IsLocked      *string `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TranscodeMode *string `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
}

func (s ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetDefaultGroup() *string {
	return s.DefaultGroup
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetGroupSymbol() *string {
	return s.GroupSymbol
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetIsLocked() *string {
	return s.IsLocked
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetName() *string {
	return s.Name
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetStatus() *string {
	return s.Status
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) GetTranscodeMode() *string {
	return s.TranscodeMode
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetDefaultGroup(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.DefaultGroup = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetGroupId(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.GroupId = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetGroupSymbol(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.GroupSymbol = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetGroupType(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.GroupType = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetIsLocked(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.IsLocked = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetName(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.Name = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetStatus(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.Status = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) SetTranscodeMode(v string) *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup {
	s.TranscodeMode = &v
	return s
}

func (s *ListTemplateGroupConsoleResponseBodyTemplateGroupsTemplateGroup) Validate() error {
	return dara.Validate(s)
}
