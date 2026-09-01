// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSkillMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDataAgentSkillMetaRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentSkillMetaRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListDataAgentSkillMetaRequest
	GetSearchKey() *string
	SetSkillFrom(v string) *ListDataAgentSkillMetaRequest
	GetSkillFrom() *string
	SetSkillId(v string) *ListDataAgentSkillMetaRequest
	GetSkillId() *string
	SetSkillName(v string) *ListDataAgentSkillMetaRequest
	GetSkillName() *string
	SetWorkspaceId(v string) *ListDataAgentSkillMetaRequest
	GetWorkspaceId() *string
}

type ListDataAgentSkillMetaRequest struct {
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword for fuzzy match.
	//
	// example:
	//
	// data-query-skill
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The source of the skill. Valid values:
	//
	// - User: a skill uploaded by the user.
	//
	// - Agent: a skill derived from Agent analysis.
	//
	// example:
	//
	// User
	SkillFrom *string `json:"SkillFrom,omitempty" xml:"SkillFrom,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// data-query-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// b5u96hud*************gq3
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentSkillMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSkillMetaRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentSkillMetaRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentSkillMetaRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentSkillMetaRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataAgentSkillMetaRequest) GetSkillFrom() *string {
	return s.SkillFrom
}

func (s *ListDataAgentSkillMetaRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *ListDataAgentSkillMetaRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ListDataAgentSkillMetaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentSkillMetaRequest) SetPageNumber(v int32) *ListDataAgentSkillMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetPageSize(v int32) *ListDataAgentSkillMetaRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetSearchKey(v string) *ListDataAgentSkillMetaRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetSkillFrom(v string) *ListDataAgentSkillMetaRequest {
	s.SkillFrom = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetSkillId(v string) *ListDataAgentSkillMetaRequest {
	s.SkillId = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetSkillName(v string) *ListDataAgentSkillMetaRequest {
	s.SkillName = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) SetWorkspaceId(v string) *ListDataAgentSkillMetaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentSkillMetaRequest) Validate() error {
	return dara.Validate(s)
}
