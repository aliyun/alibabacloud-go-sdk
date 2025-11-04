// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFilters(v string) *SearchMediaByFaceRequest
	GetCustomFilters() *string
	SetEntityId(v string) *SearchMediaByFaceRequest
	GetEntityId() *string
	SetFaceSearchToken(v string) *SearchMediaByFaceRequest
	GetFaceSearchToken() *string
	SetMediaType(v string) *SearchMediaByFaceRequest
	GetMediaType() *string
	SetNamespace(v string) *SearchMediaByFaceRequest
	GetNamespace() *string
	SetPageNo(v int32) *SearchMediaByFaceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaByFaceRequest
	GetPageSize() *int32
	SetPersonImageUrl(v string) *SearchMediaByFaceRequest
	GetPersonImageUrl() *string
	SetSearchLibName(v string) *SearchMediaByFaceRequest
	GetSearchLibName() *string
	SetUtcCreate(v string) *SearchMediaByFaceRequest
	GetUtcCreate() *string
}

type SearchMediaByFaceRequest struct {
	// example:
	//
	// {}
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 2d3bf1e35a1e42b5ab338d701efa****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The token that is used to identify the query. You can use this parameter in the SearchMediaClipByFace operation to specify the same query conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// zxtest-huangxuan-2023-3-7-V1
	FaceSearchToken *string `json:"FaceSearchToken,omitempty" xml:"FaceSearchToken,omitempty"`
	// The type of the media asset. Valid values:
	//
	// 	- image
	//
	// 	- video
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The URL of the face image.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://****.oss-cn-shanghai.aliyuncs.com/input/huangxuan****.jpg
	PersonImageUrl *string `json:"PersonImageUrl,omitempty" xml:"PersonImageUrl,omitempty"`
	// The name of the search library.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// example:
	//
	// {}
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
}

func (s SearchMediaByFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaByFaceRequest) GetCustomFilters() *string {
	return s.CustomFilters
}

func (s *SearchMediaByFaceRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaByFaceRequest) GetFaceSearchToken() *string {
	return s.FaceSearchToken
}

func (s *SearchMediaByFaceRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaByFaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchMediaByFaceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaByFaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaByFaceRequest) GetPersonImageUrl() *string {
	return s.PersonImageUrl
}

func (s *SearchMediaByFaceRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaByFaceRequest) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *SearchMediaByFaceRequest) SetCustomFilters(v string) *SearchMediaByFaceRequest {
	s.CustomFilters = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetEntityId(v string) *SearchMediaByFaceRequest {
	s.EntityId = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetFaceSearchToken(v string) *SearchMediaByFaceRequest {
	s.FaceSearchToken = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetMediaType(v string) *SearchMediaByFaceRequest {
	s.MediaType = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetNamespace(v string) *SearchMediaByFaceRequest {
	s.Namespace = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetPageNo(v int32) *SearchMediaByFaceRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetPageSize(v int32) *SearchMediaByFaceRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetPersonImageUrl(v string) *SearchMediaByFaceRequest {
	s.PersonImageUrl = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetSearchLibName(v string) *SearchMediaByFaceRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaByFaceRequest) SetUtcCreate(v string) *SearchMediaByFaceRequest {
	s.UtcCreate = &v
	return s
}

func (s *SearchMediaByFaceRequest) Validate() error {
	return dara.Validate(s)
}
