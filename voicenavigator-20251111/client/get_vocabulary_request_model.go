// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetVocabularyRequest
	GetInstanceId() *string
	SetVocabularyId(v string) *GetVocabularyRequest
	GetVocabularyId() *string
}

type GetVocabularyRequest struct {
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *GetVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVocabularyRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetVocabularyRequest) SetInstanceId(v string) *GetVocabularyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVocabularyRequest) SetVocabularyId(v string) *GetVocabularyRequest {
	s.VocabularyId = &v
	return s
}

func (s *GetVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
