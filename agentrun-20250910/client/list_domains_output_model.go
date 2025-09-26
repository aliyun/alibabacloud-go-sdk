// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DomainInfo) *ListDomainsOutput
	GetItems() []*DomainInfo
	SetPageNumber(v string) *ListDomainsOutput
	GetPageNumber() *string
	SetPageSize(v string) *ListDomainsOutput
	GetPageSize() *string
	SetTotalCount(v string) *ListDomainsOutput
	GetTotalCount() *string
}

type ListDomainsOutput struct {
	Items      []*DomainInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *string       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *string       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *string       `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDomainsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsOutput) GoString() string {
	return s.String()
}

func (s *ListDomainsOutput) GetItems() []*DomainInfo {
	return s.Items
}

func (s *ListDomainsOutput) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDomainsOutput) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDomainsOutput) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListDomainsOutput) SetItems(v []*DomainInfo) *ListDomainsOutput {
	s.Items = v
	return s
}

func (s *ListDomainsOutput) SetPageNumber(v string) *ListDomainsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsOutput) SetPageSize(v string) *ListDomainsOutput {
	s.PageSize = &v
	return s
}

func (s *ListDomainsOutput) SetTotalCount(v string) *ListDomainsOutput {
	s.TotalCount = &v
	return s
}

func (s *ListDomainsOutput) Validate() error {
	return dara.Validate(s)
}
