// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAdHocTaskResultResponseBody
	GetCode() *string
	SetExecuteResult(v *GetAdHocTaskResultResponseBodyExecuteResult) *GetAdHocTaskResultResponseBody
	GetExecuteResult() *GetAdHocTaskResultResponseBodyExecuteResult
	SetHttpStatusCode(v int32) *GetAdHocTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAdHocTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAdHocTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAdHocTaskResultResponseBody
	GetSuccess() *bool
}

type GetAdHocTaskResultResponseBody struct {
	// example:
	//
	// OK
	Code          *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ExecuteResult *GetAdHocTaskResultResponseBodyExecuteResult `json:"ExecuteResult,omitempty" xml:"ExecuteResult,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAdHocTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAdHocTaskResultResponseBody) GetExecuteResult() *GetAdHocTaskResultResponseBodyExecuteResult {
	return s.ExecuteResult
}

func (s *GetAdHocTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAdHocTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAdHocTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdHocTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAdHocTaskResultResponseBody) SetCode(v string) *GetAdHocTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdHocTaskResultResponseBody) SetExecuteResult(v *GetAdHocTaskResultResponseBodyExecuteResult) *GetAdHocTaskResultResponseBody {
	s.ExecuteResult = v
	return s
}

func (s *GetAdHocTaskResultResponseBody) SetHttpStatusCode(v int32) *GetAdHocTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAdHocTaskResultResponseBody) SetMessage(v string) *GetAdHocTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdHocTaskResultResponseBody) SetRequestId(v string) *GetAdHocTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdHocTaskResultResponseBody) SetSuccess(v bool) *GetAdHocTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAdHocTaskResultResponseBody) Validate() error {
	if s.ExecuteResult != nil {
		if err := s.ExecuteResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAdHocTaskResultResponseBodyExecuteResult struct {
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// t_2242892326444990464_20210125_2242892326444990465
	ScheduleTaskId *string `json:"ScheduleTaskId,omitempty" xml:"ScheduleTaskId,omitempty"`
	// example:
	//
	// MaxCompute_SQL_300000843_1611548758327
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAdHocTaskResultResponseBodyExecuteResult) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskResultResponseBodyExecuteResult) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) GetResult() *string {
	return s.Result
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) GetScheduleTaskId() *string {
	return s.ScheduleTaskId
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) SetResult(v string) *GetAdHocTaskResultResponseBodyExecuteResult {
	s.Result = &v
	return s
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) SetScheduleTaskId(v string) *GetAdHocTaskResultResponseBodyExecuteResult {
	s.ScheduleTaskId = &v
	return s
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) SetTaskId(v string) *GetAdHocTaskResultResponseBodyExecuteResult {
	s.TaskId = &v
	return s
}

func (s *GetAdHocTaskResultResponseBodyExecuteResult) Validate() error {
	return dara.Validate(s)
}
