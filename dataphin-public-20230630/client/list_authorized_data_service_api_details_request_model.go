// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDataServiceApiDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAuthorizedDataServiceApiDetailsRequestListQuery) *ListAuthorizedDataServiceApiDetailsRequest
	GetListQuery() *ListAuthorizedDataServiceApiDetailsRequestListQuery
	SetOpTenantId(v int64) *ListAuthorizedDataServiceApiDetailsRequest
	GetOpTenantId() *int64
}

type ListAuthorizedDataServiceApiDetailsRequest struct {
	// This parameter is required.
	ListQuery *ListAuthorizedDataServiceApiDetailsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsRequest) GetListQuery() *ListAuthorizedDataServiceApiDetailsRequestListQuery {
	return s.ListQuery
}

func (s *ListAuthorizedDataServiceApiDetailsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAuthorizedDataServiceApiDetailsRequest) SetListQuery(v *ListAuthorizedDataServiceApiDetailsRequestListQuery) *ListAuthorizedDataServiceApiDetailsRequest {
	s.ListQuery = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsRequest) SetOpTenantId(v int64) *ListAuthorizedDataServiceApiDetailsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsRequest) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedDataServiceApiDetailsRequestListQuery struct {
	// AppKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 200000000
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) SetAppKey(v int64) *ListAuthorizedDataServiceApiDetailsRequestListQuery {
	s.AppKey = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) SetPage(v int32) *ListAuthorizedDataServiceApiDetailsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) SetPageSize(v int32) *ListAuthorizedDataServiceApiDetailsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
