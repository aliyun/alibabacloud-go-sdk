// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillReferencesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSkillReferencesResponseBodyData) *ListSkillReferencesResponseBody
	GetData() *ListSkillReferencesResponseBodyData
	SetRequestId(v string) *ListSkillReferencesResponseBody
	GetRequestId() *string
}

type ListSkillReferencesResponseBody struct {
	// The skill reference relationship data returned by the paged query. The data is returned with paging.
	Data *ListSkillReferencesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID, which is used for troubleshooting.
	//
	// example:
	//
	// 5C6D9E10-1234-5678-9ABC-DEF012345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSkillReferencesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillReferencesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillReferencesResponseBody) GetData() *ListSkillReferencesResponseBodyData {
	return s.Data
}

func (s *ListSkillReferencesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillReferencesResponseBody) SetData(v *ListSkillReferencesResponseBodyData) *ListSkillReferencesResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillReferencesResponseBody) SetRequestId(v string) *ListSkillReferencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillReferencesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillReferencesResponseBodyData struct {
	// The list of skill reference relationships on the current page.
	PageItems []*ListSkillReferencesResponseBodyDataPageItems `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The total number of available pages.
	//
	// example:
	//
	// 5
	PagesAvailable *int32 `json:"pagesAvailable,omitempty" xml:"pagesAvailable,omitempty"`
	// The total number of reference relationships that match the filter conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSkillReferencesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillReferencesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillReferencesResponseBodyData) GetPageItems() []*ListSkillReferencesResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListSkillReferencesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillReferencesResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListSkillReferencesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillReferencesResponseBodyData) SetPageItems(v []*ListSkillReferencesResponseBodyDataPageItems) *ListSkillReferencesResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListSkillReferencesResponseBodyData) SetPageNumber(v int32) *ListSkillReferencesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillReferencesResponseBodyData) SetPagesAvailable(v int32) *ListSkillReferencesResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListSkillReferencesResponseBodyData) SetTotalCount(v int32) *ListSkillReferencesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillReferencesResponseBodyData) Validate() error {
	if s.PageItems != nil {
		for _, item := range s.PageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillReferencesResponseBodyDataPageItems struct {
	// The ID of the referencing entity (the identifier of the Agent or AgentSpec).
	//
	// example:
	//
	// agent-1234567890abcdef
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The type of the referencing entity. Valid values: AGENT and AGENTSPEC.
	//
	// example:
	//
	// AGENT
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The version of the referencing entity.
	//
	// example:
	//
	// 1.0.0
	OwnerVersion *string `json:"ownerVersion,omitempty" xml:"ownerVersion,omitempty"`
	// The reference selector type. Valid values: LABEL and VERSION.
	//
	// example:
	//
	// LABEL
	SelectorType *string `json:"selectorType,omitempty" xml:"selectorType,omitempty"`
	// The reference selector value, such as latest, a named label, HEAD, or a specific version.
	//
	// example:
	//
	// HEAD
	SelectorValue *string `json:"selectorValue,omitempty" xml:"selectorValue,omitempty"`
	// The name of the referenced skill.
	//
	// example:
	//
	// web-search
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The workspace ID to which the reference belongs.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListSkillReferencesResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListSkillReferencesResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetOwnerVersion() *string {
	return s.OwnerVersion
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetSelectorType() *string {
	return s.SelectorType
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetSelectorValue() *string {
	return s.SelectorValue
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetSkillName() *string {
	return s.SkillName
}

func (s *ListSkillReferencesResponseBodyDataPageItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetOwnerId(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.OwnerId = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetOwnerType(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.OwnerType = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetOwnerVersion(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.OwnerVersion = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetSelectorType(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.SelectorType = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetSelectorValue(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.SelectorValue = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetSkillName(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.SkillName = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) SetWorkspaceId(v string) *ListSkillReferencesResponseBodyDataPageItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListSkillReferencesResponseBodyDataPageItems) Validate() error {
	return dara.Validate(s)
}
