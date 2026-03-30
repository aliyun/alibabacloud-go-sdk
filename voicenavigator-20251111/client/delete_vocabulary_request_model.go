// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteVocabularyRequest
	GetInstanceId() *string
	SetVocabularyId(v string) *DeleteVocabularyRequest
	GetVocabularyId() *string
}

type DeleteVocabularyRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s DeleteVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabularyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVocabularyRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *DeleteVocabularyRequest) SetInstanceId(v string) *DeleteVocabularyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVocabularyRequest) SetVocabularyId(v string) *DeleteVocabularyRequest {
	s.VocabularyId = &v
	return s
}

func (s *DeleteVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
