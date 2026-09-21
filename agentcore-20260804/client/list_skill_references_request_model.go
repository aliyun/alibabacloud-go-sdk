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
	SetSelectorType(v string) *ListSkillReferencesRequest
	GetSelectorType() *string
	SetSelectorValue(v string) *ListSkillReferencesRequest
	GetSelectorValue() *string
}

type ListSkillReferencesRequest struct {
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// The number of entries per page. If this parameter is not specified, the server-side default value is used.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters results by reference selector type. Valid values: LABEL and VERSION.
	//
	// example:
	//
	// LABEL
	SelectorType *string `json:"selectorType,omitempty" xml:"selectorType,omitempty"`
	// Filters results by reference selector value, such as latest, a named label, HEAD, or a specific version.
	//
	// example:
	//
	// HEAD
	SelectorValue *string `json:"selectorValue,omitempty" xml:"selectorValue,omitempty"`
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

func (s *ListSkillReferencesRequest) GetSelectorType() *string {
	return s.SelectorType
}

func (s *ListSkillReferencesRequest) GetSelectorValue() *string {
	return s.SelectorValue
}

func (s *ListSkillReferencesRequest) SetPageNo(v int32) *ListSkillReferencesRequest {
	s.PageNo = &v
	return s
}

func (s *ListSkillReferencesRequest) SetPageSize(v int32) *ListSkillReferencesRequest {
	s.PageSize = &v
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

func (s *ListSkillReferencesRequest) Validate() error {
	return dara.Validate(s)
}
