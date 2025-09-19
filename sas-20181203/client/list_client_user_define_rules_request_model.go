// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUserDefineRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListClientUserDefineRulesRequest
	GetCurrentPage() *int32
	SetName(v string) *ListClientUserDefineRulesRequest
	GetName() *string
	SetPageSize(v int32) *ListClientUserDefineRulesRequest
	GetPageSize() *int32
	SetType(v []*int32) *ListClientUserDefineRulesRequest
	GetType() []*int32
}

type ListClientUserDefineRulesRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The types of rules.
	Type []*int32 `json:"Type,omitempty" xml:"Type,omitempty" type:"Repeated"`
}

func (s ListClientUserDefineRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRulesRequest) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientUserDefineRulesRequest) GetName() *string {
	return s.Name
}

func (s *ListClientUserDefineRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClientUserDefineRulesRequest) GetType() []*int32 {
	return s.Type
}

func (s *ListClientUserDefineRulesRequest) SetCurrentPage(v int32) *ListClientUserDefineRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientUserDefineRulesRequest) SetName(v string) *ListClientUserDefineRulesRequest {
	s.Name = &v
	return s
}

func (s *ListClientUserDefineRulesRequest) SetPageSize(v int32) *ListClientUserDefineRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientUserDefineRulesRequest) SetType(v []*int32) *ListClientUserDefineRulesRequest {
	s.Type = v
	return s
}

func (s *ListClientUserDefineRulesRequest) Validate() error {
	return dara.Validate(s)
}
