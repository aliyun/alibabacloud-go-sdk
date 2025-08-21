// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListAlbumDetailRequest
	GetId() *int64
	SetPageNum(v int32) *ListAlbumDetailRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAlbumDetailRequest
	GetPageSize() *int32
}

type ListAlbumDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlbumDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *ListAlbumDetailRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAlbumDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlbumDetailRequest) SetId(v int64) *ListAlbumDetailRequest {
	s.Id = &v
	return s
}

func (s *ListAlbumDetailRequest) SetPageNum(v int32) *ListAlbumDetailRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlbumDetailRequest) SetPageSize(v int32) *ListAlbumDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlbumDetailRequest) Validate() error {
	return dara.Validate(s)
}
