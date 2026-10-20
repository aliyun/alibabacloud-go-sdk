// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisplayKnowledgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *DeleteDisplayKnowledgeShrinkRequest
	GetRagId() *string
	SetRecordIdsShrink(v string) *DeleteDisplayKnowledgeShrinkRequest
	GetRecordIdsShrink() *string
}

type DeleteDisplayKnowledgeShrinkRequest struct {
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
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
}

func (s DeleteDisplayKnowledgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeShrinkRequest) GetRagId() *string {
	return s.RagId
}

func (s *DeleteDisplayKnowledgeShrinkRequest) GetRecordIdsShrink() *string {
	return s.RecordIdsShrink
}

func (s *DeleteDisplayKnowledgeShrinkRequest) SetRagId(v string) *DeleteDisplayKnowledgeShrinkRequest {
	s.RagId = &v
	return s
}

func (s *DeleteDisplayKnowledgeShrinkRequest) SetRecordIdsShrink(v string) *DeleteDisplayKnowledgeShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

func (s *DeleteDisplayKnowledgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
