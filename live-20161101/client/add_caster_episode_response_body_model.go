// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEpisodeId(v string) *AddCasterEpisodeResponseBody
	GetEpisodeId() *string
	SetRequestId(v string) *AddCasterEpisodeResponseBody
	GetRequestId() *string
}

type AddCasterEpisodeResponseBody struct {
	// The ID of the episode. You can use the ID as a request parameter in the API operation that is used to query the information about the episode list, modify the configurations of the episode, and delete the episode.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasterEpisodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeResponseBody) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *AddCasterEpisodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterEpisodeResponseBody) SetEpisodeId(v string) *AddCasterEpisodeResponseBody {
	s.EpisodeId = &v
	return s
}

func (s *AddCasterEpisodeResponseBody) SetRequestId(v string) *AddCasterEpisodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeResponseBody) Validate() error {
	return dara.Validate(s)
}
