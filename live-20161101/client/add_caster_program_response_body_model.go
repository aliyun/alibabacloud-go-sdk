// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEpisodeIds(v *AddCasterProgramResponseBodyEpisodeIds) *AddCasterProgramResponseBody
	GetEpisodeIds() *AddCasterProgramResponseBodyEpisodeIds
	SetRequestId(v string) *AddCasterProgramResponseBody
	GetRequestId() *string
}

type AddCasterProgramResponseBody struct {
	EpisodeIds *AddCasterProgramResponseBodyEpisodeIds `json:"EpisodeIds,omitempty" xml:"EpisodeIds,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasterProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponseBody) GetEpisodeIds() *AddCasterProgramResponseBodyEpisodeIds {
	return s.EpisodeIds
}

func (s *AddCasterProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterProgramResponseBody) SetEpisodeIds(v *AddCasterProgramResponseBodyEpisodeIds) *AddCasterProgramResponseBody {
	s.EpisodeIds = v
	return s
}

func (s *AddCasterProgramResponseBody) SetRequestId(v string) *AddCasterProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterProgramResponseBody) Validate() error {
	if s.EpisodeIds != nil {
		if err := s.EpisodeIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCasterProgramResponseBodyEpisodeIds struct {
	EpisodeId []*AddCasterProgramResponseBodyEpisodeIdsEpisodeId `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" type:"Repeated"`
}

func (s AddCasterProgramResponseBodyEpisodeIds) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramResponseBodyEpisodeIds) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponseBodyEpisodeIds) GetEpisodeId() []*AddCasterProgramResponseBodyEpisodeIdsEpisodeId {
	return s.EpisodeId
}

func (s *AddCasterProgramResponseBodyEpisodeIds) SetEpisodeId(v []*AddCasterProgramResponseBodyEpisodeIdsEpisodeId) *AddCasterProgramResponseBodyEpisodeIds {
	s.EpisodeId = v
	return s
}

func (s *AddCasterProgramResponseBodyEpisodeIds) Validate() error {
	if s.EpisodeId != nil {
		for _, item := range s.EpisodeId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddCasterProgramResponseBodyEpisodeIdsEpisodeId struct {
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
}

func (s AddCasterProgramResponseBodyEpisodeIdsEpisodeId) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramResponseBodyEpisodeIdsEpisodeId) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponseBodyEpisodeIdsEpisodeId) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *AddCasterProgramResponseBodyEpisodeIdsEpisodeId) SetEpisodeId(v string) *AddCasterProgramResponseBodyEpisodeIdsEpisodeId {
	s.EpisodeId = &v
	return s
}

func (s *AddCasterProgramResponseBodyEpisodeIdsEpisodeId) Validate() error {
	return dara.Validate(s)
}
