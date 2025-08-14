// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPublicMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorized(v bool) *SearchPublicMediaInfoRequest
	GetAuthorized() *bool
	SetDynamicMetaDataMatchFields(v string) *SearchPublicMediaInfoRequest
	GetDynamicMetaDataMatchFields() *string
	SetEntityId(v string) *SearchPublicMediaInfoRequest
	GetEntityId() *string
	SetFavorite(v bool) *SearchPublicMediaInfoRequest
	GetFavorite() *bool
	SetMediaIds(v string) *SearchPublicMediaInfoRequest
	GetMediaIds() *string
	SetPageNo(v int32) *SearchPublicMediaInfoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchPublicMediaInfoRequest
	GetPageSize() *int32
	SetSortBy(v string) *SearchPublicMediaInfoRequest
	GetSortBy() *string
}

type SearchPublicMediaInfoRequest struct {
	// example:
	//
	// true
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// example:
	//
	// "ApprovalStatus=\\"Available\\"&amp;MaterialBags=\\"boutiquemusic\\"&amp;Mood=\\"Nervous\\""
	DynamicMetaDataMatchFields *string `json:"DynamicMetaDataMatchFields,omitempty" xml:"DynamicMetaDataMatchFields,omitempty"`
	// example:
	//
	// Copyright_Music
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// true
	Favorite *bool `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****,****15d4a4b0448391508f2cb486****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// UsageCount:Desc,UnitPrice:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SearchPublicMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoRequest) GetAuthorized() *bool {
	return s.Authorized
}

func (s *SearchPublicMediaInfoRequest) GetDynamicMetaDataMatchFields() *string {
	return s.DynamicMetaDataMatchFields
}

func (s *SearchPublicMediaInfoRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchPublicMediaInfoRequest) GetFavorite() *bool {
	return s.Favorite
}

func (s *SearchPublicMediaInfoRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *SearchPublicMediaInfoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchPublicMediaInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchPublicMediaInfoRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchPublicMediaInfoRequest) SetAuthorized(v bool) *SearchPublicMediaInfoRequest {
	s.Authorized = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetDynamicMetaDataMatchFields(v string) *SearchPublicMediaInfoRequest {
	s.DynamicMetaDataMatchFields = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetEntityId(v string) *SearchPublicMediaInfoRequest {
	s.EntityId = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetFavorite(v bool) *SearchPublicMediaInfoRequest {
	s.Favorite = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetMediaIds(v string) *SearchPublicMediaInfoRequest {
	s.MediaIds = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetPageNo(v int32) *SearchPublicMediaInfoRequest {
	s.PageNo = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetPageSize(v int32) *SearchPublicMediaInfoRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) SetSortBy(v string) *SearchPublicMediaInfoRequest {
	s.SortBy = &v
	return s
}

func (s *SearchPublicMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
