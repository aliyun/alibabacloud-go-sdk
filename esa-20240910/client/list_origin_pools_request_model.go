// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListOriginPoolsRequest
	GetMatchType() *string
	SetName(v string) *ListOriginPoolsRequest
	GetName() *string
	SetOrderBy(v string) *ListOriginPoolsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListOriginPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOriginPoolsRequest
	GetPageSize() *int32
	SetSiteId(v int64) *ListOriginPoolsRequest
	GetSiteId() *int64
}

type ListOriginPoolsRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListOriginPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListOriginPoolsRequest) GetName() *string {
	return s.Name
}

func (s *ListOriginPoolsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListOriginPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOriginPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOriginPoolsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginPoolsRequest) SetMatchType(v string) *ListOriginPoolsRequest {
	s.MatchType = &v
	return s
}

func (s *ListOriginPoolsRequest) SetName(v string) *ListOriginPoolsRequest {
	s.Name = &v
	return s
}

func (s *ListOriginPoolsRequest) SetOrderBy(v string) *ListOriginPoolsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListOriginPoolsRequest) SetPageNumber(v int32) *ListOriginPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOriginPoolsRequest) SetPageSize(v int32) *ListOriginPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOriginPoolsRequest) SetSiteId(v int64) *ListOriginPoolsRequest {
	s.SiteId = &v
	return s
}

func (s *ListOriginPoolsRequest) Validate() error {
	return dara.Validate(s)
}
