// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOsType(v int32) *ListAppsRequest
	GetOsType() *int32
	SetPage(v string) *ListAppsRequest
	GetPage() *string
	SetPageSize(v string) *ListAppsRequest
	GetPageSize() *string
	SetProductId(v string) *ListAppsRequest
	GetProductId() *string
}

type ListAppsRequest struct {
	// example:
	//
	// 1
	OsType *int32 `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) GetOsType() *int32 {
	return s.OsType
}

func (s *ListAppsRequest) GetPage() *string {
	return s.Page
}

func (s *ListAppsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAppsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListAppsRequest) SetOsType(v int32) *ListAppsRequest {
	s.OsType = &v
	return s
}

func (s *ListAppsRequest) SetPage(v string) *ListAppsRequest {
	s.Page = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v string) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetProductId(v string) *ListAppsRequest {
	s.ProductId = &v
	return s
}

func (s *ListAppsRequest) Validate() error {
	return dara.Validate(s)
}
