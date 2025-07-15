// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopPlaylistRequest
	GetOwnerId() *int64
	SetProgramId(v string) *StopPlaylistRequest
	GetProgramId() *string
	SetRegionId(v string) *StopPlaylistRequest
	GetRegionId() *string
}

type StopPlaylistRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the episode list. If the episode list was created by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, check the value of the response parameter ProgramId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPlaylistRequest) GoString() string {
	return s.String()
}

func (s *StopPlaylistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopPlaylistRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *StopPlaylistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopPlaylistRequest) SetOwnerId(v int64) *StopPlaylistRequest {
	s.OwnerId = &v
	return s
}

func (s *StopPlaylistRequest) SetProgramId(v string) *StopPlaylistRequest {
	s.ProgramId = &v
	return s
}

func (s *StopPlaylistRequest) SetRegionId(v string) *StopPlaylistRequest {
	s.RegionId = &v
	return s
}

func (s *StopPlaylistRequest) Validate() error {
	return dara.Validate(s)
}
