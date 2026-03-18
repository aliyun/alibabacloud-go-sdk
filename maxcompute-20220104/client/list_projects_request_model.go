// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListSystemCatalog(v bool) *ListProjectsRequest
	GetListSystemCatalog() *bool
	SetMarker(v string) *ListProjectsRequest
	GetMarker() *string
	SetMaxItem(v int32) *ListProjectsRequest
	GetMaxItem() *int32
	SetPrefix(v string) *ListProjectsRequest
	GetPrefix() *string
	SetQuotaName(v string) *ListProjectsRequest
	GetQuotaName() *string
	SetQuotaNickName(v string) *ListProjectsRequest
	GetQuotaNickName() *string
	SetRegion(v string) *ListProjectsRequest
	GetRegion() *string
	SetSaleTags(v string) *ListProjectsRequest
	GetSaleTags() *string
	SetTenantId(v string) *ListProjectsRequest
	GetTenantId() *string
	SetType(v string) *ListProjectsRequest
	GetType() *string
}

type ListProjectsRequest struct {
	ListSystemCatalog *bool   `json:"listSystemCatalog,omitempty" xml:"listSystemCatalog,omitempty"`
	Marker            *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxItem           *int32  `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	Prefix            *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	QuotaName         *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	QuotaNickName     *string `json:"quotaNickName,omitempty" xml:"quotaNickName,omitempty"`
	Region            *string `json:"region,omitempty" xml:"region,omitempty"`
	SaleTags          *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	TenantId          *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type              *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetListSystemCatalog() *bool {
	return s.ListSystemCatalog
}

func (s *ListProjectsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListProjectsRequest) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListProjectsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListProjectsRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListProjectsRequest) GetQuotaNickName() *string {
	return s.QuotaNickName
}

func (s *ListProjectsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListProjectsRequest) GetSaleTags() *string {
	return s.SaleTags
}

func (s *ListProjectsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListProjectsRequest) GetType() *string {
	return s.Type
}

func (s *ListProjectsRequest) SetListSystemCatalog(v bool) *ListProjectsRequest {
	s.ListSystemCatalog = &v
	return s
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxItem(v int32) *ListProjectsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaName(v string) *ListProjectsRequest {
	s.QuotaName = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaNickName(v string) *ListProjectsRequest {
	s.QuotaNickName = &v
	return s
}

func (s *ListProjectsRequest) SetRegion(v string) *ListProjectsRequest {
	s.Region = &v
	return s
}

func (s *ListProjectsRequest) SetSaleTags(v string) *ListProjectsRequest {
	s.SaleTags = &v
	return s
}

func (s *ListProjectsRequest) SetTenantId(v string) *ListProjectsRequest {
	s.TenantId = &v
	return s
}

func (s *ListProjectsRequest) SetType(v string) *ListProjectsRequest {
	s.Type = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
