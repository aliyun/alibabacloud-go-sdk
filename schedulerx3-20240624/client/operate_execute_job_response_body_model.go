// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateExecuteJobResponseBody
	GetCode() *int32
	SetData(v *OperateExecuteJobResponseBodyData) *OperateExecuteJobResponseBody
	GetData() *OperateExecuteJobResponseBodyData
	SetMessage(v string) *OperateExecuteJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateExecuteJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateExecuteJobResponseBody
	GetSuccess() *bool
}

type OperateExecuteJobResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *OperateExecuteJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6305893D-517D-5131-A767-644EDA81CEC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateExecuteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateExecuteJobResponseBody) GetData() *OperateExecuteJobResponseBodyData {
	return s.Data
}

func (s *OperateExecuteJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateExecuteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateExecuteJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateExecuteJobResponseBody) SetCode(v int32) *OperateExecuteJobResponseBody {
	s.Code = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetData(v *OperateExecuteJobResponseBodyData) *OperateExecuteJobResponseBody {
	s.Data = v
	return s
}

func (s *OperateExecuteJobResponseBody) SetMessage(v string) *OperateExecuteJobResponseBody {
	s.Message = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetRequestId(v string) *OperateExecuteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateExecuteJobResponseBody) SetSuccess(v bool) *OperateExecuteJobResponseBody {
	s.Success = &v
	return s
}

func (s *OperateExecuteJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OperateExecuteJobResponseBodyData struct {
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s OperateExecuteJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponseBodyData) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateExecuteJobResponseBodyData) SetJobExecutionId(v string) *OperateExecuteJobResponseBodyData {
	s.JobExecutionId = &v
	return s
}

func (s *OperateExecuteJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
