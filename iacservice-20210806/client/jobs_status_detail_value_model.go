// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobsStatusDetailValue interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *JobsStatusDetailValue
	GetComment() *string
	SetJobResult(v string) *JobsStatusDetailValue
	GetJobResult() *string
	SetTimeStamps(v string) *JobsStatusDetailValue
	GetTimeStamps() *string
}

type JobsStatusDetailValue struct {
	// 备注
	//
	// example:
	//
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 作业执行结果
	//
	// example:
	//
	// Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
	JobResult *string `json:"jobResult,omitempty" xml:"jobResult,omitempty"`
	// 到达状态时间
	//
	// example:
	//
	// 2022-06-13 17:11:34
	TimeStamps *string `json:"timeStamps,omitempty" xml:"timeStamps,omitempty"`
}

func (s JobsStatusDetailValue) String() string {
	return dara.Prettify(s)
}

func (s JobsStatusDetailValue) GoString() string {
	return s.String()
}

func (s *JobsStatusDetailValue) GetComment() *string {
	return s.Comment
}

func (s *JobsStatusDetailValue) GetJobResult() *string {
	return s.JobResult
}

func (s *JobsStatusDetailValue) GetTimeStamps() *string {
	return s.TimeStamps
}

func (s *JobsStatusDetailValue) SetComment(v string) *JobsStatusDetailValue {
	s.Comment = &v
	return s
}

func (s *JobsStatusDetailValue) SetJobResult(v string) *JobsStatusDetailValue {
	s.JobResult = &v
	return s
}

func (s *JobsStatusDetailValue) SetTimeStamps(v string) *JobsStatusDetailValue {
	s.TimeStamps = &v
	return s
}

func (s *JobsStatusDetailValue) Validate() error {
	return dara.Validate(s)
}
