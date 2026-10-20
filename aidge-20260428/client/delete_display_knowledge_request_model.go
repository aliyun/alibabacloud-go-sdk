// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisplayKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *DeleteDisplayKnowledgeRequest
	GetRagId() *string
	SetRecordIds(v []*int64) *DeleteDisplayKnowledgeRequest
	GetRecordIds() []*int64
}

type DeleteDisplayKnowledgeRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_x x x
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The IDs of the records to delete. You can specify up to 500 record IDs.
	//
	// This parameter is required.
	RecordIds []*int64 `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
}

func (s DeleteDisplayKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeRequest) GetRagId() *string {
	return s.RagId
}

func (s *DeleteDisplayKnowledgeRequest) GetRecordIds() []*int64 {
	return s.RecordIds
}

func (s *DeleteDisplayKnowledgeRequest) SetRagId(v string) *DeleteDisplayKnowledgeRequest {
	s.RagId = &v
	return s
}

func (s *DeleteDisplayKnowledgeRequest) SetRecordIds(v []*int64) *DeleteDisplayKnowledgeRequest {
	s.RecordIds = v
	return s
}

func (s *DeleteDisplayKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
