// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePracticeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BatchDeletePracticeTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BatchDeletePracticeTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *BatchDeletePracticeTaskResponseBody
	GetRequestId() *string
	SetResults(v []*BatchDeletePracticeTaskResponseBodyResults) *BatchDeletePracticeTaskResponseBody
	GetResults() []*BatchDeletePracticeTaskResponseBodyResults
	SetSuccess(v bool) *BatchDeletePracticeTaskResponseBody
	GetSuccess() *bool
}

type BatchDeletePracticeTaskResponseBody struct {
	ErrorCode    *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results      []*BatchDeletePracticeTaskResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Success      *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchDeletePracticeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePracticeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeletePracticeTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchDeletePracticeTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchDeletePracticeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeletePracticeTaskResponseBody) GetResults() []*BatchDeletePracticeTaskResponseBodyResults {
	return s.Results
}

func (s *BatchDeletePracticeTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeletePracticeTaskResponseBody) SetErrorCode(v string) *BatchDeletePracticeTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBody) SetErrorMessage(v string) *BatchDeletePracticeTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBody) SetRequestId(v string) *BatchDeletePracticeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBody) SetResults(v []*BatchDeletePracticeTaskResponseBodyResults) *BatchDeletePracticeTaskResponseBody {
	s.Results = v
	return s
}

func (s *BatchDeletePracticeTaskResponseBody) SetSuccess(v bool) *BatchDeletePracticeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchDeletePracticeTaskResponseBodyResults struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskId  *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchDeletePracticeTaskResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePracticeTaskResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchDeletePracticeTaskResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *BatchDeletePracticeTaskResponseBodyResults) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchDeletePracticeTaskResponseBodyResults) SetMessage(v string) *BatchDeletePracticeTaskResponseBodyResults {
	s.Message = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBodyResults) SetTaskId(v string) *BatchDeletePracticeTaskResponseBodyResults {
	s.TaskId = &v
	return s
}

func (s *BatchDeletePracticeTaskResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
