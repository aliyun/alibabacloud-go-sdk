// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobStatusDetailValue interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *JobStatusDetailValue
	GetComment() *string
	SetJobResult(v string) *JobStatusDetailValue
	GetJobResult() *string
	SetTimeStamps(v string) *JobStatusDetailValue
	GetTimeStamps() *string
}

type JobStatusDetailValue struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// {}
	JobResult *string `json:"jobResult,omitempty" xml:"jobResult,omitempty"`
	// example:
	//
	// 2022-06-13 17:11:34
	TimeStamps *string `json:"timeStamps,omitempty" xml:"timeStamps,omitempty"`
}

func (s JobStatusDetailValue) String() string {
	return dara.Prettify(s)
}

func (s JobStatusDetailValue) GoString() string {
	return s.String()
}

func (s *JobStatusDetailValue) GetComment() *string {
	return s.Comment
}

func (s *JobStatusDetailValue) GetJobResult() *string {
	return s.JobResult
}

func (s *JobStatusDetailValue) GetTimeStamps() *string {
	return s.TimeStamps
}

func (s *JobStatusDetailValue) SetComment(v string) *JobStatusDetailValue {
	s.Comment = &v
	return s
}

func (s *JobStatusDetailValue) SetJobResult(v string) *JobStatusDetailValue {
	s.JobResult = &v
	return s
}

func (s *JobStatusDetailValue) SetTimeStamps(v string) *JobStatusDetailValue {
	s.TimeStamps = &v
	return s
}

func (s *JobStatusDetailValue) Validate() error {
	return dara.Validate(s)
}
