// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *GetExecuteStateResponseBody
	GetErrorMessage() *string
	SetLogFile(v map[string]interface{}) *GetExecuteStateResponseBody
	GetLogFile() map[string]interface{}
	SetRequestId(v string) *GetExecuteStateResponseBody
	GetRequestId() *string
	SetState(v string) *GetExecuteStateResponseBody
	GetState() *string
	SetStatus(v string) *GetExecuteStateResponseBody
	GetStatus() *string
}

type GetExecuteStateResponseBody struct {
	// The error message.
	//
	// example:
	//
	// Your account does not have enough balance to order postpaid product.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The execution log.
	//
	// example:
	//
	// {"tf-plan.run.log":"xxx"}
	LogFile map[string]interface{} `json:"logFile,omitempty" xml:"logFile,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B4672AE3-C313-5B7A-BB24-45345570D398
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The content of the state file.
	//
	// example:
	//
	// {"version": 4, "terraform_version": "1.5.7", "serial": 3, "lineage": "cb71b0b2-1ec2-6483-d409-8cae23186ec6",  "outputs": {}, "resources": [], "check_results": null}
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The status. Valid values:
	//
	// - Pending: The task is ready.
	//
	// - PlanQueued: The plan task is queued and waiting for an available workflow.
	//
	// - ApplyQueued: The apply task is queued and waiting for an available workflow.
	//
	// - Planning: The plan phase is being executed.
	//
	// - Planned: The plan execution is completed.
	//
	// - Confirmed: The execution is confirmed after the plan phase.
	//
	// - PlannedAndFinished: No differences are found after the plan execution is completed, and the job enters the final state.
	//
	// - Applying: The apply phase is being executed.
	//
	// - Applied: The apply execution is completed.
	//
	// - Discarded: The execution task is discarded and enters the final state.
	//
	// - Errored: An error occurs during the execution, and the job enters the final state.
	//
	// - Canceled: The execution is canceled and enters the final state.
	//
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetExecuteStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteStateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetExecuteStateResponseBody) GetLogFile() map[string]interface{} {
	return s.LogFile
}

func (s *GetExecuteStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecuteStateResponseBody) GetState() *string {
	return s.State
}

func (s *GetExecuteStateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExecuteStateResponseBody) SetErrorMessage(v string) *GetExecuteStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetLogFile(v map[string]interface{}) *GetExecuteStateResponseBody {
	s.LogFile = v
	return s
}

func (s *GetExecuteStateResponseBody) SetRequestId(v string) *GetExecuteStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetState(v string) *GetExecuteStateResponseBody {
	s.State = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetStatus(v string) *GetExecuteStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetExecuteStateResponseBody) Validate() error {
	return dara.Validate(s)
}
