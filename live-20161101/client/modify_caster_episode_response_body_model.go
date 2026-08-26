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
	// The ID of the production studio. You can use this ID to perform operations on the scenario. These operations include copying, updating, querying, starting, and stopping the scenario. You can also use the ID to query the audio configuration of the scenario.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the episode. You can use this ID to manage the episode list and its items. For the episode list, you can perform operations such as query, edit, delete, start, and stop. For episode items, you can perform operations such as create, add, delete, and query.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938623****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The ID of the request.
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
