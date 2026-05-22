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
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RouteName  *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// This parameter is required.
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
