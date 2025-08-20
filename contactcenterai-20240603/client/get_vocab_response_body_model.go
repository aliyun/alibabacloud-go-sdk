// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetVocabResponseBodyData) *GetVocabResponseBody
	GetData() *GetVocabResponseBodyData
	SetRequestId(v string) *GetVocabResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetVocabResponseBody
	GetSuccess() *string
}

type GetVocabResponseBody struct {
	Data *GetVocabResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-*******F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVocabResponseBody) GoString() string {
	return s.String()
}

func (s *GetVocabResponseBody) GetData() *GetVocabResponseBodyData {
	return s.Data
}

func (s *GetVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVocabResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetVocabResponseBody) SetData(v *GetVocabResponseBodyData) *GetVocabResponseBody {
	s.Data = v
	return s
}

func (s *GetVocabResponseBody) SetRequestId(v string) *GetVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVocabResponseBody) SetSuccess(v string) *GetVocabResponseBody {
	s.Success = &v
	return s
}

func (s *GetVocabResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVocabResponseBodyData struct {
	// example:
	//
	// nls
	AudioModelCode *string `json:"audioModelCode,omitempty" xml:"audioModelCode,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rrbe***jrvrdd
	VocabularyId   *string                                   `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	WordWeightList []*GetVocabResponseBodyDataWordWeightList `json:"wordWeightList,omitempty" xml:"wordWeightList,omitempty" type:"Repeated"`
}

func (s GetVocabResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVocabResponseBodyData) GetAudioModelCode() *string {
	return s.AudioModelCode
}

func (s *GetVocabResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetVocabResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetVocabResponseBodyData) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetVocabResponseBodyData) GetWordWeightList() []*GetVocabResponseBodyDataWordWeightList {
	return s.WordWeightList
}

func (s *GetVocabResponseBodyData) SetAudioModelCode(v string) *GetVocabResponseBodyData {
	s.AudioModelCode = &v
	return s
}

func (s *GetVocabResponseBodyData) SetDescription(v string) *GetVocabResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetVocabResponseBodyData) SetName(v string) *GetVocabResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVocabResponseBodyData) SetVocabularyId(v string) *GetVocabResponseBodyData {
	s.VocabularyId = &v
	return s
}

func (s *GetVocabResponseBodyData) SetWordWeightList(v []*GetVocabResponseBodyDataWordWeightList) *GetVocabResponseBodyData {
	s.WordWeightList = v
	return s
}

func (s *GetVocabResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetVocabResponseBodyDataWordWeightList struct {
	// example:
	//
	// 1
	Weight *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
	Word   *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s GetVocabResponseBodyDataWordWeightList) String() string {
	return dara.Prettify(s)
}

func (s GetVocabResponseBodyDataWordWeightList) GoString() string {
	return s.String()
}

func (s *GetVocabResponseBodyDataWordWeightList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetVocabResponseBodyDataWordWeightList) GetWord() *string {
	return s.Word
}

func (s *GetVocabResponseBodyDataWordWeightList) SetWeight(v int32) *GetVocabResponseBodyDataWordWeightList {
	s.Weight = &v
	return s
}

func (s *GetVocabResponseBodyDataWordWeightList) SetWord(v string) *GetVocabResponseBodyDataWordWeightList {
	s.Word = &v
	return s
}

func (s *GetVocabResponseBodyDataWordWeightList) Validate() error {
	return dara.Validate(s)
}
