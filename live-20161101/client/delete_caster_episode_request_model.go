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
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the episode. If the episode was added by calling the [AddCasterEpisode](https://help.aliyun.com/document_detail/2848068.html) operation, check the value of the response parameter EpisodeId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf932738****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
