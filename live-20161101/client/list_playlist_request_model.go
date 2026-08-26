// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListPlaylistRequest
	GetOwnerId() *int64
	SetPage(v int32) *ListPlaylistRequest
	GetPage() *int32
	SetPageSize(v int32) *ListPlaylistRequest
	GetPageSize() *int32
	SetProgramId(v string) *ListPlaylistRequest
	GetProgramId() *string
	SetRegionId(v string) *ListPlaylistRequest
	GetRegionId() *string
}

type ListPlaylistRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the playlist. If you specify this parameter, only the information about the specified playlist is returned. If you leave this parameter empty, the information about all playlists that belong to your account is returned. Use the ProgramId value returned by the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation.
	//
	// example:
	//
	// c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistRequest) GoString() string {
	return s.String()
}

func (s *ListPlaylistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPlaylistRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListPlaylistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPlaylistRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *ListPlaylistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPlaylistRequest) SetOwnerId(v int64) *ListPlaylistRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPlaylistRequest) SetPage(v int32) *ListPlaylistRequest {
	s.Page = &v
	return s
}

func (s *ListPlaylistRequest) SetPageSize(v int32) *ListPlaylistRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlaylistRequest) SetProgramId(v string) *ListPlaylistRequest {
	s.ProgramId = &v
	return s
}

func (s *ListPlaylistRequest) SetRegionId(v string) *ListPlaylistRequest {
	s.RegionId = &v
	return s
}

func (s *ListPlaylistRequest) Validate() error {
	return dara.Validate(s)
}
