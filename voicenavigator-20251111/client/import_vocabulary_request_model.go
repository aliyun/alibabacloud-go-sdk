// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *ImportVocabularyRequest
	GetFileKey() *string
	SetInstanceId(v string) *ImportVocabularyRequest
	GetInstanceId() *string
}

type ImportVocabularyRequest struct {
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ImportVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportVocabularyRequest) GoString() string {
	return s.String()
}

func (s *ImportVocabularyRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *ImportVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportVocabularyRequest) SetFileKey(v string) *ImportVocabularyRequest {
	s.FileKey = &v
	return s
}

func (s *ImportVocabularyRequest) SetInstanceId(v string) *ImportVocabularyRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
