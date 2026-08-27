// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActiveAnnouncementsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListActiveAnnouncementsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListActiveAnnouncementsRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListActiveAnnouncementsRequest
	GetTenantId() *string
}

type ListActiveAnnouncementsRequest struct {
	// The page number for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page for pagination. Default value: 100. Maximum value: 500.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListActiveAnnouncementsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActiveAnnouncementsRequest) GoString() string {
	return s.String()
}

func (s *ListActiveAnnouncementsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListActiveAnnouncementsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListActiveAnnouncementsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListActiveAnnouncementsRequest) SetPageNumber(v int64) *ListActiveAnnouncementsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetPageSize(v int64) *ListActiveAnnouncementsRequest {
	s.PageSize = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetTenantId(v string) *ListActiveAnnouncementsRequest {
	s.TenantId = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) Validate() error {
	return dara.Validate(s)
}
