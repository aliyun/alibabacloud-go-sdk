// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteConnQuestionRequest
	GetAgentKey() *string
	SetOutlineId(v int64) *DeleteConnQuestionRequest
	GetOutlineId() *int64
}

type DeleteConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 877397683
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s DeleteConnQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteConnQuestionRequest) GetOutlineId() *int64 {
	return s.OutlineId
}

func (s *DeleteConnQuestionRequest) SetAgentKey(v string) *DeleteConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteConnQuestionRequest) SetOutlineId(v int64) *DeleteConnQuestionRequest {
	s.OutlineId = &v
	return s
}

func (s *DeleteConnQuestionRequest) Validate() error {
	return dara.Validate(s)
}
