// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSimQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateSimQuestionRequest
	GetAgentKey() *string
	SetSimQuestionId(v int64) *UpdateSimQuestionRequest
	GetSimQuestionId() *int64
	SetTitle(v string) *UpdateSimQuestionRequest
	GetTitle() *string
}

type UpdateSimQuestionRequest struct {
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
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateSimQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateSimQuestionRequest) GetSimQuestionId() *int64 {
	return s.SimQuestionId
}

func (s *UpdateSimQuestionRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateSimQuestionRequest) SetAgentKey(v string) *UpdateSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateSimQuestionRequest) SetSimQuestionId(v int64) *UpdateSimQuestionRequest {
	s.SimQuestionId = &v
	return s
}

func (s *UpdateSimQuestionRequest) SetTitle(v string) *UpdateSimQuestionRequest {
	s.Title = &v
	return s
}

func (s *UpdateSimQuestionRequest) Validate() error {
	return dara.Validate(s)
}
