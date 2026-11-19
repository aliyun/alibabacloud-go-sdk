// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillReferencesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSkillReferencesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSkillReferencesRequest
	GetPageSize() *int32
	SetRefType(v string) *ListSkillReferencesRequest
	GetRefType() *string
	SetSelectorType(v string) *ListSkillReferencesRequest
	GetSelectorType() *string
	SetSelectorValue(v string) *ListSkillReferencesRequest
	GetSelectorValue() *string
	SetSourceId(v string) *ListSkillReferencesRequest
	GetSourceId() *string
}

type ListSkillReferencesRequest struct {
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// The number of entries per page. If not specified, the server default value is used.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters by reference strength. Valid values: STRONG, WEAK.
	//
	// example:
	//
	// STRONG
	RefType *string `json:"refType,omitempty" xml:"refType,omitempty"`
	// Filters by reference selector type. Valid values: LABEL, VERSION.
	//
	// example:
	//
	// LABEL
	SelectorType *string `json:"selectorType,omitempty" xml:"selectorType,omitempty"`
	// Filters by reference selector value. Valid values: latest, a named label, HEAD, or a specific version.
	//
	// example:
	//
	// HEAD
	SelectorValue *string `json:"selectorValue,omitempty" xml:"selectorValue,omitempty"`
	// Filters by reference source identifier. The value must be agentcore or a configured source.
	//
	// example:
	//
	// agentcore
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ListSkillReferencesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillReferencesRequest) GoString() string {
	return s.String()
}

func (s *ListSkillReferencesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSkillReferencesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillReferencesRequest) GetRefType() *string {
	return s.RefType
}

func (s *ListSkillReferencesRequest) GetSelectorType() *string {
	return s.SelectorType
}

func (s *ListSkillReferencesRequest) GetSelectorValue() *string {
	return s.SelectorValue
}

func (s *ListSkillReferencesRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListSkillReferencesRequest) SetPageNo(v int32) *ListSkillReferencesRequest {
	s.PageNo = &v
	return s
}

func (s *ListSkillReferencesRequest) SetPageSize(v int32) *ListSkillReferencesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillReferencesRequest) SetRefType(v string) *ListSkillReferencesRequest {
	s.RefType = &v
	return s
}

func (s *ListSkillReferencesRequest) SetSelectorType(v string) *ListSkillReferencesRequest {
	s.SelectorType = &v
	return s
}

func (s *ListSkillReferencesRequest) SetSelectorValue(v string) *ListSkillReferencesRequest {
	s.SelectorValue = &v
	return s
}

func (s *ListSkillReferencesRequest) SetSourceId(v string) *ListSkillReferencesRequest {
	s.SourceId = &v
	return s
}

func (s *ListSkillReferencesRequest) Validate() error {
	return dara.Validate(s)
}
