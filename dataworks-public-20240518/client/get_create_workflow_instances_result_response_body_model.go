// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateWorkflowInstancesResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCreateWorkflowInstancesResultResponseBody
	GetRequestId() *string
	SetResult(v *GetCreateWorkflowInstancesResultResponseBodyResult) *GetCreateWorkflowInstancesResultResponseBody
	GetResult() *GetCreateWorkflowInstancesResultResponseBodyResult
}

type GetCreateWorkflowInstancesResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The creation result of the workflow instance.
	Result *GetCreateWorkflowInstancesResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetCreateWorkflowInstancesResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreateWorkflowInstancesResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateWorkflowInstancesResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreateWorkflowInstancesResultResponseBody) GetResult() *GetCreateWorkflowInstancesResultResponseBodyResult {
	return s.Result
}

func (s *GetCreateWorkflowInstancesResultResponseBody) SetRequestId(v string) *GetCreateWorkflowInstancesResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBody) SetResult(v *GetCreateWorkflowInstancesResultResponseBodyResult) *GetCreateWorkflowInstancesResultResponseBody {
	s.Result = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCreateWorkflowInstancesResultResponseBodyResult struct {
	// The error message. This parameter is returned only if the creation fails.
	//
	// example:
	//
	// Invalid Param xxx
	FailureMessage *string `json:"FailureMessage,omitempty" xml:"FailureMessage,omitempty"`
	// The creation status. Valid values:
	//
	// 	- Creating
	//
	// 	- Created
	//
	// 	- CreateFailure
	//
	// example:
	//
	// Created
	Status                     *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	UnifiedWorkflowInstanceIds []*int64 `json:"UnifiedWorkflowInstanceIds,omitempty" xml:"UnifiedWorkflowInstanceIds,omitempty" type:"Repeated"`
	// The workflow instance IDs. This parameter is returned only if the creation is successful.
	WorkflowInstanceIds     []*int64 `json:"WorkflowInstanceIds,omitempty" xml:"WorkflowInstanceIds,omitempty" type:"Repeated"`
	WorkflowTaskInstanceIds []*int64 `json:"WorkflowTaskInstanceIds,omitempty" xml:"WorkflowTaskInstanceIds,omitempty" type:"Repeated"`
}

func (s GetCreateWorkflowInstancesResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCreateWorkflowInstancesResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) GetFailureMessage() *string {
	return s.FailureMessage
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) GetUnifiedWorkflowInstanceIds() []*int64 {
	return s.UnifiedWorkflowInstanceIds
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) GetWorkflowInstanceIds() []*int64 {
	return s.WorkflowInstanceIds
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) GetWorkflowTaskInstanceIds() []*int64 {
	return s.WorkflowTaskInstanceIds
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) SetFailureMessage(v string) *GetCreateWorkflowInstancesResultResponseBodyResult {
	s.FailureMessage = &v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) SetStatus(v string) *GetCreateWorkflowInstancesResultResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) SetUnifiedWorkflowInstanceIds(v []*int64) *GetCreateWorkflowInstancesResultResponseBodyResult {
	s.UnifiedWorkflowInstanceIds = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) SetWorkflowInstanceIds(v []*int64) *GetCreateWorkflowInstancesResultResponseBodyResult {
	s.WorkflowInstanceIds = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) SetWorkflowTaskInstanceIds(v []*int64) *GetCreateWorkflowInstancesResultResponseBodyResult {
	s.WorkflowTaskInstanceIds = v
	return s
}

func (s *GetCreateWorkflowInstancesResultResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
