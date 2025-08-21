// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeFaqRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *DescribeFaqRequest
	GetKnowledgeId() *int64
}

type DescribeFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DescribeFaqRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaqRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeFaqRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DescribeFaqRequest) SetAgentKey(v string) *DescribeFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeFaqRequest) SetKnowledgeId(v int64) *DescribeFaqRequest {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeFaqRequest) Validate() error {
	return dara.Validate(s)
}
