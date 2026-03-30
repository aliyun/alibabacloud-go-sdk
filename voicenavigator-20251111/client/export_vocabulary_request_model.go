// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportVocabularyRequest
  GetInstanceId() *string 
  SetVocabularyIds(v []*string) *ExportVocabularyRequest
  GetVocabularyIds() []*string 
}

type ExportVocabularyRequest struct {
  // example:
  // 
  // af81a389-91f0-4157-8d82-720edd02b66a
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  VocabularyIds []*string `json:"VocabularyIds,omitempty" xml:"VocabularyIds,omitempty" type:"Repeated"`
}

func (s ExportVocabularyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyRequest) GoString() string {
  return s.String()
}

func (s *ExportVocabularyRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportVocabularyRequest) GetVocabularyIds() []*string  {
  return s.VocabularyIds
}

func (s *ExportVocabularyRequest) SetInstanceId(v string) *ExportVocabularyRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportVocabularyRequest) SetVocabularyIds(v []*string) *ExportVocabularyRequest {
  s.VocabularyIds = v
  return s
}

func (s *ExportVocabularyRequest) Validate() error {
  return dara.Validate(s)
}

