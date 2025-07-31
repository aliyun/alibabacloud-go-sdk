// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataDomainsRequestListQuery) *ListDataDomainsRequest
	GetListQuery() *ListDataDomainsRequestListQuery
	SetOpTenantId(v int64) *ListDataDomainsRequest
	GetOpTenantId() *int64
}

type ListDataDomainsRequest struct {
	// This parameter is required.
	ListQuery *ListDataDomainsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDataDomainsRequest) GetListQuery() *ListDataDomainsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataDomainsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataDomainsRequest) SetListQuery(v *ListDataDomainsRequestListQuery) *ListDataDomainsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataDomainsRequest) SetOpTenantId(v int64) *ListDataDomainsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataDomainsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDataDomainsRequestListQuery struct {
	BizUnitIdList []*int64 `json:"BizUnitIdList,omitempty" xml:"BizUnitIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Keyword      *string  `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ParentIdList []*int64 `json:"ParentIdList,omitempty" xml:"ParentIdList,omitempty" type:"Repeated"`
}

func (s ListDataDomainsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataDomainsRequestListQuery) GetBizUnitIdList() []*int64 {
	return s.BizUnitIdList
}

func (s *ListDataDomainsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataDomainsRequestListQuery) GetParentIdList() []*int64 {
	return s.ParentIdList
}

func (s *ListDataDomainsRequestListQuery) SetBizUnitIdList(v []*int64) *ListDataDomainsRequestListQuery {
	s.BizUnitIdList = v
	return s
}

func (s *ListDataDomainsRequestListQuery) SetKeyword(v string) *ListDataDomainsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataDomainsRequestListQuery) SetParentIdList(v []*int64) *ListDataDomainsRequestListQuery {
	s.ParentIdList = v
	return s
}

func (s *ListDataDomainsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
