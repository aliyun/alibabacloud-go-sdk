// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIdentityProviderSyncJobResponseBody
	GetCode() *string
	SetData(v *CreateIdentityProviderSyncJobResponseBodyData) *CreateIdentityProviderSyncJobResponseBody
	GetData() *CreateIdentityProviderSyncJobResponseBodyData
	SetHttpStatusCode(v int32) *CreateIdentityProviderSyncJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateIdentityProviderSyncJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIdentityProviderSyncJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIdentityProviderSyncJobResponseBody
	GetSuccess() *bool
}

type CreateIdentityProviderSyncJobResponseBody struct {
	// The business status code of the API request. SUCCESS indicates that the request is successful, but does not indicate that the asynchronous synchronization task is completed.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data returned in the response.
	//
	// example:
	//
	// {"taskId":"12345678-1234-4234-8234-123456789abc","status":"PENDING"}
	Data *CreateIdentityProviderSyncJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the API request, which does not indicate the execution result of the asynchronous task.
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
	// The unique identifier of the request. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. To check whether the asynchronous task is completed, query the task status.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateIdentityProviderSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetData() *CreateIdentityProviderSyncJobResponseBodyData {
	return s.Data
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderSyncJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetCode(v string) *CreateIdentityProviderSyncJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetData(v *CreateIdentityProviderSyncJobResponseBodyData) *CreateIdentityProviderSyncJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetHttpStatusCode(v int32) *CreateIdentityProviderSyncJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetMessage(v string) *CreateIdentityProviderSyncJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetRequestId(v string) *CreateIdentityProviderSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) SetSuccess(v bool) *CreateIdentityProviderSyncJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderSyncJobResponseBodyData struct {
	// The current execution step of the AgentCore asynchronous task. This field may not be returned if the task has not started.
	//
	// example:
	//
	// SYNC
	CurrentStep *string `json:"currentStep,omitempty" xml:"currentStep,omitempty"`
	// The failure summary of the asynchronous task. This field is not returned if no failure information is available. You can use this field along with the taskId and requestId for troubleshooting.
	//
	// example:
	//
	// The identity provider synchronization job failed.
	FailureReason *string `json:"failureReason,omitempty" xml:"failureReason,omitempty"`
	// The status of the synchronization task. Valid values: PENDING: pending. RUNNING: running. PAUSED: paused. SUCCESS: successful. FAILED: failed. CANCELED: canceled. The first three are non-terminal states, and the last three are terminal states. Query the status until a terminal state is reached to determine the result.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of user entries that have been read locally and written or updated, including existing users. This does not represent the number of newly added users or the synchronization percentage. This field is not returned when no statistics are available.
	//
	// example:
	//
	// 12
	SyncedUserCount *int32 `json:"syncedUserCount,omitempty" xml:"syncedUserCount,omitempty"`
	// The ID of the downstream synchronization task in AgentIdentity. This field is not returned if the task has not been created. You cannot use this ID in place of taskId to call GetIdentityProviderSyncJob.
	//
	// example:
	//
	// syncjob_0123456789abcdef0123456789abcdef
	SynchronizationJobId *string `json:"synchronizationJobId,omitempty" xml:"synchronizationJobId,omitempty"`
	// The local task ID in AgentCore returned by CreateIdentityProviderSyncJob, which is different from the synchronizationJobId in AgentIdentity. You must use the same workspace and identity provider when querying the task.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateIdentityProviderSyncJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSyncJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetFailureReason() *string {
	return s.FailureReason
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetSyncedUserCount() *int32 {
	return s.SyncedUserCount
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetCurrentStep(v string) *CreateIdentityProviderSyncJobResponseBodyData {
	s.CurrentStep = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetFailureReason(v string) *CreateIdentityProviderSyncJobResponseBodyData {
	s.FailureReason = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetStatus(v string) *CreateIdentityProviderSyncJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetSyncedUserCount(v int32) *CreateIdentityProviderSyncJobResponseBodyData {
	s.SyncedUserCount = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetSynchronizationJobId(v string) *CreateIdentityProviderSyncJobResponseBodyData {
	s.SynchronizationJobId = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) SetTaskId(v string) *CreateIdentityProviderSyncJobResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateIdentityProviderSyncJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
