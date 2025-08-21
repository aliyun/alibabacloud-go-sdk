// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateConnQuestionRequest
	GetAgentKey() *string
	SetConnQuestionId(v int64) *UpdateConnQuestionRequest
	GetConnQuestionId() *int64
	SetOutlineId(v int64) *UpdateConnQuestionRequest
	GetOutlineId() *int64
}

type UpdateConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000000295
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 877397683
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s UpdateConnQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateConnQuestionRequest) GetConnQuestionId() *int64 {
	return s.ConnQuestionId
}

func (s *UpdateConnQuestionRequest) GetOutlineId() *int64 {
	return s.OutlineId
}

func (s *UpdateConnQuestionRequest) SetAgentKey(v string) *UpdateConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateConnQuestionRequest) SetConnQuestionId(v int64) *UpdateConnQuestionRequest {
	s.ConnQuestionId = &v
	return s
}

func (s *UpdateConnQuestionRequest) SetOutlineId(v int64) *UpdateConnQuestionRequest {
	s.OutlineId = &v
	return s
}

func (s *UpdateConnQuestionRequest) Validate() error {
	return dara.Validate(s)
}
