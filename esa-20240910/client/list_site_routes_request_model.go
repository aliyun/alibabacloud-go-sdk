// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListSiteRoutesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListSiteRoutesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListSiteRoutesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSiteRoutesRequest
	GetPageSize() *int32
	SetRouteName(v string) *ListSiteRoutesRequest
	GetRouteName() *string
	SetSiteId(v int64) *ListSiteRoutesRequest
	GetSiteId() *int64
}

type ListSiteRoutesRequest struct {
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Use this parameter to query global or feature-specific configurations. This parameter takes effect only if the `functionName` parameter is also specified.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number. The default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page, with a maximum of 500. The default is 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The route name. Use this parameter to find a route by its name. This parameter takes effect only if the `functionName` parameter is also specified.
	//
	// example:
	//
	// test_route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The site ID. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListSiteRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSiteRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListSiteRoutesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteRoutesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListSiteRoutesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSiteRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSiteRoutesRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *ListSiteRoutesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSiteRoutesRequest) SetConfigId(v int64) *ListSiteRoutesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListSiteRoutesRequest) SetConfigType(v string) *ListSiteRoutesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListSiteRoutesRequest) SetPageNumber(v int32) *ListSiteRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSiteRoutesRequest) SetPageSize(v int32) *ListSiteRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSiteRoutesRequest) SetRouteName(v string) *ListSiteRoutesRequest {
	s.RouteName = &v
	return s
}

func (s *ListSiteRoutesRequest) SetSiteId(v int64) *ListSiteRoutesRequest {
	s.SiteId = &v
	return s
}

func (s *ListSiteRoutesRequest) Validate() error {
	return dara.Validate(s)
}
