// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderSyncJobResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderSyncJobResponseBodyData) *GetIdentityProviderSyncJobResponseBody
	GetData() *GetIdentityProviderSyncJobResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderSyncJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderSyncJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderSyncJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderSyncJobResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderSyncJobResponseBody struct {
	// The business status code of the API request. SUCCESS indicates that the request is successful, but does not indicate that the asynchronous synchronization task is completed.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data returned by the request.
	//
	// example:
	//
	// {"taskId":"12345678-1234-4234-8234-123456789abc","status":"SUCCESS","currentStep":"SYNC","synchronizationJobId":"syncjob_0123456789abcdef0123456789abcdef","syncedUserCount":12}
	Data *GetIdentityProviderSyncJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the API request. It does not indicate the execution result of the asynchronous task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message of the API request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request. Provide this value when you troubleshoot issues.
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. To determine whether the asynchronous task is completed, query the task status.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIdentityProviderSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderSyncJobResponseBody) GetData() *GetIdentityProviderSyncJobResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderSyncJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderSyncJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderSyncJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderSyncJobResponseBody) SetCode(v string) *GetIdentityProviderSyncJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) SetData(v *GetIdentityProviderSyncJobResponseBodyData) *GetIdentityProviderSyncJobResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderSyncJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) SetMessage(v string) *GetIdentityProviderSyncJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) SetRequestId(v string) *GetIdentityProviderSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) SetSuccess(v bool) *GetIdentityProviderSyncJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSyncJobResponseBodyData struct {
	// The current execution step of the AgentCore asynchronous task. This field may not be returned if the task has not started.
	//
	// example:
	//
	// SYNC
	CurrentStep *string `json:"currentStep,omitempty" xml:"currentStep,omitempty"`
	// The failure summary of the asynchronous task. This field is not returned if there is no failure information. You can use the taskId and requestId for troubleshooting.
	//
	// example:
	//
	// The identity provider synchronization job failed.
	FailureReason *string `json:"failureReason,omitempty" xml:"failureReason,omitempty"`
	// The status of the synchronization task. Valid values:
	//
	// - PENDING: The task is pending execution.
	//
	// - RUNNING: The task is running.
	//
	// - PAUSED: The task is paused.
	//
	// - SUCCESS: The task is successful.
	//
	// - FAILED: The task failed.
	//
	// - CANCELED: The task is canceled.
	//
	// The first three values are non-terminal states, and the last three values are terminal states. Query the task until it reaches a terminal state to determine the result.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of user entries that have been locally read and written or updated, including existing users. This value does not represent the number of newly added users or the synchronization percentage. This field is not returned when no statistics are available.
	//
	// example:
	//
	// 12
	SyncedUserCount *int32 `json:"syncedUserCount,omitempty" xml:"syncedUserCount,omitempty"`
	// The downstream synchronization task ID of AgentIdentity. This field is not returned if the task has not been created. It cannot replace the taskId to call GetIdentityProviderSyncJob.
	//
	// example:
	//
	// syncjob_0123456789abcdef0123456789abcdef
	SynchronizationJobId *string `json:"synchronizationJobId,omitempty" xml:"synchronizationJobId,omitempty"`
	// The local task ID of AgentCore returned by CreateIdentityProviderSyncJob. It is not the synchronizationJobId of AgentIdentity. The same workspace and identity provider must be used for the query.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetIdentityProviderSyncJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetFailureReason() *string {
	return s.FailureReason
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetSyncedUserCount() *int32 {
	return s.SyncedUserCount
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GetIdentityProviderSyncJobResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetCurrentStep(v string) *GetIdentityProviderSyncJobResponseBodyData {
	s.CurrentStep = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetFailureReason(v string) *GetIdentityProviderSyncJobResponseBodyData {
	s.FailureReason = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetStatus(v string) *GetIdentityProviderSyncJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetSyncedUserCount(v int32) *GetIdentityProviderSyncJobResponseBodyData {
	s.SyncedUserCount = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetSynchronizationJobId(v string) *GetIdentityProviderSyncJobResponseBodyData {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) SetTaskId(v string) *GetIdentityProviderSyncJobResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetIdentityProviderSyncJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
