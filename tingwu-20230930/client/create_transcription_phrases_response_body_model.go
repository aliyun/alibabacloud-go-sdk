// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTranscriptionPhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTranscriptionPhrasesResponseBody
	GetCode() *string
	SetData(v *CreateTranscriptionPhrasesResponseBodyData) *CreateTranscriptionPhrasesResponseBody
	GetData() *CreateTranscriptionPhrasesResponseBodyData
	SetMessage(v string) *CreateTranscriptionPhrasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTranscriptionPhrasesResponseBody
	GetRequestId() *string
}

type CreateTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateTranscriptionPhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTranscriptionPhrasesResponseBody) GetData() *CreateTranscriptionPhrasesResponseBodyData {
	return s.Data
}

func (s *CreateTranscriptionPhrasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTranscriptionPhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTranscriptionPhrasesResponseBody) SetCode(v string) *CreateTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetData(v *CreateTranscriptionPhrasesResponseBodyData) *CreateTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetMessage(v string) *CreateTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetRequestId(v string) *CreateTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateTranscriptionPhrasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTranscriptionPhrasesResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTranscriptionPhrasesResponseBodyData) GetPhraseId() *string {
	return s.PhraseId
}

func (s *CreateTranscriptionPhrasesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetPhraseId(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.PhraseId = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetStatus(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
