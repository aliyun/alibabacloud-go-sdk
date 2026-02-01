// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAgentTaskResponseBody
	GetRequestId() *string
	SetCode(v string) *GetAgentTaskResponseBody
	GetCode() *string
	SetData(v *GetAgentTaskResponseBodyData) *GetAgentTaskResponseBody
	GetData() *GetAgentTaskResponseBodyData
	SetMessage(v string) *GetAgentTaskResponseBody
	GetMessage() *string
}

type GetAgentTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAgentTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentTaskResponseBody) GetData() *GetAgentTaskResponseBodyData {
	return s.Data
}

func (s *GetAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentTaskResponseBody) SetRequestId(v string) *GetAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentTaskResponseBody) SetCode(v string) *GetAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentTaskResponseBody) SetData(v *GetAgentTaskResponseBodyData) *GetAgentTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentTaskResponseBody) SetMessage(v string) *GetAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentTaskResponseBodyData struct {
	Jobs   []*GetAgentTaskResponseBodyDataJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	Status *string                             `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                             `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAgentTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBodyData) GetJobs() []*GetAgentTaskResponseBodyDataJobs {
	return s.Jobs
}

func (s *GetAgentTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAgentTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentTaskResponseBodyData) SetJobs(v []*GetAgentTaskResponseBodyDataJobs) *GetAgentTaskResponseBodyData {
	s.Jobs = v
	return s
}

func (s *GetAgentTaskResponseBodyData) SetStatus(v string) *GetAgentTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentTaskResponseBodyData) SetTaskId(v string) *GetAgentTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAgentTaskResponseBodyData) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentTaskResponseBodyDataJobs struct {
	Error        *string     `json:"error,omitempty" xml:"error,omitempty"`
	ErrorCode    *string     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Instance     *string     `json:"instance,omitempty" xml:"instance,omitempty"`
	Params       interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Region       *string     `json:"region,omitempty" xml:"region,omitempty"`
	Result       *string     `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAgentTaskResponseBodyDataJobs) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBodyDataJobs) GetError() *string {
	return s.Error
}

func (s *GetAgentTaskResponseBodyDataJobs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgentTaskResponseBodyDataJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgentTaskResponseBodyDataJobs) GetInstance() *string {
	return s.Instance
}

func (s *GetAgentTaskResponseBodyDataJobs) GetParams() interface{} {
	return s.Params
}

func (s *GetAgentTaskResponseBodyDataJobs) GetRegion() *string {
	return s.Region
}

func (s *GetAgentTaskResponseBodyDataJobs) GetResult() *string {
	return s.Result
}

func (s *GetAgentTaskResponseBodyDataJobs) GetStatus() *string {
	return s.Status
}

func (s *GetAgentTaskResponseBodyDataJobs) SetError(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Error = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetErrorCode(v string) *GetAgentTaskResponseBodyDataJobs {
	s.ErrorCode = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetErrorMessage(v string) *GetAgentTaskResponseBodyDataJobs {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetInstance(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Instance = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetParams(v interface{}) *GetAgentTaskResponseBodyDataJobs {
	s.Params = v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetRegion(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Region = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetResult(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Result = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetStatus(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Status = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) Validate() error {
	return dara.Validate(s)
}
