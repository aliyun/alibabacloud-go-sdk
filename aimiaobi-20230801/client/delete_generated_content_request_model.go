// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneratedContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteGeneratedContentRequest
	GetAgentKey() *string
	SetId(v int64) *DeleteGeneratedContentRequest
	GetId() *int64
	SetRegionId(v string) *DeleteGeneratedContentRequest
	GetRegionId() *string
}

type DeleteGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGeneratedContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteGeneratedContentRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteGeneratedContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGeneratedContentRequest) SetAgentKey(v string) *DeleteGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteGeneratedContentRequest) SetId(v int64) *DeleteGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *DeleteGeneratedContentRequest) SetRegionId(v string) *DeleteGeneratedContentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGeneratedContentRequest) Validate() error {
	return dara.Validate(s)
}
