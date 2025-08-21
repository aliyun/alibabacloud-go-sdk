// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *RetryDocRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *RetryDocRequest
	GetKnowledgeId() *int64
}

type RetryDocRequest struct {
	// example:
	//
	// e2a20f74cd9042558002c0f7dc873739_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s RetryDocRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryDocRequest) GoString() string {
	return s.String()
}

func (s *RetryDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *RetryDocRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *RetryDocRequest) SetAgentKey(v string) *RetryDocRequest {
	s.AgentKey = &v
	return s
}

func (s *RetryDocRequest) SetKnowledgeId(v int64) *RetryDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *RetryDocRequest) Validate() error {
	return dara.Validate(s)
}
