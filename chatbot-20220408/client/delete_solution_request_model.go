// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteSolutionRequest
	GetAgentKey() *string
	SetSolutionId(v int64) *DeleteSolutionRequest
	GetSolutionId() *int64
}

type DeleteSolutionRequest struct {
	// The key of the business space. If this parameter is not specified, the default business space is accessed. You can obtain the key on the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The ID of the answer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001321580
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DeleteSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSolutionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSolutionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteSolutionRequest) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *DeleteSolutionRequest) SetAgentKey(v string) *DeleteSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteSolutionRequest) SetSolutionId(v int64) *DeleteSolutionRequest {
	s.SolutionId = &v
	return s
}

func (s *DeleteSolutionRequest) Validate() error {
	return dara.Validate(s)
}
