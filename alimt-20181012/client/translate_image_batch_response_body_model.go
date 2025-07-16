// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateImageBatchResponseBody
	GetCode() *int32
	SetData(v *TranslateImageBatchResponseBodyData) *TranslateImageBatchResponseBody
	GetData() *TranslateImageBatchResponseBodyData
	SetMessage(v string) *TranslateImageBatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateImageBatchResponseBody
	GetRequestId() *string
}

type TranslateImageBatchResponseBody struct {
	// example:
	//
	// 200
	Code *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateImageBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D774D33D-F1CB-5A2C-A787-E0A2179239CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateImageBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageBatchResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateImageBatchResponseBody) GetData() *TranslateImageBatchResponseBodyData {
	return s.Data
}

func (s *TranslateImageBatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateImageBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateImageBatchResponseBody) SetCode(v int32) *TranslateImageBatchResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateImageBatchResponseBody) SetData(v *TranslateImageBatchResponseBodyData) *TranslateImageBatchResponseBody {
	s.Data = v
	return s
}

func (s *TranslateImageBatchResponseBody) SetMessage(v string) *TranslateImageBatchResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateImageBatchResponseBody) SetRequestId(v string) *TranslateImageBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateImageBatchResponseBody) Validate() error {
	return dara.Validate(s)
}

type TranslateImageBatchResponseBodyData struct {
	// example:
	//
	// EEA28E6D-4828-5031-BD8C-8FF1B3216842
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TranslateImageBatchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *TranslateImageBatchResponseBodyData) SetTaskId(v string) *TranslateImageBatchResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *TranslateImageBatchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
