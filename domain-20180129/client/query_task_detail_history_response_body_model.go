// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageCursor(v *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) *QueryTaskDetailHistoryResponseBody
	GetCurrentPageCursor() *QueryTaskDetailHistoryResponseBodyCurrentPageCursor
	SetNextPageCursor(v *QueryTaskDetailHistoryResponseBodyNextPageCursor) *QueryTaskDetailHistoryResponseBody
	GetNextPageCursor() *QueryTaskDetailHistoryResponseBodyNextPageCursor
	SetObjects(v []*QueryTaskDetailHistoryResponseBodyObjects) *QueryTaskDetailHistoryResponseBody
	GetObjects() []*QueryTaskDetailHistoryResponseBodyObjects
	SetPageSize(v int32) *QueryTaskDetailHistoryResponseBody
	GetPageSize() *int32
	SetPrePageCursor(v *QueryTaskDetailHistoryResponseBodyPrePageCursor) *QueryTaskDetailHistoryResponseBody
	GetPrePageCursor() *QueryTaskDetailHistoryResponseBodyPrePageCursor
	SetRequestId(v string) *QueryTaskDetailHistoryResponseBody
	GetRequestId() *string
}

type QueryTaskDetailHistoryResponseBody struct {
	// Current page cursor.
	CurrentPageCursor *QueryTaskDetailHistoryResponseBodyCurrentPageCursor `json:"CurrentPageCursor,omitempty" xml:"CurrentPageCursor,omitempty" type:"Struct"`
	// Cursor for the next page.
	NextPageCursor *QueryTaskDetailHistoryResponseBodyNextPageCursor `json:"NextPageCursor,omitempty" xml:"NextPageCursor,omitempty" type:"Struct"`
	// Task detail information.
	Objects []*QueryTaskDetailHistoryResponseBodyObjects `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// Paging size.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Cursor for the previous page.
	PrePageCursor *QueryTaskDetailHistoryResponseBodyPrePageCursor `json:"PrePageCursor,omitempty" xml:"PrePageCursor,omitempty" type:"Struct"`
	// Unique Request access token.
	//
	// example:
	//
	// 548CAE74-88F8-402F-8C12-97E747389C51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBody) GetCurrentPageCursor() *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	return s.CurrentPageCursor
}

func (s *QueryTaskDetailHistoryResponseBody) GetNextPageCursor() *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	return s.NextPageCursor
}

func (s *QueryTaskDetailHistoryResponseBody) GetObjects() []*QueryTaskDetailHistoryResponseBodyObjects {
	return s.Objects
}

func (s *QueryTaskDetailHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskDetailHistoryResponseBody) GetPrePageCursor() *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	return s.PrePageCursor
}

func (s *QueryTaskDetailHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskDetailHistoryResponseBody) SetCurrentPageCursor(v *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) *QueryTaskDetailHistoryResponseBody {
	s.CurrentPageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetNextPageCursor(v *QueryTaskDetailHistoryResponseBodyNextPageCursor) *QueryTaskDetailHistoryResponseBody {
	s.NextPageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetObjects(v []*QueryTaskDetailHistoryResponseBodyObjects) *QueryTaskDetailHistoryResponseBody {
	s.Objects = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetPageSize(v int32) *QueryTaskDetailHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetPrePageCursor(v *QueryTaskDetailHistoryResponseBodyPrePageCursor) *QueryTaskDetailHistoryResponseBody {
	s.PrePageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetRequestId(v string) *QueryTaskDetailHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) Validate() error {
	if s.CurrentPageCursor != nil {
		if err := s.CurrentPageCursor.Validate(); err != nil {
			return err
		}
	}
	if s.NextPageCursor != nil {
		if err := s.NextPageCursor.Validate(); err != nil {
			return err
		}
	}
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrePageCursor != nil {
		if err := s.PrePageCursor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskDetailHistoryResponseBodyCurrentPageCursor struct {
	// Job Creation Time.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Result of task execution.
	//
	// example:
	//
	// 执行成功
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Domain instance ID.
	//
	// example:
	//
	// S1234456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task detail ID.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-2342
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// Job number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// Task Status. Valid values:
	//
	// - **WAITING_EXECUTE**: Waiting for execution.
	//
	// - **EXECUTING**: Executing.
	//
	// - **EXECUTE_SUCCESS**: Execution succeeded.
	//
	// - **EXECUTE_FAILURE**: Execution failed.
	//
	// example:
	//
	// EXECUTE_SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Job Status code. Valid values:
	//
	// - **0**: Waiting to execute.
	//
	// - **1**: Executing.
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// example:
	//
	// 2
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// Task Type. Valid values:
	//
	// - **CHG_HOLDER**: Modify registrant information.
	//
	// - **CHG_DNS**: Modify DNS.
	//
	// - **SET_WHOIS_PROTECT**: Enable privacy protection.
	//
	// - **UPDATE_ADMIN_CONTACT**: Modify administrative contact information.
	//
	// - **UPDATE_BILLING_CONTACT**: Modify billing contact information.
	//
	// - **UPDATE_TECH_CONTACT**: Modify technical contact information.
	//
	// - **SET_UPDATE_PROHIBITED**: Enable domain name edit lock.
	//
	// - **SET_TRANSFER_PROHIBITED**: Enable domain name transfer lock.
	//
	// - **ORDER_ACTIVATE**: Create a registration order.
	//
	// - **ORDER_RENEW**: Create a renewal order.
	//
	// - **ORDER_REDEEM**: Create a redemption order.
	//
	// - **CREATE_DNSHOST**: Create a DNS host.
	//
	// - **UPDATE_DNSHOST**: Update a DNS host.
	//
	// - **SYNC_DNSHOST**: Synchronize a DNS host.
	//
	// example:
	//
	// CHG_DNS
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Description of the task type.
	//
	// example:
	//
	// 修改DNS
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// Retry Count of job details.
	//
	// example:
	//
	// 0
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// The most recent task execution time.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyCurrentPageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.UpdateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) Validate() error {
	return dara.Validate(s)
}

type QueryTaskDetailHistoryResponseBodyNextPageCursor struct {
	// Creation time of the job.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Result of task execution.
	//
	// example:
	//
	// 域名有禁止更新锁
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Domain name instance ID.
	//
	// example:
	//
	// S1234567890
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task detail number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-2424
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// Job number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// Task Status. Valid values:
	//
	// - **WAITING_EXECUTE**: Waiting for execution.
	//
	// - **EXECUTING**: Executing.
	//
	// - **EXECUTE_SUCCESS**: Succeeded.
	//
	// - **EXECUTE_FAILURE**: Failed.
	//
	// example:
	//
	// EXECUTE_FAILURE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Task status code. Valid values:
	//
	// - **0**: Waiting for execution.
	//
	// - **1**: Executing.
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// Task Type. Valid values:
	//
	// - **CHG_HOLDER**: Modify registrant information.
	//
	// - **CHG_DNS**: Modify DNS.
	//
	// - **SET_WHOIS_PROTECT**: Enable privacy protection.
	//
	// - **UPDATE_ADMIN_CONTACT**: Modify administrator contact information.
	//
	// - **UPDATE_BILLING_CONTACT**: Modify billing contact information.
	//
	// - **UPDATE_TECH_CONTACT**: Modify technical contact information.
	//
	// - **SET_UPDATE_PROHIBITED**: Enable Edit Lock.
	//
	// - **SET_TRANSFER_PROHIBITED**: Enable transfer lock.
	//
	// - **ORDER_ACTIVATE**: Create a registration order.
	//
	// - **ORDER_RENEW**: Create a renewal order.
	//
	// - **ORDER_REDEEM**: Create a redemption order.
	//
	// - **CREATE_DNSHOST**: Create a DNS host.
	//
	// - **UPDATE_DNSHOST**: Update a DNS host.
	//
	// - **SYNC_DNSHOST**: Synchronize a DNS host.
	//
	// example:
	//
	// CHG_DNS
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Task Type Description.
	//
	// example:
	//
	// 修改DNS
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// Number of retries for the task details.
	//
	// example:
	//
	// 5
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// The most recent running time of the job details.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyNextPageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyNextPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.UpdateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) Validate() error {
	return dara.Validate(s)
}

type QueryTaskDetailHistoryResponseBodyObjects struct {
	// The creation time of the job.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The result of the job execution.
	//
	// example:
	//
	// 域名有禁止更新锁
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The instance ID of the domain name.
	//
	// example:
	//
	// S123456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task detail number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-4234
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// The job number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// Task Status. Valid values:
	//
	// - **WAITING_EXECUTE**: Waiting for execution.
	//
	// - **EXECUTING**: Executing.
	//
	// - **EXECUTE_SUCCESS**: Execution succeeded.
	//
	// - **EXECUTE_FAILURE**: Execution failed.
	//
	// example:
	//
	// EXECUTE_FAILURE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The job status code. Valid values:
	//
	// - **0**: Waiting for execution.
	//
	// - **1**: Executing.
	//
	// - **2**: Succeeded.
	//
	// - **3**: Failed.
	//
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// The task type. Valid values:
	//
	// - **CHG_HOLDER**: Modify registrant information.
	//
	// - **CHG_DNS**: Modify DNS settings.
	//
	// - **SET_WHOIS_PROTECT**: Enable privacy protection.
	//
	// - **UPDATE_ADMIN_CONTACT**: Update administrative contact information.
	//
	// - **UPDATE_BILLING_CONTACT**: Update billing contact information.
	//
	// - **UPDATE_TECH_CONTACT**: Update technical contact information.
	//
	// - **SET_UPDATE_PROHIBITED**: Enable the Edit Lock for the domain name.
	//
	// - **SET_TRANSFER_PROHIBITED**: Enable the transfer lock for the domain name.
	//
	// - **ORDER_ACTIVATE**: Create a registration order.
	//
	// - **ORDER_RENEW**: Create a renewal order.
	//
	// - **ORDER_REDEEM**: Create a redemption order.
	//
	// - **CREATE_DNSHOST**: Create a DNS host.
	//
	// - **UPDATE_DNSHOST**: Update a DNS host.
	//
	// - **SYNC_DNSHOST**: Synchronize a DNS host.
	//
	// example:
	//
	// CHG_DNS
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Task Type description.
	//
	// example:
	//
	// 修改DNS
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// Number of retries for the task detail.
	//
	// example:
	//
	// 5
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// The running time of the most recent job execution.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyObjects) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.UpdateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) Validate() error {
	return dara.Validate(s)
}

type QueryTaskDetailHistoryResponseBodyPrePageCursor struct {
	// Task creation time.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Result of task execution.
	//
	// example:
	//
	// 域名有禁止更新锁
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Domain instance ID.
	//
	// example:
	//
	// S123456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task detail number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-123
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// Task number.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// Task Status. Valid values:
	//
	// - **WAITING_EXECUTE**: Waiting for execution.
	//
	// - **EXECUTING**: Executing.
	//
	// - **EXECUTE_SUCCESS**: Execution succeeded.
	//
	// - **EXECUTE_FAILURE**: Execution failed.
	//
	// example:
	//
	// EXECUTE_FAILURE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Task status code. Valid values:
	//
	// - **0**: Waiting for execution.
	//
	// - **1**: Executing.
	//
	// - **2**: Execution succeeded.
	//
	// - **3**: Execution failed.
	//
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// Task Type. Valid values:
	//
	// - **CHG_HOLDER**: Modify registrant information.
	//
	// - **CHG_DNS**: Modify DNS.
	//
	// - **SET_WHOIS_PROTECT**: Enable privacy protection.
	//
	// - **UPDATE_ADMIN_CONTACT**: Modify administrative contact information.
	//
	// - **UPDATE_BILLING_CONTACT**: Modify billing contact information.
	//
	// - **UPDATE_TECH_CONTACT**: Modify technical contact information.
	//
	// - **SET_UPDATE_PROHIBITED**: Enable domain name edit lock.
	//
	// - **SET_TRANSFER_PROHIBITED**: Enable domain name transfer lock.
	//
	// - **ORDER_ACTIVATE**: Create a registration order.
	//
	// - **ORDER_RENEW**: Create a renewal order.
	//
	// - **ORDER_REDEEM**: Create a redemption order.
	//
	// - **CREATE_DNSHOST**: Create a DNS host.
	//
	// - **UPDATE_DNSHOST**: Update a DNS host.
	//
	// - **SYNC_DNSHOST**: Synchronize a DNS host.
	//
	// example:
	//
	// CHG_DNS
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Description of the task type.
	//
	// example:
	//
	// 修改DNS
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// Number of retries for the task detail.
	//
	// example:
	//
	// 5
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// The most recent running time of the task details.
	//
	// example:
	//
	// 2019-07-30 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyPrePageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyPrePageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.UpdateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) Validate() error {
	return dara.Validate(s)
}
