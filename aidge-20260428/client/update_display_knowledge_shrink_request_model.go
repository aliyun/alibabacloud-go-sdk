// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisplayKnowledgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRagId(v string) *UpdateDisplayKnowledgeShrinkRequest
	GetRagId() *string
	SetRecordsShrink(v string) *UpdateDisplayKnowledgeShrinkRequest
	GetRecordsShrink() *string
}

type UpdateDisplayKnowledgeShrinkRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The records to update. A maximum of 500 records can be specified.
	//
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
}

func (s UpdateDisplayKnowledgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeShrinkRequest) GetRagId() *string {
	return s.RagId
}

func (s *UpdateDisplayKnowledgeShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *UpdateDisplayKnowledgeShrinkRequest) SetRagId(v string) *UpdateDisplayKnowledgeShrinkRequest {
	s.RagId = &v
	return s
}

func (s *UpdateDisplayKnowledgeShrinkRequest) SetRecordsShrink(v string) *UpdateDisplayKnowledgeShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *UpdateDisplayKnowledgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
