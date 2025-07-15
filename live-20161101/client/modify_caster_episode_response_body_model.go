// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterEpisodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyCasterEpisodeResponseBody
	GetCasterId() *string
	SetEpisodeId(v string) *ModifyCasterEpisodeResponseBody
	GetEpisodeId() *string
	SetRequestId(v string) *ModifyCasterEpisodeResponseBody
	GetRequestId() *string
}

type ModifyCasterEpisodeResponseBody struct {
	// The ID of the production studio. You can use the ID as a request parameter in the API operation that is used to copy the configurations of a scene, update the configurations of a scene, query the scenes of a production studio, query the audio configurations of a scene, start a scene in the production studio, or stop a scene in the production studio.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the episode. You can use the ID as a request parameter in the API operation that is used to delete an episode list from a production studio, add episodes to an episode list, remove episodes from an episode list, query the information about episodes in an episode list, update episodes in an episode list, delete an episode list, query the information about an episode list, start an episode list, or stop an episode list.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938623****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCasterEpisodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterEpisodeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCasterEpisodeResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterEpisodeResponseBody) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *ModifyCasterEpisodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCasterEpisodeResponseBody) SetCasterId(v string) *ModifyCasterEpisodeResponseBody {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterEpisodeResponseBody) SetEpisodeId(v string) *ModifyCasterEpisodeResponseBody {
	s.EpisodeId = &v
	return s
}

func (s *ModifyCasterEpisodeResponseBody) SetRequestId(v string) *ModifyCasterEpisodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterEpisodeResponseBody) Validate() error {
	return dara.Validate(s)
}
