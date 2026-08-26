// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeletePlaylistRequest
	GetOwnerId() *int64
	SetProgramId(v string) *DeletePlaylistRequest
	GetProgramId() *string
	SetRegionId(v string) *DeletePlaylistRequest
	GetRegionId() *string
}

type DeletePlaylistRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the playlist. You can obtain the playlist ID from the ProgramId parameter in the response of the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaylistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePlaylistRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *DeletePlaylistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePlaylistRequest) SetOwnerId(v int64) *DeletePlaylistRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePlaylistRequest) SetProgramId(v string) *DeletePlaylistRequest {
	s.ProgramId = &v
	return s
}

func (s *DeletePlaylistRequest) SetRegionId(v string) *DeletePlaylistRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePlaylistRequest) Validate() error {
	return dara.Validate(s)
}
