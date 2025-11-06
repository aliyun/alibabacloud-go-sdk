// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryTaskDetailListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryTaskDetailListResponseBodyData) *QueryTaskDetailListResponseBody
	GetData() *QueryTaskDetailListResponseBodyData
	SetNextPage(v bool) *QueryTaskDetailListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryTaskDetailListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryTaskDetailListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryTaskDetailListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryTaskDetailListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryTaskDetailListResponseBody
	GetTotalPageNum() *int32
}

type QueryTaskDetailListResponseBody struct {
	// The page number returned.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The tasks.
	Data *QueryTaskDetailListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the current page is preceded by a page.
	//
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A2320E4-D870-49C9-A6DC-test
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTaskDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryTaskDetailListResponseBody) GetData() *QueryTaskDetailListResponseBodyData {
	return s.Data
}

func (s *QueryTaskDetailListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryTaskDetailListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskDetailListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryTaskDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskDetailListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryTaskDetailListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryTaskDetailListResponseBody) SetCurrentPageNum(v int32) *QueryTaskDetailListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetData(v *QueryTaskDetailListResponseBodyData) *QueryTaskDetailListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetNextPage(v bool) *QueryTaskDetailListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetPageSize(v int32) *QueryTaskDetailListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetPrePage(v bool) *QueryTaskDetailListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetRequestId(v string) *QueryTaskDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetTotalItemNum(v int32) *QueryTaskDetailListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetTotalPageNum(v int32) *QueryTaskDetailListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskDetailListResponseBodyData struct {
	TaskDetail []*QueryTaskDetailListResponseBodyDataTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Repeated"`
}

func (s QueryTaskDetailListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBodyData) GetTaskDetail() []*QueryTaskDetailListResponseBodyDataTaskDetail {
	return s.TaskDetail
}

func (s *QueryTaskDetailListResponseBodyData) SetTaskDetail(v []*QueryTaskDetailListResponseBodyDataTaskDetail) *QueryTaskDetailListResponseBodyData {
	s.TaskDetail = v
	return s
}

func (s *QueryTaskDetailListResponseBodyData) Validate() error {
	if s.TaskDetail != nil {
		for _, item := range s.TaskDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTaskDetailListResponseBodyDataTaskDetail struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2018-01-25 20:46:57
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The error message returned when the task failed.
	//
	// example:
	//
	// The operation is successful.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The cause of a domain name task failure.
	//
	// example:
	//
	// Purchase record exists for the domain name in Aliyun
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The instance ID of the domain name.
	//
	// example:
	//
	// S20179H1BBI9test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task details.
	//
	// example:
	//
	// 75addb07-28a3-450e-b5ec-test
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 60d6e201-8ee5-47ab-8fdc-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// The result of the task.
	//
	// example:
	//
	// 12345
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The task status. Valid values:
	//
	// 	- **WAITING_EXECUTE**: To be executed
	//
	// 	- **EXECUTING**: being executed
	//
	// 	- **EXECUTE_SUCCESS**: successful
	//
	// 	- **EXECUTE_FAILURE**: failed
	//
	// example:
	//
	// EXECUTE_SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The status code of the task. Valid values:
	//
	// 	- **0**: waiting for execution
	//
	// 	- **1**: being executed
	//
	// 	- **2**: successful
	//
	// 	- **3**: failed
	//
	// example:
	//
	// 2
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// The task type. Valid values:
	//
	// 	- **CHG_HOLDER**: The task is run to modify the domain name registrant.
	//
	// 	- **CHG_DNS**: The task is run to change the Domain Name System (DNS) servers.
	//
	// 	- **SET_WHOIS_PROTECT**: The task is run to configure privacy protection for the domain name.
	//
	// 	- **UPDATE_ADMIN_CONTACT**: The task is run to modify the information about the administrator of the domain name.
	//
	// 	- **UPDATE_BILLING_CONTACT**: The task is run to modify the information about the payer for the domain name.
	//
	// 	- **UPDATE_TECH_CONTACT**: The task is run to modify the information about the technical support for the domain name.
	//
	// 	- **SET_UPDATE_PROHIBITED**: The task is run to configure the security lock for the domain name.
	//
	// 	- **SET_TRANSFER_PROHIBITED**: The task is run to configure the transfer lock for the domain name.
	//
	// 	- **ORDER_ACTIVATE**: The task is run to create a registration order for the domain name.
	//
	// 	- **ORDER_RENEW**: The task is run to create a renewal order for the domain name.
	//
	// 	- **ORDER_REDEEM**: The task is run to create a redemption order for the domain name.
	//
	// 	- **CREATE_DNSHOST**: The task is run to create a DNS server for the domain name.
	//
	// 	- **UPDATE_DNSHOST**: The task is run to update the information about a DNS server for the domain name.
	//
	// 	- **SYNC_DNSHOST**: The task is run to synchronize a DNS server for the domain name.
	//
	// example:
	//
	// ORDER_RENEW
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The description of the task type.
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// The number of times the task was retried.
	//
	// example:
	//
	// 0
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// The last time when the task was run.
	//
	// example:
	//
	// 2018-01-25 20:47:01
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailListResponseBodyDataTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailListResponseBodyDataTaskDetail) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetFailReason() *string {
	return s.FailReason
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskResult() *string {
	return s.TaskResult
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetCreateTime(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetDomainName(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetErrorMsg(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetFailReason(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.FailReason = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetInstanceId(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskDetailNo(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskNo(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskResult(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskResult = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskStatus(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskStatusCode(v int32) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskType(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskTypeDescription(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTryCount(v int32) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetUpdateTime(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.UpdateTime = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) Validate() error {
	return dara.Validate(s)
}
