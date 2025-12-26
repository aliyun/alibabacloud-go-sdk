// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseFileChunks(v []*KnowledgeBaseFileChunk) *RetrieveKnowledgeBaseResponseBody
	GetKnowledgeBaseFileChunks() []*KnowledgeBaseFileChunk
}

type RetrieveKnowledgeBaseResponseBody struct {
	KnowledgeBaseFileChunks []*KnowledgeBaseFileChunk `json:"KnowledgeBaseFileChunks,omitempty" xml:"KnowledgeBaseFileChunks,omitempty" type:"Repeated"`
}

func (s RetrieveKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponseBody) GetKnowledgeBaseFileChunks() []*KnowledgeBaseFileChunk {
	return s.KnowledgeBaseFileChunks
}

func (s *RetrieveKnowledgeBaseResponseBody) SetKnowledgeBaseFileChunks(v []*KnowledgeBaseFileChunk) *RetrieveKnowledgeBaseResponseBody {
	s.KnowledgeBaseFileChunks = v
	return s
}

func (s *RetrieveKnowledgeBaseResponseBody) Validate() error {
	if s.KnowledgeBaseFileChunks != nil {
		for _, item := range s.KnowledgeBaseFileChunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
