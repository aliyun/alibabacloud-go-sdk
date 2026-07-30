// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTuringTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CheckTuringTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CheckTuringTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CheckTuringTaskResponseBody
	GetRequestId() *string
	SetResult(v *CheckTuringTaskResponseBodyResult) *CheckTuringTaskResponseBody
	GetResult() *CheckTuringTaskResponseBodyResult
	SetSuccess(v bool) *CheckTuringTaskResponseBody
	GetSuccess() *bool
}

type CheckTuringTaskResponseBody struct {
	ErrorCode *string                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                            `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CheckTuringTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckTuringTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTuringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTuringTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckTuringTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CheckTuringTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTuringTaskResponseBody) GetResult() *CheckTuringTaskResponseBodyResult {
	return s.Result
}

func (s *CheckTuringTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckTuringTaskResponseBody) SetErrorCode(v string) *CheckTuringTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckTuringTaskResponseBody) SetErrorMsg(v string) *CheckTuringTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckTuringTaskResponseBody) SetRequestId(v string) *CheckTuringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTuringTaskResponseBody) SetResult(v *CheckTuringTaskResponseBodyResult) *CheckTuringTaskResponseBody {
	s.Result = v
	return s
}

func (s *CheckTuringTaskResponseBody) SetSuccess(v bool) *CheckTuringTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CheckTuringTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckTuringTaskResponseBodyResult struct {
	FailBizCode *string `json:"failBizCode,omitempty" xml:"failBizCode,omitempty"`
	FailCode    *string `json:"failCode,omitempty" xml:"failCode,omitempty"`
	FailMsg     *string `json:"failMsg,omitempty" xml:"failMsg,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId      *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	VideoUrl    *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s CheckTuringTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CheckTuringTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckTuringTaskResponseBodyResult) GetFailBizCode() *string {
	return s.FailBizCode
}

func (s *CheckTuringTaskResponseBodyResult) GetFailCode() *string {
	return s.FailCode
}

func (s *CheckTuringTaskResponseBodyResult) GetFailMsg() *string {
	return s.FailMsg
}

func (s *CheckTuringTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CheckTuringTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CheckTuringTaskResponseBodyResult) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *CheckTuringTaskResponseBodyResult) SetFailBizCode(v string) *CheckTuringTaskResponseBodyResult {
	s.FailBizCode = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) SetFailCode(v string) *CheckTuringTaskResponseBodyResult {
	s.FailCode = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) SetFailMsg(v string) *CheckTuringTaskResponseBodyResult {
	s.FailMsg = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) SetStatus(v string) *CheckTuringTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) SetTaskId(v string) *CheckTuringTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) SetVideoUrl(v string) *CheckTuringTaskResponseBodyResult {
	s.VideoUrl = &v
	return s
}

func (s *CheckTuringTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
