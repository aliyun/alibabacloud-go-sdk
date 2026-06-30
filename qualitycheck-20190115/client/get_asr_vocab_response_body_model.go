// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAsrVocabResponseBody
	GetCode() *string
	SetData(v *GetAsrVocabResponseBodyData) *GetAsrVocabResponseBody
	GetData() *GetAsrVocabResponseBodyData
	SetMessage(v string) *GetAsrVocabResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsrVocabResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsrVocabResponseBody
	GetSuccess() *bool
}

type GetAsrVocabResponseBody struct {
	// The result code. A value of **200*	- means success. Any other value means failure. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error details if the request fails. If successful, the value is **successful**.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Use this field to check the result:
	//
	// - **true*	- means success
	//
	// - false or **null*	- means failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAsrVocabResponseBody) GetData() *GetAsrVocabResponseBodyData {
	return s.Data
}

func (s *GetAsrVocabResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsrVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsrVocabResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsrVocabResponseBody) SetCode(v string) *GetAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetData(v *GetAsrVocabResponseBodyData) *GetAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrVocabResponseBody) SetMessage(v string) *GetAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetRequestId(v string) *GetAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetSuccess(v bool) *GetAsrVocabResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsrVocabResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAsrVocabResponseBodyData struct {
	// The ASR version.
	//
	// - 2 or **null**: V2 (Intelligent Speech Interaction ASR)
	//
	// - **3**: V3 (Paraformer ASR)
	//
	// example:
	//
	// 3
	AsrVersion *int32 `json:"AsrVersion,omitempty" xml:"AsrVersion,omitempty"`
	// The language model ID. This field appears only for V3.
	//
	// example:
	//
	// paraformer-8k-v2
	ModelCustomizationId *string `json:"ModelCustomizationId,omitempty" xml:"ModelCustomizationId,omitempty"`
	// The hotword group name.
	//
	// example:
	//
	// test
	Name  *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Words *GetAsrVocabResponseBodyDataWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Struct"`
}

func (s GetAsrVocabResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyData) GetAsrVersion() *int32 {
	return s.AsrVersion
}

func (s *GetAsrVocabResponseBodyData) GetModelCustomizationId() *string {
	return s.ModelCustomizationId
}

func (s *GetAsrVocabResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAsrVocabResponseBodyData) GetWords() *GetAsrVocabResponseBodyDataWords {
	return s.Words
}

func (s *GetAsrVocabResponseBodyData) SetAsrVersion(v int32) *GetAsrVocabResponseBodyData {
	s.AsrVersion = &v
	return s
}

func (s *GetAsrVocabResponseBodyData) SetModelCustomizationId(v string) *GetAsrVocabResponseBodyData {
	s.ModelCustomizationId = &v
	return s
}

func (s *GetAsrVocabResponseBodyData) SetName(v string) *GetAsrVocabResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAsrVocabResponseBodyData) SetWords(v *GetAsrVocabResponseBodyDataWords) *GetAsrVocabResponseBodyData {
	s.Words = v
	return s
}

func (s *GetAsrVocabResponseBodyData) Validate() error {
	if s.Words != nil {
		if err := s.Words.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAsrVocabResponseBodyDataWords struct {
	Word []*GetAsrVocabResponseBodyDataWordsWord `json:"Word,omitempty" xml:"Word,omitempty" type:"Repeated"`
}

func (s GetAsrVocabResponseBodyDataWords) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWords) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWords) GetWord() []*GetAsrVocabResponseBodyDataWordsWord {
	return s.Word
}

func (s *GetAsrVocabResponseBodyDataWords) SetWord(v []*GetAsrVocabResponseBodyDataWordsWord) *GetAsrVocabResponseBodyDataWords {
	s.Word = v
	return s
}

func (s *GetAsrVocabResponseBodyDataWords) Validate() error {
	if s.Word != nil {
		for _, item := range s.Word {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAsrVocabResponseBodyDataWordsWord struct {
	Weight *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Word   *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s GetAsrVocabResponseBodyDataWordsWord) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWordsWord) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWordsWord) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAsrVocabResponseBodyDataWordsWord) GetWord() *string {
	return s.Word
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWeight(v int32) *GetAsrVocabResponseBodyDataWordsWord {
	s.Weight = &v
	return s
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWord(v string) *GetAsrVocabResponseBodyDataWordsWord {
	s.Word = &v
	return s
}

func (s *GetAsrVocabResponseBodyDataWordsWord) Validate() error {
	return dara.Validate(s)
}
