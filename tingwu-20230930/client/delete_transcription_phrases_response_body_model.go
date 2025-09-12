// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscriptionPhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteTranscriptionPhrasesResponseBodyData) *DeleteTranscriptionPhrasesResponseBody
	GetData() *DeleteTranscriptionPhrasesResponseBodyData
	SetErrorCode(v string) *DeleteTranscriptionPhrasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTranscriptionPhrasesResponseBody
	GetErrorMessage() *string
	SetStatus(v string) *DeleteTranscriptionPhrasesResponseBody
	GetStatus() *string
}

type DeleteTranscriptionPhrasesResponseBody struct {
	Data *DeleteTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteTranscriptionPhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTranscriptionPhrasesResponseBody) GetData() *DeleteTranscriptionPhrasesResponseBodyData {
	return s.Data
}

func (s *DeleteTranscriptionPhrasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTranscriptionPhrasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTranscriptionPhrasesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetData(v *DeleteTranscriptionPhrasesResponseBodyData) *DeleteTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetErrorCode(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetErrorMessage(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetStatus(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteTranscriptionPhrasesResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteTranscriptionPhrasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *DeleteTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *DeleteTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) SetStatus(v string) *DeleteTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
