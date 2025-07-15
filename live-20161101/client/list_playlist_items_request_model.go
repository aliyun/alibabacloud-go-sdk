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
	// The ID of the episode list. If the episode list was created by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, check the value of the response parameter ProgramId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The IDs of the episodes that you want to query. Separate episode IDs with commas (,). If you set this parameter, only the information about the specified episodes is returned. If you do not set this parameter, the information about all episodes in the episode list is returned.
	//
	// example:
	//
	// c10f3d63-eacf-4fbf-bd48-a07a6ba7****,c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ProgramItemIds *string `json:"ProgramItemIds,omitempty" xml:"ProgramItemIds,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
