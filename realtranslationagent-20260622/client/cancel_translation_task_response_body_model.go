// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTranslationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelTranslationTaskResponseBody
	GetCode() *string
	SetData(v *CancelTranslationTaskResponseBodyData) *CancelTranslationTaskResponseBody
	GetData() *CancelTranslationTaskResponseBodyData
	SetMessage(v string) *CancelTranslationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelTranslationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelTranslationTaskResponseBody
	GetSuccess() *bool
}

type CancelTranslationTaskResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *CancelTranslationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F22CA868-601F-50D1-99B2-520491F25EB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelTranslationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTranslationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTranslationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelTranslationTaskResponseBody) GetData() *CancelTranslationTaskResponseBodyData {
	return s.Data
}

func (s *CancelTranslationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelTranslationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTranslationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelTranslationTaskResponseBody) SetCode(v string) *CancelTranslationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTranslationTaskResponseBody) SetData(v *CancelTranslationTaskResponseBodyData) *CancelTranslationTaskResponseBody {
	s.Data = v
	return s
}

func (s *CancelTranslationTaskResponseBody) SetMessage(v string) *CancelTranslationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelTranslationTaskResponseBody) SetRequestId(v string) *CancelTranslationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTranslationTaskResponseBody) SetSuccess(v bool) *CancelTranslationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelTranslationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelTranslationTaskResponseBodyData struct {
	// The task status.
	//
	// example:
	//
	// CANCEL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTranslationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelTranslationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelTranslationTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CancelTranslationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelTranslationTaskResponseBodyData) SetStatus(v string) *CancelTranslationTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CancelTranslationTaskResponseBodyData) SetTaskId(v string) *CancelTranslationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CancelTranslationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
