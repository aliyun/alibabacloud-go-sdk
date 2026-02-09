// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTuringTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitTuringTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SubmitTuringTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *SubmitTuringTaskResponseBody
	GetRequestId() *string
	SetResult(v *SubmitTuringTaskResponseBodyResult) *SubmitTuringTaskResponseBody
	GetResult() *SubmitTuringTaskResponseBodyResult
	SetSuccess(v bool) *SubmitTuringTaskResponseBody
	GetSuccess() *bool
}

type SubmitTuringTaskResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 400
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 551FF252-6CFC-5DDA-9F84-9B07302385C2
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *SubmitTuringTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitTuringTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTuringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTuringTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitTuringTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitTuringTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTuringTaskResponseBody) GetResult() *SubmitTuringTaskResponseBodyResult {
	return s.Result
}

func (s *SubmitTuringTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitTuringTaskResponseBody) SetErrorCode(v string) *SubmitTuringTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitTuringTaskResponseBody) SetErrorMsg(v string) *SubmitTuringTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitTuringTaskResponseBody) SetRequestId(v string) *SubmitTuringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTuringTaskResponseBody) SetResult(v *SubmitTuringTaskResponseBodyResult) *SubmitTuringTaskResponseBody {
	s.Result = v
	return s
}

func (s *SubmitTuringTaskResponseBody) SetSuccess(v bool) *SubmitTuringTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitTuringTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTuringTaskResponseBodyResult struct {
	// example:
	//
	// 874890065171169280
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SubmitTuringTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitTuringTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SubmitTuringTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTuringTaskResponseBodyResult) SetTaskId(v string) *SubmitTuringTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *SubmitTuringTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
