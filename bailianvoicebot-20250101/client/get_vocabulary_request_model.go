// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *GetVocabularyRequest
	GetBusinessUnitId() *string
	SetVocabularyId(v string) *GetVocabularyRequest
	GetVocabularyId() *string
}

type GetVocabularyRequest struct {
	// example:
	//
	// llm-zop7ukgtksltamo4
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4061
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s GetVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVocabularyRequest) GoString() string {
	return s.String()
}

func (s *GetVocabularyRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *GetVocabularyRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetVocabularyRequest) SetBusinessUnitId(v string) *GetVocabularyRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *GetVocabularyRequest) SetVocabularyId(v string) *GetVocabularyRequest {
	s.VocabularyId = &v
	return s
}

func (s *GetVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
