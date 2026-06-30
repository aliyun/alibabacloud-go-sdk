// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranslationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitTranslationTaskResponseBody
	GetCode() *string
	SetData(v *SubmitTranslationTaskResponseBodyData) *SubmitTranslationTaskResponseBody
	GetData() *SubmitTranslationTaskResponseBodyData
	SetMessage(v string) *SubmitTranslationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTranslationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitTranslationTaskResponseBody
	GetSuccess() *bool
}

type SubmitTranslationTaskResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitTranslationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitTranslationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitTranslationTaskResponseBody) GetData() *SubmitTranslationTaskResponseBodyData {
	return s.Data
}

func (s *SubmitTranslationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTranslationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTranslationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitTranslationTaskResponseBody) SetCode(v string) *SubmitTranslationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTranslationTaskResponseBody) SetData(v *SubmitTranslationTaskResponseBodyData) *SubmitTranslationTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTranslationTaskResponseBody) SetMessage(v string) *SubmitTranslationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTranslationTaskResponseBody) SetRequestId(v string) *SubmitTranslationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTranslationTaskResponseBody) SetSuccess(v bool) *SubmitTranslationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitTranslationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranslationTaskResponseBodyData struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitTranslationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitTranslationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTranslationTaskResponseBodyData) SetStatus(v string) *SubmitTranslationTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitTranslationTaskResponseBodyData) SetTaskId(v string) *SubmitTranslationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitTranslationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
