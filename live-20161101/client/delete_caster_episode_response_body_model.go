// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterEpisodeResponseBody
	GetCasterId() *string
	SetEpisodeId(v string) *DeleteCasterEpisodeResponseBody
	GetEpisodeId() *string
	SetRequestId(v string) *DeleteCasterEpisodeResponseBody
	GetRequestId() *string
}

type DeleteCasterEpisodeResponseBody struct {
	// The ID of the production studio. You can specify the ID as a parameter in the request to modify the episode in the production studio.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the episode. You can specify the ID as a parameter in the request to modify the episode in the production studio.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf932738****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterEpisodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterEpisodeResponseBody) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *DeleteCasterEpisodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterEpisodeResponseBody) SetCasterId(v string) *DeleteCasterEpisodeResponseBody {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterEpisodeResponseBody) SetEpisodeId(v string) *DeleteCasterEpisodeResponseBody {
	s.EpisodeId = &v
	return s
}

func (s *DeleteCasterEpisodeResponseBody) SetRequestId(v string) *DeleteCasterEpisodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterEpisodeResponseBody) Validate() error {
	return dara.Validate(s)
}
