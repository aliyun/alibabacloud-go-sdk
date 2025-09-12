// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscriptionPhrasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTranscriptionPhrasesResponseBody
	GetCode() *string
	SetData(v *UpdateTranscriptionPhrasesResponseBodyData) *UpdateTranscriptionPhrasesResponseBody
	GetData() *UpdateTranscriptionPhrasesResponseBodyData
	SetMessage(v string) *UpdateTranscriptionPhrasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTranscriptionPhrasesResponseBody
	GetRequestId() *string
}

type UpdateTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateTranscriptionPhrasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTranscriptionPhrasesResponseBody) GetData() *UpdateTranscriptionPhrasesResponseBodyData {
	return s.Data
}

func (s *UpdateTranscriptionPhrasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTranscriptionPhrasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetCode(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetData(v *UpdateTranscriptionPhrasesResponseBodyData) *UpdateTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetMessage(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetRequestId(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTranscriptionPhrasesResponseBodyData struct {
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

func (s UpdateTranscriptionPhrasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetStatus(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
