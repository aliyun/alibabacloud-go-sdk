// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscriptionPhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTranscriptionPhrasesResponseBody
	GetCode() *string
	SetData(v *GetTranscriptionPhrasesResponseBodyData) *GetTranscriptionPhrasesResponseBody
	GetData() *GetTranscriptionPhrasesResponseBodyData
	SetMessage(v string) *GetTranscriptionPhrasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranscriptionPhrasesResponseBody
	GetRequestId() *string
}

type GetTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTranscriptionPhrasesResponseBody) GetData() *GetTranscriptionPhrasesResponseBodyData {
	return s.Data
}

func (s *GetTranscriptionPhrasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranscriptionPhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscriptionPhrasesResponseBody) SetCode(v string) *GetTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetData(v *GetTranscriptionPhrasesResponseBodyData) *GetTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetMessage(v string) *GetTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetRequestId(v string) *GetTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Phrases      []*GetTranscriptionPhrasesResponseBodyDataPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTranscriptionPhrasesResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTranscriptionPhrasesResponseBodyData) GetPhrases() []*GetTranscriptionPhrasesResponseBodyDataPhrases {
	return s.Phrases
}

func (s *GetTranscriptionPhrasesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetPhrases(v []*GetTranscriptionPhrasesResponseBodyDataPhrases) *GetTranscriptionPhrasesResponseBodyData {
	s.Phrases = v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetStatus(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTranscriptionPhrasesResponseBodyDataPhrases struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// example:
	//
	// {"苹果":3,"西瓜":3}
	WordWeights *string `json:"WordWeights,omitempty" xml:"WordWeights,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBodyDataPhrases) String() string {
	return dara.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBodyDataPhrases) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) GetDescription() *string {
	return s.Description
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) GetName() *string {
	return s.Name
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) GetPhraseId() *string {
	return s.PhraseId
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) GetWordWeights() *string {
	return s.WordWeights
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetDescription(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.Description = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetName(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.Name = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetPhraseId(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.PhraseId = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetWordWeights(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.WordWeights = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) Validate() error {
	return dara.Validate(s)
}
