// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListPlaylistItemsRequest
	GetOwnerId() *int64
	SetProgramId(v string) *ListPlaylistItemsRequest
	GetProgramId() *string
	SetProgramItemIds(v string) *ListPlaylistItemsRequest
	GetProgramItemIds() *string
	SetRegionId(v string) *ListPlaylistItemsRequest
	GetRegionId() *string
}

type ListPlaylistItemsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the playlist. You can obtain the ID from the ProgramId parameter in the response of the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The IDs of the playlist items. Separate multiple IDs with commas (,). If you specify this parameter, only the information about the specified items is returned. If you leave this parameter empty, the information about all items in the playlist is returned.
	//
	// example:
	//
	// c10f3d63-eacf-4fbf-bd48-a07a6ba7****,c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ProgramItemIds *string `json:"ProgramItemIds,omitempty" xml:"ProgramItemIds,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPlaylistItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistItemsRequest) GoString() string {
	return s.String()
}

func (s *ListPlaylistItemsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPlaylistItemsRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *ListPlaylistItemsRequest) GetProgramItemIds() *string {
	return s.ProgramItemIds
}

func (s *ListPlaylistItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPlaylistItemsRequest) SetOwnerId(v int64) *ListPlaylistItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPlaylistItemsRequest) SetProgramId(v string) *ListPlaylistItemsRequest {
	s.ProgramId = &v
	return s
}

func (s *ListPlaylistItemsRequest) SetProgramItemIds(v string) *ListPlaylistItemsRequest {
	s.ProgramItemIds = &v
	return s
}

func (s *ListPlaylistItemsRequest) SetRegionId(v string) *ListPlaylistItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListPlaylistItemsRequest) Validate() error {
	return dara.Validate(s)
}
