// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBusinessUnitId(v string) *ExportVocabularyRequest
  GetBusinessUnitId() *string 
  SetVocabularyIds(v []*string) *ExportVocabularyRequest
  GetVocabularyIds() []*string 
}

type ExportVocabularyRequest struct {
  // example:
  // 
  // llm-c11iig67g863rih8
  BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
  VocabularyIds []*string `json:"VocabularyIds,omitempty" xml:"VocabularyIds,omitempty" type:"Repeated"`
}

func (s ExportVocabularyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyRequest) GoString() string {
  return s.String()
}

func (s *ExportVocabularyRequest) GetBusinessUnitId() *string  {
  return s.BusinessUnitId
}

func (s *ExportVocabularyRequest) GetVocabularyIds() []*string  {
  return s.VocabularyIds
}

func (s *ExportVocabularyRequest) SetBusinessUnitId(v string) *ExportVocabularyRequest {
  s.BusinessUnitId = &v
  return s
}

func (s *ExportVocabularyRequest) SetVocabularyIds(v []*string) *ExportVocabularyRequest {
  s.VocabularyIds = v
  return s
}

func (s *ExportVocabularyRequest) Validate() error {
  return dara.Validate(s)
}

