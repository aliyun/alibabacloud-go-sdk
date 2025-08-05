// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTaskResponseBody
	GetCode() *string
	SetCompletedTime(v int64) *DescribeTaskResponseBody
	GetCompletedTime() *int64
	SetCreatedTime(v int64) *DescribeTaskResponseBody
	GetCreatedTime() *int64
	SetDescription(v string) *DescribeTaskResponseBody
	GetDescription() *string
	SetMessage(v string) *DescribeTaskResponseBody
	GetMessage() *string
	SetName(v string) *DescribeTaskResponseBody
	GetName() *string
	SetProgress(v int32) *DescribeTaskResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeTaskResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeTaskResponseBody
	GetResult() *string
	SetSuccess(v bool) *DescribeTaskResponseBody
	GetSuccess() *bool
	SetUpdatedTime(v int64) *DescribeTaskResponseBody
	GetUpdatedTime() *int64
}

type DescribeTaskResponseBody struct {
	// HttpCode
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the task was complete. The time is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615607706
	CompletedTime *int64 `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// The time when the job was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615607706
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **created**: The job is created.
	//
	// 	- **expired**: The job expires.
	//
	// 	- **completed**: The job is completed.
	//
	// 	- **cancelled**: The job is canceled.
	//
	// example:
	//
	// completed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// InstallBackupClientsTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The progress of the job. Valid values: 0 to 100. Unit: percentage (%). If the job fails, the value -1 is returned.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the job.
	//
	// example:
	//
	// {}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the call is successful.
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the job was updated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615607706
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTaskResponseBody) GetCompletedTime() *int64 {
	return s.CompletedTime
}

func (s *DescribeTaskResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeTaskResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTaskResponseBody) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeTaskResponseBody) SetCode(v string) *DescribeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCompletedTime(v int64) *DescribeTaskResponseBody {
	s.CompletedTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCreatedTime(v int64) *DescribeTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDescription(v string) *DescribeTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTaskResponseBody) SetMessage(v string) *DescribeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBody) SetName(v string) *DescribeTaskResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTaskResponseBody) SetProgress(v int32) *DescribeTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetResult(v string) *DescribeTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeTaskResponseBody) SetSuccess(v bool) *DescribeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdatedTime(v int64) *DescribeTaskResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
