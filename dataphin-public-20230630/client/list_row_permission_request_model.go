// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListRowPermissionRequest
	GetOpTenantId() *int64
	SetPageRowPermissionQuery(v *ListRowPermissionRequestPageRowPermissionQuery) *ListRowPermissionRequest
	GetPageRowPermissionQuery() *ListRowPermissionRequestPageRowPermissionQuery
}

type ListRowPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PageRowPermissionQuery *ListRowPermissionRequestPageRowPermissionQuery `json:"PageRowPermissionQuery,omitempty" xml:"PageRowPermissionQuery,omitempty" type:"Struct"`
}

func (s ListRowPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionRequest) GoString() string {
	return s.String()
}

func (s *ListRowPermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRowPermissionRequest) GetPageRowPermissionQuery() *ListRowPermissionRequestPageRowPermissionQuery {
	return s.PageRowPermissionQuery
}

func (s *ListRowPermissionRequest) SetOpTenantId(v int64) *ListRowPermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRowPermissionRequest) SetPageRowPermissionQuery(v *ListRowPermissionRequestPageRowPermissionQuery) *ListRowPermissionRequest {
	s.PageRowPermissionQuery = v
	return s
}

func (s *ListRowPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionRequestPageRowPermissionQuery struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRowPermissionRequestPageRowPermissionQuery) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionRequestPageRowPermissionQuery) GoString() string {
	return s.String()
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) SetKeyword(v string) *ListRowPermissionRequestPageRowPermissionQuery {
	s.Keyword = &v
	return s
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) SetPageNum(v int32) *ListRowPermissionRequestPageRowPermissionQuery {
	s.PageNum = &v
	return s
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) SetPageSize(v int32) *ListRowPermissionRequestPageRowPermissionQuery {
	s.PageSize = &v
	return s
}

func (s *ListRowPermissionRequestPageRowPermissionQuery) Validate() error {
	return dara.Validate(s)
}
