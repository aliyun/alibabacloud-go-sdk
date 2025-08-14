// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaClipByFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *SearchMediaClipByFaceRequest
	GetEntityId() *string
	SetFaceSearchToken(v string) *SearchMediaClipByFaceRequest
	GetFaceSearchToken() *string
	SetMediaId(v string) *SearchMediaClipByFaceRequest
	GetMediaId() *string
	SetPageNo(v int32) *SearchMediaClipByFaceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaClipByFaceRequest
	GetPageSize() *int32
	SetSearchLibName(v string) *SearchMediaClipByFaceRequest
	GetSearchLibName() *string
}

type SearchMediaClipByFaceRequest struct {
	// The ID of the entity.
	//
	// example:
	//
	// 2d3bf1e35a1e42b5ab338d701efa****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The value of this parameter is the same as that of the FaceSearchToken parameter in the SearchMediaByFace request. This specifies to return media asset clips that meet the same query conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// zxtest-huangxuan-2023-3-7-V1
	FaceSearchToken *string `json:"FaceSearchToken,omitempty" xml:"FaceSearchToken,omitempty"`
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b187b3620c8490886cfc2a9578c****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
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
	// The name of the search library.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s SearchMediaClipByFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *SearchMediaClipByFaceRequest) GetFaceSearchToken() *string {
	return s.FaceSearchToken
}

func (s *SearchMediaClipByFaceRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaClipByFaceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaClipByFaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaClipByFaceRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaClipByFaceRequest) SetEntityId(v string) *SearchMediaClipByFaceRequest {
	s.EntityId = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) SetFaceSearchToken(v string) *SearchMediaClipByFaceRequest {
	s.FaceSearchToken = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) SetMediaId(v string) *SearchMediaClipByFaceRequest {
	s.MediaId = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) SetPageNo(v int32) *SearchMediaClipByFaceRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) SetPageSize(v int32) *SearchMediaClipByFaceRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) SetSearchLibName(v string) *SearchMediaClipByFaceRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaClipByFaceRequest) Validate() error {
	return dara.Validate(s)
}
