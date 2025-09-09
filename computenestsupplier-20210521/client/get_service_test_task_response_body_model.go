// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionDetails(v []*GetServiceTestTaskResponseBodyExecutionDetails) *GetServiceTestTaskResponseBody
	GetExecutionDetails() []*GetServiceTestTaskResponseBodyExecutionDetails
	SetRequestId(v string) *GetServiceTestTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetServiceTestTaskResponseBody
	GetStatus() *string
	SetTaskName(v string) *GetServiceTestTaskResponseBody
	GetTaskName() *string
	SetTaskRegionId(v string) *GetServiceTestTaskResponseBody
	GetTaskRegionId() *string
}

type GetServiceTestTaskResponseBody struct {
	// The execution details.
	ExecutionDetails []*GetServiceTestTaskResponseBodyExecutionDetails `json:"ExecutionDetails,omitempty" xml:"ExecutionDetails,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service test task. Valid values:
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Failure
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// nametest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task execution region.
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
}

func (s GetServiceTestTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponseBody) GetExecutionDetails() []*GetServiceTestTaskResponseBodyExecutionDetails {
	return s.ExecutionDetails
}

func (s *GetServiceTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceTestTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceTestTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *GetServiceTestTaskResponseBody) GetTaskRegionId() *string {
	return s.TaskRegionId
}

func (s *GetServiceTestTaskResponseBody) SetExecutionDetails(v []*GetServiceTestTaskResponseBodyExecutionDetails) *GetServiceTestTaskResponseBody {
	s.ExecutionDetails = v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetRequestId(v string) *GetServiceTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetStatus(v string) *GetServiceTestTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetTaskName(v string) *GetServiceTestTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetTaskRegionId(v string) *GetServiceTestTaskResponseBody {
	s.TaskRegionId = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceTestTaskResponseBodyExecutionDetails struct {
	// The service test case name.
	//
	// example:
	//
	// case1
	CaseName *string `json:"CaseName,omitempty" xml:"CaseName,omitempty"`
	// The execution report
	//
	// example:
	//
	// -----------------------------------------------------------------------------
	//
	// Region: cn-qingdao
	//
	// StackName: iact3-default-cn-qingd
	//
	// StackId: 009d2991-f494-d
	//
	// *****************************************************************************
	ExecutionReport *string `json:"ExecutionReport,omitempty" xml:"ExecutionReport,omitempty"`
	// The sub task status.
	//
	// example:
	//
	// Runing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub task id.
	//
	// example:
	//
	// stt-xxxx
	SubTaskId *string `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
}

func (s GetServiceTestTaskResponseBodyExecutionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTestTaskResponseBodyExecutionDetails) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) GetCaseName() *string {
	return s.CaseName
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) GetExecutionReport() *string {
	return s.ExecutionReport
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) GetStatus() *string {
	return s.Status
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) GetSubTaskId() *string {
	return s.SubTaskId
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetCaseName(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.CaseName = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetExecutionReport(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.ExecutionReport = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetStatus(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.Status = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetSubTaskId(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.SubTaskId = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) Validate() error {
	return dara.Validate(s)
}
