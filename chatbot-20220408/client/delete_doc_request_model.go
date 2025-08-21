// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteDocRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *DeleteDocRequest
	GetKnowledgeId() *int64
}

type DeleteDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteDocRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteDocRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DeleteDocRequest) SetAgentKey(v string) *DeleteDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDocRequest) SetKnowledgeId(v int64) *DeleteDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *DeleteDocRequest) Validate() error {
	return dara.Validate(s)
}
