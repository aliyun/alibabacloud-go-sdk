// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateVocabResponseBodyData) *CreateVocabResponseBody
	GetData() *CreateVocabResponseBodyData
	SetRequestId(v string) *CreateVocabResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateVocabResponseBody
	GetSuccess() *string
}

type CreateVocabResponseBody struct {
	Data *CreateVocabResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-*******F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVocabResponseBody) GetData() *CreateVocabResponseBodyData {
	return s.Data
}

func (s *CreateVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVocabResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateVocabResponseBody) SetData(v *CreateVocabResponseBodyData) *CreateVocabResponseBody {
	s.Data = v
	return s
}

func (s *CreateVocabResponseBody) SetRequestId(v string) *CreateVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVocabResponseBody) SetSuccess(v string) *CreateVocabResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVocabResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateVocabResponseBodyData struct {
	// example:
	//
	// f3d82*******7
	VocabularyId *string `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
}

func (s CreateVocabResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVocabResponseBodyData) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *CreateVocabResponseBodyData) SetVocabularyId(v string) *CreateVocabResponseBodyData {
	s.VocabularyId = &v
	return s
}

func (s *CreateVocabResponseBodyData) Validate() error {
	return dara.Validate(s)
}
