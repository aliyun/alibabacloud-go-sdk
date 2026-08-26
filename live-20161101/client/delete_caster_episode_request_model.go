// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterEpisodeRequest
	GetCasterId() *string
	SetEpisodeId(v string) *DeleteCasterEpisodeRequest
	GetEpisodeId() *string
	SetOwnerId(v int64) *DeleteCasterEpisodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterEpisodeRequest
	GetRegionId() *string
}

type DeleteCasterEpisodeRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value from the response.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The production studio name on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The episode ID. If you added the episode by calling the [AddCasterEpisode](https://help.aliyun.com/document_detail/2848068.html) operation, use the EpisodeId value from the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf932738****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCasterEpisodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterEpisodeRequest) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *DeleteCasterEpisodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterEpisodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterEpisodeRequest) SetCasterId(v string) *DeleteCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterEpisodeRequest) SetEpisodeId(v string) *DeleteCasterEpisodeRequest {
	s.EpisodeId = &v
	return s
}

func (s *DeleteCasterEpisodeRequest) SetOwnerId(v int64) *DeleteCasterEpisodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterEpisodeRequest) SetRegionId(v string) *DeleteCasterEpisodeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterEpisodeRequest) Validate() error {
	return dara.Validate(s)
}
