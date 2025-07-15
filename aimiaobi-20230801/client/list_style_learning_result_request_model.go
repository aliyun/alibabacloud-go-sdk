// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStyleLearningResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListStyleLearningResultRequest
	GetAgentKey() *string
	SetCurrent(v int32) *ListStyleLearningResultRequest
	GetCurrent() *int32
	SetSize(v int32) *ListStyleLearningResultRequest
	GetSize() *int32
}

type ListStyleLearningResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListStyleLearningResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStyleLearningResultRequest) GoString() string {
	return s.String()
}

func (s *ListStyleLearningResultRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListStyleLearningResultRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListStyleLearningResultRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListStyleLearningResultRequest) SetAgentKey(v string) *ListStyleLearningResultRequest {
	s.AgentKey = &v
	return s
}

func (s *ListStyleLearningResultRequest) SetCurrent(v int32) *ListStyleLearningResultRequest {
	s.Current = &v
	return s
}

func (s *ListStyleLearningResultRequest) SetSize(v int32) *ListStyleLearningResultRequest {
	s.Size = &v
	return s
}

func (s *ListStyleLearningResultRequest) Validate() error {
	return dara.Validate(s)
}
