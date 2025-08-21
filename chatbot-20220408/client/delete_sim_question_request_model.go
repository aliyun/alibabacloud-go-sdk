// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSimQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteSimQuestionRequest
	GetAgentKey() *string
	SetSimQuestionId(v int64) *DeleteSimQuestionRequest
	GetSimQuestionId() *int64
}

type DeleteSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000002788
	SimQuestionId *int64 `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
}

func (s DeleteSimQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteSimQuestionRequest) GetSimQuestionId() *int64 {
	return s.SimQuestionId
}

func (s *DeleteSimQuestionRequest) SetAgentKey(v string) *DeleteSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteSimQuestionRequest) SetSimQuestionId(v int64) *DeleteSimQuestionRequest {
	s.SimQuestionId = &v
	return s
}

func (s *DeleteSimQuestionRequest) Validate() error {
	return dara.Validate(s)
}
