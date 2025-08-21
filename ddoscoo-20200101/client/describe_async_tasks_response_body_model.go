// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAsyncTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTasks(v []*DescribeAsyncTasksResponseBodyAsyncTasks) *DescribeAsyncTasksResponseBody
	GetAsyncTasks() []*DescribeAsyncTasksResponseBodyAsyncTasks
	SetRequestId(v string) *DescribeAsyncTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAsyncTasksResponseBody
	GetTotalCount() *int32
}

type DescribeAsyncTasksResponseBody struct {
	// An array that consists of the details of the asynchronous export tasks.
	AsyncTasks []*DescribeAsyncTasksResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of asynchronous export tasks that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAsyncTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponseBody) GetAsyncTasks() []*DescribeAsyncTasksResponseBodyAsyncTasks {
	return s.AsyncTasks
}

func (s *DescribeAsyncTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAsyncTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAsyncTasksResponseBody) SetAsyncTasks(v []*DescribeAsyncTasksResponseBodyAsyncTasks) *DescribeAsyncTasksResponseBody {
	s.AsyncTasks = v
	return s
}

func (s *DescribeAsyncTasksResponseBody) SetRequestId(v string) *DescribeAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAsyncTasksResponseBody) SetTotalCount(v int32) *DescribeAsyncTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAsyncTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAsyncTasksResponseBodyAsyncTasks struct {
	// The end time of the task. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 157927362000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the task. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 156927362000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task parameter. The value is a JSON string. The returned field in the value varies based on the value of **TaskType**.
	//
	// If **TaskType*	- is set to **1**, **3**, **4**, **5**, or **6**, the following filed is returned:
	//
	// 	- **instanceId**: the ID of the instance. Data type: string.
	//
	// If **TaskType*	- is set to **2**, the following field is returned:
	//
	// 	- **domain**: the domain name of the website. Data type: string.
	//
	// example:
	//
	// {"instanceId": "ddoscoo-cn-mp91j1ao****"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The execution result of the task. The value is a JSON string. The returned fields in the value vary based on the value of **TaskType**.
	//
	// If **TaskType*	- is set to **1**, **3**, **4**, **5**, or **6**, the following fields are returned:
	//
	// 	- **instanceId**: the ID of the instance. Data type: string.
	//
	// 	- **url**: the URL to download the exported file from Object Storage Service (OSS). Data type: string.
	//
	// If **TaskType*	- is set to **2**, the following fields are returned:
	//
	// 	- **domain**: the domain name of the website. Data type: string.
	//
	// 	- **url**: the URL to download the exported file from OSS. Data type: string.
	//
	// example:
	//
	// {"instanceId": "ddoscoo-cn-mp91j1ao****","url": "https://****.oss-cn-beijing.aliyuncs.com/heap.bin?Expires=1584785140&OSSAccessKeyId=TMP.3KfzD82FyRJevJdEkRX6JEFHhbvRBBb75PZJnyJmksA2QkMm47xFAFDgMhEV8Nm6Vxr8xExMfiy9LsUFAcLcTBrN3r****&Signature=Sj8BNcsxJLE8l5qm4cjNlDt8gv****"}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: indicates that the task is being initialized.
	//
	// 	- **1**: indicates that the task is in progress.
	//
	// 	- **2**: indicates that the task is successful.
	//
	// 	- **3**: indicates that the task failed.
	//
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **1**: the task to export the port forwarding rules of an instance
	//
	// 	- **2**: the task to export the forwarding rules of a website protected by an instance
	//
	// 	- **3**: the task to export the sessions and health check settings of an instance
	//
	// 	- **4**: the task to export the mitigation policies of an instance
	//
	// 	- **5**: the task to download the blacklist for destination IP addresses of an instance
	//
	// 	- **6**: the task to download the whitelist for destination IP addresses of an instance
	//
	// example:
	//
	// 5
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeAsyncTasksResponseBodyAsyncTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAsyncTasksResponseBodyAsyncTasks) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetTaskResult() *string {
	return s.TaskResult
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) GetTaskType() *int32 {
	return s.TaskType
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetEndTime(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetStartTime(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskId(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskParams(v string) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskResult(v string) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskStatus(v int32) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskType(v int32) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) Validate() error {
	return dara.Validate(s)
}
