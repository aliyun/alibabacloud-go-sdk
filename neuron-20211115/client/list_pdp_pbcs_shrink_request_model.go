// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpPbcsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ListPdpPbcsShrinkRequest
	GetCompanyId() *int64
	SetDeveloperId(v string) *ListPdpPbcsShrinkRequest
	GetDeveloperId() *string
	SetKeyword(v string) *ListPdpPbcsShrinkRequest
	GetKeyword() *string
	SetOrderBy(v string) *ListPdpPbcsShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpPbcsShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpPbcsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpPbcsShrinkRequest
	GetPageSize() *int32
	SetPbcIdsShrink(v string) *ListPdpPbcsShrinkRequest
	GetPbcIdsShrink() *string
}

type ListPdpPbcsShrinkRequest struct {
	CompanyId      *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	DeveloperId    *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	Keyword        *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PbcIdsShrink   *string `json:"pbcIds,omitempty" xml:"pbcIds,omitempty"`
}

func (s ListPdpPbcsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpPbcsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPdpPbcsShrinkRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPdpPbcsShrinkRequest) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *ListPdpPbcsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPdpPbcsShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpPbcsShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpPbcsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpPbcsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpPbcsShrinkRequest) GetPbcIdsShrink() *string {
	return s.PbcIdsShrink
}

func (s *ListPdpPbcsShrinkRequest) SetCompanyId(v int64) *ListPdpPbcsShrinkRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetDeveloperId(v string) *ListPdpPbcsShrinkRequest {
	s.DeveloperId = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetKeyword(v string) *ListPdpPbcsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetOrderBy(v string) *ListPdpPbcsShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetOrderDirection(v string) *ListPdpPbcsShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetPageNumber(v int32) *ListPdpPbcsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetPageSize(v int32) *ListPdpPbcsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) SetPbcIdsShrink(v string) *ListPdpPbcsShrinkRequest {
	s.PbcIdsShrink = &v
	return s
}

func (s *ListPdpPbcsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
