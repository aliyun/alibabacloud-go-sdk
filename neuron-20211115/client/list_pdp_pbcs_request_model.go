// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpPbcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPdpPbcsRequest
	GetCompanyId() *int64
	SetDeveloperId(v string) *ListPdpPbcsRequest
	GetDeveloperId() *string
	SetKeyword(v string) *ListPdpPbcsRequest
	GetKeyword() *string
	SetOrderBy(v string) *ListPdpPbcsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpPbcsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpPbcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpPbcsRequest
	GetPageSize() *int32
	SetPbcIds(v []*int64) *ListPdpPbcsRequest
	GetPbcIds() []*int64
}

type ListPdpPbcsRequest struct {
	CompanyId      *int64   `json:"companyId,omitempty" xml:"companyId,omitempty"`
	DeveloperId    *string  `json:"developerId,omitempty" xml:"developerId,omitempty"`
	Keyword        *string  `json:"keyword,omitempty" xml:"keyword,omitempty"`
	OrderBy        *string  `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string  `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PbcIds         []*int64 `json:"pbcIds,omitempty" xml:"pbcIds,omitempty" type:"Repeated"`
}

func (s ListPdpPbcsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpPbcsRequest) GoString() string {
	return s.String()
}

func (s *ListPdpPbcsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPdpPbcsRequest) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *ListPdpPbcsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPdpPbcsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpPbcsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpPbcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpPbcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpPbcsRequest) GetPbcIds() []*int64 {
	return s.PbcIds
}

func (s *ListPdpPbcsRequest) SetCompanyId(v int64) *ListPdpPbcsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPdpPbcsRequest) SetDeveloperId(v string) *ListPdpPbcsRequest {
	s.DeveloperId = &v
	return s
}

func (s *ListPdpPbcsRequest) SetKeyword(v string) *ListPdpPbcsRequest {
	s.Keyword = &v
	return s
}

func (s *ListPdpPbcsRequest) SetOrderBy(v string) *ListPdpPbcsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpPbcsRequest) SetOrderDirection(v string) *ListPdpPbcsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpPbcsRequest) SetPageNumber(v int32) *ListPdpPbcsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpPbcsRequest) SetPageSize(v int32) *ListPdpPbcsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpPbcsRequest) SetPbcIds(v []*int64) *ListPdpPbcsRequest {
	s.PbcIds = v
	return s
}

func (s *ListPdpPbcsRequest) Validate() error {
	return dara.Validate(s)
}
