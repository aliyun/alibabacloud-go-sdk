// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleLearningResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetStyleLearningResultRequest
	GetAgentKey() *string
	SetId(v int64) *GetStyleLearningResultRequest
	GetId() *int64
}

type GetStyleLearningResultRequest struct {
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
	// 39
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetStyleLearningResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStyleLearningResultRequest) GoString() string {
	return s.String()
}

func (s *GetStyleLearningResultRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetStyleLearningResultRequest) GetId() *int64 {
	return s.Id
}

func (s *GetStyleLearningResultRequest) SetAgentKey(v string) *GetStyleLearningResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetStyleLearningResultRequest) SetId(v int64) *GetStyleLearningResultRequest {
	s.Id = &v
	return s
}

func (s *GetStyleLearningResultRequest) Validate() error {
	return dara.Validate(s)
}
