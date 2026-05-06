// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ImportVocabularyRequest
	GetBusinessUnitId() *string
	SetFileKey(v string) *ImportVocabularyRequest
	GetFileKey() *string
}

type ImportVocabularyRequest struct {
	// example:
	//
	// llm-zop7ukgtksltamo4
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	FileKey        *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
}

func (s ImportVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportVocabularyRequest) GoString() string {
	return s.String()
}

func (s *ImportVocabularyRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ImportVocabularyRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *ImportVocabularyRequest) SetBusinessUnitId(v string) *ImportVocabularyRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ImportVocabularyRequest) SetFileKey(v string) *ImportVocabularyRequest {
	s.FileKey = &v
	return s
}

func (s *ImportVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
