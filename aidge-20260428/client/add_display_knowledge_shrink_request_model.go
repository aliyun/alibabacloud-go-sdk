// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisplayKnowledgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *AddDisplayKnowledgeShrinkRequest
	GetRagId() *string
	SetRecordsShrink(v string) *AddDisplayKnowledgeShrinkRequest
	GetRecordsShrink() *string
}

type AddDisplayKnowledgeShrinkRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The records to add. A maximum of 500 records are supported.
	//
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
}

func (s AddDisplayKnowledgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeShrinkRequest) GetRagId() *string {
	return s.RagId
}

func (s *AddDisplayKnowledgeShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *AddDisplayKnowledgeShrinkRequest) SetRagId(v string) *AddDisplayKnowledgeShrinkRequest {
	s.RagId = &v
	return s
}

func (s *AddDisplayKnowledgeShrinkRequest) SetRecordsShrink(v string) *AddDisplayKnowledgeShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *AddDisplayKnowledgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
