// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeletePlaylistItemsRequest
	GetOwnerId() *int64
	SetProgramId(v string) *DeletePlaylistItemsRequest
	GetProgramId() *string
	SetProgramItemIds(v string) *DeletePlaylistItemsRequest
	GetProgramItemIds() *string
	SetRegionId(v string) *DeletePlaylistItemsRequest
	GetRegionId() *string
}

type DeletePlaylistItemsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the episode list. If the episode list was created by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, check the value of the response parameter ProgramId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The IDs of the episodes that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["c09f3d63-eacf-4fbf-bd48-a07a6ba7****","c10f3d63-eacf-4fbf-bd48-a07a6ba7****"]
	ProgramItemIds *string `json:"ProgramItemIds,omitempty" xml:"ProgramItemIds,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePlaylistItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistItemsRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaylistItemsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePlaylistItemsRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *DeletePlaylistItemsRequest) GetProgramItemIds() *string {
	return s.ProgramItemIds
}

func (s *DeletePlaylistItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePlaylistItemsRequest) SetOwnerId(v int64) *DeletePlaylistItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePlaylistItemsRequest) SetProgramId(v string) *DeletePlaylistItemsRequest {
	s.ProgramId = &v
	return s
}

func (s *DeletePlaylistItemsRequest) SetProgramItemIds(v string) *DeletePlaylistItemsRequest {
	s.ProgramItemIds = &v
	return s
}

func (s *DeletePlaylistItemsRequest) SetRegionId(v string) *DeletePlaylistItemsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePlaylistItemsRequest) Validate() error {
	return dara.Validate(s)
}
