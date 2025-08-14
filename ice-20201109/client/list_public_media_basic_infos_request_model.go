// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicMediaBasicInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *ListPublicMediaBasicInfosRequest
	GetBusinessType() *string
	SetIncludeFileBasicInfo(v bool) *ListPublicMediaBasicInfosRequest
	GetIncludeFileBasicInfo() *bool
	SetMaxResults(v int32) *ListPublicMediaBasicInfosRequest
	GetMaxResults() *int32
	SetMediaTagId(v string) *ListPublicMediaBasicInfosRequest
	GetMediaTagId() *string
	SetNextToken(v string) *ListPublicMediaBasicInfosRequest
	GetNextToken() *string
	SetPageNo(v int32) *ListPublicMediaBasicInfosRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListPublicMediaBasicInfosRequest
	GetPageSize() *int32
}

type ListPublicMediaBasicInfosRequest struct {
	// The business type of the media asset. Valid values:
	//
	// 	- sticker
	//
	// 	- bgm
	//
	// 	- bgi
	//
	// example:
	//
	// sticker
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// Specifies whether to return the basic information of the media asset.
	//
	// example:
	//
	// true
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
	// The maximum number of entries to return.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The media tag. All media assets that contain the specified media tag are returned. Valid values:
	//
	// 	- Sticker tags:
	//
	//     	- sticker-atmosphere
	//
	//     	- sticker-bubble
	//
	//     	- sticker-cute
	//
	//     	- sticker-daily
	//
	//     	- sticker-expression
	//
	//     	- sticker-gif
	//
	// 	- Background music (BGM) tags:
	//
	//     	- bgm-romantic
	//
	//     	- bgm-cuisine
	//
	//     	- bgm-chinese-style
	//
	//     	- bgm-upbeat
	//
	//     	- bgm-dynamic
	//
	//     	- bgm-relaxing
	//
	//     	- bgm-quirky
	//
	//     	- bgm-beauty
	//
	// 	- Background image (BGI) tags:
	//
	//     	- bgi-grad
	//
	//     	- bgi-solid
	//
	//     	- bgi-pic
	//
	// example:
	//
	// ticker-atmosphere
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// pSa1SQ0wCe5pzVrQ6mWZEw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPublicMediaBasicInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListPublicMediaBasicInfosRequest) GetIncludeFileBasicInfo() *bool {
	return s.IncludeFileBasicInfo
}

func (s *ListPublicMediaBasicInfosRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicMediaBasicInfosRequest) GetMediaTagId() *string {
	return s.MediaTagId
}

func (s *ListPublicMediaBasicInfosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicMediaBasicInfosRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPublicMediaBasicInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublicMediaBasicInfosRequest) SetBusinessType(v string) *ListPublicMediaBasicInfosRequest {
	s.BusinessType = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListPublicMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetMaxResults(v int32) *ListPublicMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetMediaTagId(v string) *ListPublicMediaBasicInfosRequest {
	s.MediaTagId = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetNextToken(v string) *ListPublicMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetPageNo(v int32) *ListPublicMediaBasicInfosRequest {
	s.PageNo = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetPageSize(v int32) *ListPublicMediaBasicInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) Validate() error {
	return dara.Validate(s)
}
