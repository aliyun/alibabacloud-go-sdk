// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetadataInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListMetadataInfosRequest
	GetEnv() *string
	SetNamespaceId(v int64) *ListMetadataInfosRequest
	GetNamespaceId() *int64
	SetNamespaceName(v string) *ListMetadataInfosRequest
	GetNamespaceName() *string
	SetOrderBy(v string) *ListMetadataInfosRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetadataInfosRequest
	GetOrderDirection() *string
	SetOrgId(v int32) *ListMetadataInfosRequest
	GetOrgId() *int32
	SetPageNumber(v int32) *ListMetadataInfosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMetadataInfosRequest
	GetPageSize() *int32
	SetPbcId(v int64) *ListMetadataInfosRequest
	GetPbcId() *int64
}

type ListMetadataInfosRequest struct {
	Env            *string `json:"env,omitempty" xml:"env,omitempty"`
	NamespaceId    *int64  `json:"namespace_id,omitempty" xml:"namespace_id,omitempty"`
	NamespaceName  *string `json:"namespace_name,omitempty" xml:"namespace_name,omitempty"`
	OrderBy        *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	OrgId          *int32  `json:"org_id,omitempty" xml:"org_id,omitempty"`
	PageNumber     *int32  `json:"page_number,omitempty" xml:"page_number,omitempty"`
	PageSize       *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// This parameter is required.
	PbcId *int64 `json:"pbc_id,omitempty" xml:"pbc_id,omitempty"`
}

func (s ListMetadataInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetadataInfosRequest) GoString() string {
	return s.String()
}

func (s *ListMetadataInfosRequest) GetEnv() *string {
	return s.Env
}

func (s *ListMetadataInfosRequest) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *ListMetadataInfosRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListMetadataInfosRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetadataInfosRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetadataInfosRequest) GetOrgId() *int32 {
	return s.OrgId
}

func (s *ListMetadataInfosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetadataInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetadataInfosRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListMetadataInfosRequest) SetEnv(v string) *ListMetadataInfosRequest {
	s.Env = &v
	return s
}

func (s *ListMetadataInfosRequest) SetNamespaceId(v int64) *ListMetadataInfosRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListMetadataInfosRequest) SetNamespaceName(v string) *ListMetadataInfosRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListMetadataInfosRequest) SetOrderBy(v string) *ListMetadataInfosRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetadataInfosRequest) SetOrderDirection(v string) *ListMetadataInfosRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetadataInfosRequest) SetOrgId(v int32) *ListMetadataInfosRequest {
	s.OrgId = &v
	return s
}

func (s *ListMetadataInfosRequest) SetPageNumber(v int32) *ListMetadataInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetadataInfosRequest) SetPageSize(v int32) *ListMetadataInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetadataInfosRequest) SetPbcId(v int64) *ListMetadataInfosRequest {
	s.PbcId = &v
	return s
}

func (s *ListMetadataInfosRequest) Validate() error {
	return dara.Validate(s)
}
