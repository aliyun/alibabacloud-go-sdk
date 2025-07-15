// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStyleLearningResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteStyleLearningResultRequest
	GetAgentKey() *string
	SetId(v int64) *DeleteStyleLearningResultRequest
	GetId() *int64
}

type DeleteStyleLearningResultRequest struct {
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
	// 20
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteStyleLearningResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStyleLearningResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteStyleLearningResultRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteStyleLearningResultRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteStyleLearningResultRequest) SetAgentKey(v string) *DeleteStyleLearningResultRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteStyleLearningResultRequest) SetId(v int64) *DeleteStyleLearningResultRequest {
	s.Id = &v
	return s
}

func (s *DeleteStyleLearningResultRequest) Validate() error {
	return dara.Validate(s)
}
