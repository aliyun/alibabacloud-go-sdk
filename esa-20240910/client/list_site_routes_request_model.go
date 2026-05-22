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
	// The configuration type. You can use this parameter to query global configurations or feature configurations. This parameter takes effect only if the functionName parameter is passed.
	//
	// Valid values:
	//
	// 	- global
	//
	// 	- rule
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 500. Default value: 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. This parameter takes effect only when parameter functionName is specified.
	//
	// example:
	//
	// test_route
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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
