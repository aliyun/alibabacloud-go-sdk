// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListPdpServicesRequest
	GetAlias() *string
	SetEnterpriseId(v int64) *ListPdpServicesRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListPdpServicesRequest
	GetName() *string
	SetOrderBy(v string) *ListPdpServicesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListPdpServicesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListPdpServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpServicesRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListPdpServicesRequest
	GetPbcId() *int64
	SetPdpServiceIds(v []*int64) *ListPdpServicesRequest
	GetPdpServiceIds() []*int64
	SetProductId(v int64) *ListPdpServicesRequest
	GetProductId() *int64
}

type ListPdpServicesRequest struct {
	Alias          *string  `json:"alias,omitempty" xml:"alias,omitempty"`
	EnterpriseId   *int64   `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string  `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string  `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PbcId          *int64   `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	PdpServiceIds  []*int64 `json:"pdpServiceIds,omitempty" xml:"pdpServiceIds,omitempty" type:"Repeated"`
	ProductId      *int64   `json:"productId,omitempty" xml:"productId,omitempty"`
}

func (s ListPdpServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpServicesRequest) GoString() string {
	return s.String()
}

func (s *ListPdpServicesRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListPdpServicesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListPdpServicesRequest) GetName() *string {
	return s.Name
}

func (s *ListPdpServicesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPdpServicesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListPdpServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpServicesRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListPdpServicesRequest) GetPdpServiceIds() []*int64 {
	return s.PdpServiceIds
}

func (s *ListPdpServicesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListPdpServicesRequest) SetAlias(v string) *ListPdpServicesRequest {
	s.Alias = &v
	return s
}

func (s *ListPdpServicesRequest) SetEnterpriseId(v int64) *ListPdpServicesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListPdpServicesRequest) SetName(v string) *ListPdpServicesRequest {
	s.Name = &v
	return s
}

func (s *ListPdpServicesRequest) SetOrderBy(v string) *ListPdpServicesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPdpServicesRequest) SetOrderDirection(v string) *ListPdpServicesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPdpServicesRequest) SetPageNumber(v int32) *ListPdpServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpServicesRequest) SetPageSize(v int32) *ListPdpServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpServicesRequest) SetPbcId(v int64) *ListPdpServicesRequest {
	s.PbcId = &v
	return s
}

func (s *ListPdpServicesRequest) SetPdpServiceIds(v []*int64) *ListPdpServicesRequest {
	s.PdpServiceIds = v
	return s
}

func (s *ListPdpServicesRequest) SetProductId(v int64) *ListPdpServicesRequest {
	s.ProductId = &v
	return s
}

func (s *ListPdpServicesRequest) Validate() error {
	return dara.Validate(s)
}
