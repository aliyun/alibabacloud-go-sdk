// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPbcAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustry(v string) *SearchPbcAssetsRequest
	GetIndustry() *string
	SetKeyword(v string) *SearchPbcAssetsRequest
	GetKeyword() *string
	SetOrderBy(v string) *SearchPbcAssetsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchPbcAssetsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *SearchPbcAssetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchPbcAssetsRequest
	GetPageSize() *int32
	SetType(v string) *SearchPbcAssetsRequest
	GetType() *string
}

type SearchPbcAssetsRequest struct {
	Industry       *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Keyword        *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	OrderBy        *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	PageNumber     *int32  `json:"page_number,omitempty" xml:"page_number,omitempty"`
	PageSize       *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchPbcAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchPbcAssetsRequest) GoString() string {
	return s.String()
}

func (s *SearchPbcAssetsRequest) GetIndustry() *string {
	return s.Industry
}

func (s *SearchPbcAssetsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchPbcAssetsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchPbcAssetsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchPbcAssetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchPbcAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchPbcAssetsRequest) GetType() *string {
	return s.Type
}

func (s *SearchPbcAssetsRequest) SetIndustry(v string) *SearchPbcAssetsRequest {
	s.Industry = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetKeyword(v string) *SearchPbcAssetsRequest {
	s.Keyword = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetOrderBy(v string) *SearchPbcAssetsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetOrderDirection(v string) *SearchPbcAssetsRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetPageNumber(v int32) *SearchPbcAssetsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetPageSize(v int32) *SearchPbcAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPbcAssetsRequest) SetType(v string) *SearchPbcAssetsRequest {
	s.Type = &v
	return s
}

func (s *SearchPbcAssetsRequest) Validate() error {
	return dara.Validate(s)
}
