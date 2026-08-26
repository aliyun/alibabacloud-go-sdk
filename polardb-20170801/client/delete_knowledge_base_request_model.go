// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *DeleteKnowledgeBaseRequest
	GetRegionId() *string
}

type DeleteKnowledgeBaseRequest struct {
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) SetRegionId(v string) *DeleteKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
