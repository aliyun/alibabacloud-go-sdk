// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBusinessUnitId(v string) *ExportVocabularyShrinkRequest
  GetBusinessUnitId() *string 
  SetVocabularyIdsShrink(v string) *ExportVocabularyShrinkRequest
  GetVocabularyIdsShrink() *string 
}

type ExportVocabularyShrinkRequest struct {
  // example:
  // 
  // llm-c11iig67g863rih8
  BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
  VocabularyIdsShrink *string `json:"VocabularyIds,omitempty" xml:"VocabularyIds,omitempty"`
}

func (s ExportVocabularyShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportVocabularyShrinkRequest) GetBusinessUnitId() *string  {
  return s.BusinessUnitId
}

func (s *ExportVocabularyShrinkRequest) GetVocabularyIdsShrink() *string  {
  return s.VocabularyIdsShrink
}

func (s *ExportVocabularyShrinkRequest) SetBusinessUnitId(v string) *ExportVocabularyShrinkRequest {
  s.BusinessUnitId = &v
  return s
}

func (s *ExportVocabularyShrinkRequest) SetVocabularyIdsShrink(v string) *ExportVocabularyShrinkRequest {
  s.VocabularyIdsShrink = &v
  return s
}

func (s *ExportVocabularyShrinkRequest) Validate() error {
  return dara.Validate(s)
}

