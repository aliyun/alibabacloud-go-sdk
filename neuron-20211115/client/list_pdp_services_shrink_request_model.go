// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListPdpServicesShrinkRequest
	GetAlias() *string
	SetEnterpriseId(v int64) *ListPdpServicesShrinkRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListPdpServicesShrinkRequest
	GetName() *string
	SetOrderBy(v string) *ListPdpServicesShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpServicesShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpServicesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpServicesShrinkRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListPdpServicesShrinkRequest
	GetPbcId() *int64
	SetPdpServiceIdsShrink(v string) *ListPdpServicesShrinkRequest
	GetPdpServiceIdsShrink() *string
	SetProductId(v int64) *ListPdpServicesShrinkRequest
	GetProductId() *int64
}

type ListPdpServicesShrinkRequest struct {
	Alias               *string `json:"alias,omitempty" xml:"alias,omitempty"`
	EnterpriseId        *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy             *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection      *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber          *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize            *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PbcId               *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	PdpServiceIdsShrink *string `json:"pdpServiceIds,omitempty" xml:"pdpServiceIds,omitempty"`
	ProductId           *int64  `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s ListPdpServicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPdpServicesShrinkRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListPdpServicesShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListPdpServicesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListPdpServicesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpServicesShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpServicesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpServicesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpServicesShrinkRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListPdpServicesShrinkRequest) GetPdpServiceIdsShrink() *string {
	return s.PdpServiceIdsShrink
}

func (s *ListPdpServicesShrinkRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListPdpServicesShrinkRequest) SetAlias(v string) *ListPdpServicesShrinkRequest {
	s.Alias = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetEnterpriseId(v int64) *ListPdpServicesShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetName(v string) *ListPdpServicesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetOrderBy(v string) *ListPdpServicesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetOrderDirection(v string) *ListPdpServicesShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetPageNumber(v int32) *ListPdpServicesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetPageSize(v int32) *ListPdpServicesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetPbcId(v int64) *ListPdpServicesShrinkRequest {
	s.PbcId = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetPdpServiceIdsShrink(v string) *ListPdpServicesShrinkRequest {
	s.PdpServiceIdsShrink = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) SetProductId(v int64) *ListPdpServicesShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *ListPdpServicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
