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
	// Request ID, which can be used for end-to-end Diagnosis
	//
	// example:
	//
	// 2E75336A-0DB2-5263-B201-A6488EC97B50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code
	//
	// - `code == Success` indicates that authorization Succeeded.
	//
	// - Any other status code indicates Failed to Authorize. When authorization fails, View the `message` field to obtain the detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned Data.
	Data *GetAgentTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Error message
	//
	// - If `code == Success`, this field is empty;
	//
	// - Otherwise, this field contains the Request error message.
	//
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
	// List of sub-Jobs
	Jobs   []*GetAgentTaskResponseBodyDataJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	Status *string                             `json:"status,omitempty" xml:"status,omitempty"`
	// Job ID.
	//
	// example:
	//
	// c41d8e3506224184a714682fea86d22d
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
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
	// When Job execution fails, this field contains the error message indicating the cause of the failure.
	//
	// example:
	//
	// 已废弃（误用）
	Error *string `json:"error,omitempty" xml:"error,omitempty"`
	// The error code indicating the reason for sub-job failure. Possible values:
	//
	// 	- empty: The job executed normally.
	//
	// 	- INSTANCE_NOT_SUPPORTED: The instance type is not supported.
	//
	// 	- INSTANCE_NOT_EXISTS: The instance does not exist.
	//
	// 	- INSTANCE_RELEASED: The instance has been released.
	//
	// 	- INSTANCE_NOT_RUNNING: The instance is not running.
	//
	// 	- INSTANCE_NOT_OWNED: The instance does not belong to the current account.
	//
	// 	- AGENT_ALREADY_INSTALLED: The agent is already installed.
	//
	// 	- AGENT_NOT_INSTALLED: The agent is not installed.
	//
	// 	- AGENT_SAME_VERSION: The version is the same.
	//
	// 	- HAS_RUNNING_JOB: There is a running job.
	//
	// 	- RPM_LOCK_HELD: The RPM lock is held.
	//
	// 	- DISK_SPACE_INSUFFICIENT: Insufficient disk space.
	//
	// 	- NODE_LOAD_HIGH: High edge zone load.
	//
	// 	- COMMAND_FAILED: Command execution failed.
	//
	// 	- CLIENT_NOT_RUNNING: The Cloud Assistant agent is not running.
	//
	// 	- CLIENT_NOT_RESPONSE: The Cloud Assistant agent is unresponsive.
	//
	// 	- DELIVERY_TIMEOUT: Command delivery timeout.
	//
	// 	- EXECUTION_TIMEOUT: Command execution timeout.
	//
	// 	- TASK_CONCURRENCY_LIMIT: Task concurrency limit reached.
	//
	// example:
	//
	// DISK_SPACE_INSUFFICIENT
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// Detailed reason for subtask execution failure. Possible values:
	//
	// 	- Instance type is not supported
	//
	// 	- Instance does not exist
	//
	// 	- Instance has been released
	//
	// 	- Instance is not running
	//
	// 	- Instance does not belong to the current account
	//
	// 	- Agent is already installed
	//
	// 	- Agent is not installed
	//
	// 	- Agent version is the same; no upgrade is required
	//
	// 	- A task is currently running; please retry later
	//
	// 	- RPM lock is occupied; please retry later
	//
	// 	- Insufficient disk space
	//
	// 	- Edge zone payload is too high; please retry later
	//
	// 	- Command execution failed; please retry later
	//
	// 	- Cloud Assistant Agent is not running
	//
	// 	- Cloud Assistant Agent is unresponsive
	//
	// 	- Command sending timeout
	//
	// 	- Command execution timeout
	//
	// 	- Task concurrency limit has been reached
	//
	// example:
	//
	// 磁盘空间不足
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// i-2zehme0rs1tc090fdl3n
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// Parameters of the sub-Job
	//
	// example:
	//
	// {
	//
	//     "agent_version": "3.5.0-beta",
	//
	//     "opt": "install",
	//
	//     "agent_id": "74a86327-3170-412c-8fd67-da3389ec56a9",
	//
	//     "install_type": "InstallAndUpgrade"
	//
	// }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Result of sub-Job execution
	//
	// example:
	//
	// 已废弃（误用）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// Sub-Job status:
	//
	// - Created: Created
	//
	// - Running: Running
	//
	// - Success: Job Run Succeeded
	//
	// - Fail: Job Run failed
	//
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
