// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportVocabularyShrinkRequest
  GetInstanceId() *string 
  SetVocabularyIdsShrink(v string) *ExportVocabularyShrinkRequest
  GetVocabularyIdsShrink() *string 
}

type ExportVocabularyShrinkRequest struct {
  // example:
  // 
  // af81a389-91f0-4157-8d82-720edd02b66a
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  VocabularyIdsShrink *string `json:"VocabularyIds,omitempty" xml:"VocabularyIds,omitempty"`
}

func (s ExportVocabularyShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportVocabularyShrinkRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportVocabularyShrinkRequest) GetVocabularyIdsShrink() *string  {
  return s.VocabularyIdsShrink
}

func (s *ExportVocabularyShrinkRequest) SetInstanceId(v string) *ExportVocabularyShrinkRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportVocabularyShrinkRequest) SetVocabularyIdsShrink(v string) *ExportVocabularyShrinkRequest {
  s.VocabularyIdsShrink = &v
  return s
}

func (s *ExportVocabularyShrinkRequest) Validate() error {
  return dara.Validate(s)
}

