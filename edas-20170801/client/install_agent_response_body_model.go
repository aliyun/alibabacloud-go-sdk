// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InstallAgentResponseBody
	GetCode() *int32
	SetExecutionResultList(v *InstallAgentResponseBodyExecutionResultList) *InstallAgentResponseBody
	GetExecutionResultList() *InstallAgentResponseBodyExecutionResultList
	SetMessage(v string) *InstallAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallAgentResponseBody
	GetRequestId() *string
}

type InstallAgentResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	ExecutionResultList *InstallAgentResponseBodyExecutionResultList `json:"ExecutionResultList,omitempty" xml:"ExecutionResultList,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallAgentResponseBody) GetExecutionResultList() *InstallAgentResponseBodyExecutionResultList {
	return s.ExecutionResultList
}

func (s *InstallAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAgentResponseBody) SetCode(v int32) *InstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentResponseBody) SetExecutionResultList(v *InstallAgentResponseBodyExecutionResultList) *InstallAgentResponseBody {
	s.ExecutionResultList = v
	return s
}

func (s *InstallAgentResponseBody) SetMessage(v string) *InstallAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAgentResponseBody) SetRequestId(v string) *InstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentResponseBody) Validate() error {
	if s.ExecutionResultList != nil {
		if err := s.ExecutionResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallAgentResponseBodyExecutionResultList struct {
	ExecutionResult []*InstallAgentResponseBodyExecutionResultListExecutionResult `json:"ExecutionResult,omitempty" xml:"ExecutionResult,omitempty" type:"Repeated"`
}

func (s InstallAgentResponseBodyExecutionResultList) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponseBodyExecutionResultList) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyExecutionResultList) GetExecutionResult() []*InstallAgentResponseBodyExecutionResultListExecutionResult {
	return s.ExecutionResult
}

func (s *InstallAgentResponseBodyExecutionResultList) SetExecutionResult(v []*InstallAgentResponseBodyExecutionResultListExecutionResult) *InstallAgentResponseBodyExecutionResultList {
	s.ExecutionResult = v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultList) Validate() error {
	if s.ExecutionResult != nil {
		for _, item := range s.ExecutionResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallAgentResponseBodyExecutionResultListExecutionResult struct {
	// The time when the installation was complete.
	//
	// example:
	//
	// 20**-11-10T07:02:17Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-2ze7s2v0b789k*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The state of the installation.
	//
	// example:
	//
	// Finished
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The state of the installation command.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the installation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallAgentResponseBodyExecutionResultListExecutionResult) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentResponseBodyExecutionResultListExecutionResult) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) GetStatus() *string {
	return s.Status
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) GetSuccess() *bool {
	return s.Success
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetFinishedTime(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.FinishedTime = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetInstanceId(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.InstanceId = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetInvokeRecordStatus(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetStatus(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.Status = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetSuccess(v bool) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.Success = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) Validate() error {
	return dara.Validate(s)
}
