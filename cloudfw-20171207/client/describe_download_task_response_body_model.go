// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDownloadTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeDownloadTaskResponseBodyTasks) *DescribeDownloadTaskResponseBody
	GetTasks() []*DescribeDownloadTaskResponseBodyTasks
	SetTotalCount(v int32) *DescribeDownloadTaskResponseBody
	GetTotalCount() *int32
}

type DescribeDownloadTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3F2BED13-F3D0-5984-80D6-D5F298CFEA88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*DescribeDownloadTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDownloadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadTaskResponseBody) GetTasks() []*DescribeDownloadTaskResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeDownloadTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDownloadTaskResponseBody) SetRequestId(v string) *DescribeDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetTasks(v []*DescribeDownloadTaskResponseBodyTasks) *DescribeDownloadTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetTotalCount(v int32) *DescribeDownloadTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDownloadTaskResponseBodyTasks struct {
	// The time when the task was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1706595827
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task expires. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1714371828
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The size of the file.
	//
	// example:
	//
	// 24.04KB
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The URL of the OSS file.
	//
	// example:
	//
	// https://cfw-table-download-cn.oss-cn-hangzhou.aliyuncs.com/%E4%BA%92%E8%81%94%E7%BD%91%E8%BE%B9%E7%95%8C%E9%98%B2%E7%81%AB%E5%A2%99%E8%B5%84%E4%BA%A7-IPv4_1069.csv?Expires=1708583913&OSSAccessKeyId=****&Signature=******%3D
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **finish**
	//
	// 	- **start**
	//
	// 	- **error**
	//
	// 	- **expire**: The task file is invalid and cannot be downloaded.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1111
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Internet Boundary Firewall Assets - IPv4
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// InternetFirewallAsset
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetFileSize() *string {
	return s.FileSize
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetFileURL() *string {
	return s.FileURL
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeDownloadTaskResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetCreateTime(v int64) *DescribeDownloadTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetExpireTime(v int64) *DescribeDownloadTaskResponseBodyTasks {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetFileSize(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.FileSize = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetFileURL(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.FileURL = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetStatus(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskId(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskName(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskType(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
