// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListVocabResponseBodyData) *ListVocabResponseBody
	GetData() []*ListVocabResponseBodyData
	SetRequestId(v string) *ListVocabResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListVocabResponseBody
	GetSuccess() *string
}

type ListVocabResponseBody struct {
	Data []*ListVocabResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-*******F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVocabResponseBody) GoString() string {
	return s.String()
}

func (s *ListVocabResponseBody) GetData() []*ListVocabResponseBodyData {
	return s.Data
}

func (s *ListVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVocabResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListVocabResponseBody) SetData(v []*ListVocabResponseBodyData) *ListVocabResponseBody {
	s.Data = v
	return s
}

func (s *ListVocabResponseBody) SetRequestId(v string) *ListVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVocabResponseBody) SetSuccess(v string) *ListVocabResponseBody {
	s.Success = &v
	return s
}

func (s *ListVocabResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVocabResponseBodyData struct {
	// example:
	//
	// nls
	AudioModelCode *string `json:"audioModelCode,omitempty" xml:"audioModelCode,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// dv*****erverve
	VocabularyId   *string                                    `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	WordWeightList []*ListVocabResponseBodyDataWordWeightList `json:"wordWeightList,omitempty" xml:"wordWeightList,omitempty" type:"Repeated"`
}

func (s ListVocabResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVocabResponseBodyData) GetAudioModelCode() *string {
	return s.AudioModelCode
}

func (s *ListVocabResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListVocabResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListVocabResponseBodyData) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *ListVocabResponseBodyData) GetWordWeightList() []*ListVocabResponseBodyDataWordWeightList {
	return s.WordWeightList
}

func (s *ListVocabResponseBodyData) SetAudioModelCode(v string) *ListVocabResponseBodyData {
	s.AudioModelCode = &v
	return s
}

func (s *ListVocabResponseBodyData) SetDescription(v string) *ListVocabResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListVocabResponseBodyData) SetName(v string) *ListVocabResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListVocabResponseBodyData) SetVocabularyId(v string) *ListVocabResponseBodyData {
	s.VocabularyId = &v
	return s
}

func (s *ListVocabResponseBodyData) SetWordWeightList(v []*ListVocabResponseBodyDataWordWeightList) *ListVocabResponseBodyData {
	s.WordWeightList = v
	return s
}

func (s *ListVocabResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListVocabResponseBodyDataWordWeightList struct {
	// example:
	//
	// 3
	Weight *int32  `json:"weight,omitempty" xml:"weight,omitempty"`
	Word   *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListVocabResponseBodyDataWordWeightList) String() string {
	return dara.Prettify(s)
}

func (s ListVocabResponseBodyDataWordWeightList) GoString() string {
	return s.String()
}

func (s *ListVocabResponseBodyDataWordWeightList) GetWeight() *int32 {
	return s.Weight
}

func (s *ListVocabResponseBodyDataWordWeightList) GetWord() *string {
	return s.Word
}

func (s *ListVocabResponseBodyDataWordWeightList) SetWeight(v int32) *ListVocabResponseBodyDataWordWeightList {
	s.Weight = &v
	return s
}

func (s *ListVocabResponseBodyDataWordWeightList) SetWord(v string) *ListVocabResponseBodyDataWordWeightList {
	s.Word = &v
	return s
}

func (s *ListVocabResponseBodyDataWordWeightList) Validate() error {
	return dara.Validate(s)
}
