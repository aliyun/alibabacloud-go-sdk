// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscriptionPhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTranscriptionPhrasesResponseBody
	GetCode() *string
	SetData(v *ListTranscriptionPhrasesResponseBodyData) *ListTranscriptionPhrasesResponseBody
	GetData() *ListTranscriptionPhrasesResponseBodyData
	SetMessage(v string) *ListTranscriptionPhrasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTranscriptionPhrasesResponseBody
	GetRequestId() *string
}

type ListTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListTranscriptionPhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTranscriptionPhrasesResponseBody) GetData() *ListTranscriptionPhrasesResponseBodyData {
	return s.Data
}

func (s *ListTranscriptionPhrasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTranscriptionPhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranscriptionPhrasesResponseBody) SetCode(v string) *ListTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetData(v *ListTranscriptionPhrasesResponseBodyData) *ListTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetMessage(v string) *ListTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetRequestId(v string) *ListTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// PHS.Exceed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The num of the phrase exceeds the upper limit.
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Phrases      []*ListTranscriptionPhrasesResponseBodyDataPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTranscriptionPhrasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTranscriptionPhrasesResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTranscriptionPhrasesResponseBodyData) GetPhrases() []*ListTranscriptionPhrasesResponseBodyDataPhrases {
	return s.Phrases
}

func (s *ListTranscriptionPhrasesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetPhrases(v []*ListTranscriptionPhrasesResponseBodyDataPhrases) *ListTranscriptionPhrasesResponseBodyData {
	s.Phrases = v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetStatus(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTranscriptionPhrasesResponseBodyDataPhrases struct {
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
}

func (s ListTranscriptionPhrasesResponseBodyDataPhrases) String() string {
	return dara.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBodyDataPhrases) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) GetDescription() *string {
	return s.Description
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) GetName() *string {
	return s.Name
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) GetPhraseId() *string {
	return s.PhraseId
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetDescription(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.Description = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetName(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.Name = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetPhraseId(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.PhraseId = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) Validate() error {
	return dara.Validate(s)
}
