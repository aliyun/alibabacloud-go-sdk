// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteCasterEpisodeGroupRequest
	GetOwnerId() *int64
	SetProgramId(v string) *DeleteCasterEpisodeGroupRequest
	GetProgramId() *string
	SetRegionId(v string) *DeleteCasterEpisodeGroupRequest
	GetRegionId() *string
}

type DeleteCasterEpisodeGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the episode list. If the episode list was added by calling the [AddCasterEpisodeGroup](https://help.aliyun.com/document_detail/2848071.html) operation, check the value of the response parameter ProgramId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf932738****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCasterEpisodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterEpisodeGroupRequest) GetProgramId() *string {
	return s.ProgramId
}

func (s *DeleteCasterEpisodeGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterEpisodeGroupRequest) SetOwnerId(v int64) *DeleteCasterEpisodeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterEpisodeGroupRequest) SetProgramId(v string) *DeleteCasterEpisodeGroupRequest {
	s.ProgramId = &v
	return s
}

func (s *DeleteCasterEpisodeGroupRequest) SetRegionId(v string) *DeleteCasterEpisodeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterEpisodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
