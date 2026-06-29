// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJob interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v *SimpleUser) *Job
	GetCreator() *SimpleUser
	SetGmtCreateTime(v string) *Job
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Job
	GetGmtModifiedTime() *string
	SetJobId(v string) *Job
	GetJobId() *string
	SetJobResult(v *JobJobResult) *Job
	GetJobResult() *JobJobResult
	SetJobType(v string) *Job
	GetJobType() *string
	SetStatus(v string) *Job
	GetStatus() *string
}

type Job struct {
	// Job creator.
	Creator *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Creation Time.
	//
	// example:
	//
	// 2021-12-28 11:42:19
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At.
	//
	// example:
	//
	// 2021-12-28 11:42:20
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 147***441221943296
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Task Result.
	JobResult *JobJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	// Task Type. Currently, only DOWNLOWD_MARKRESULT_FLOW is supported.
	//
	// example:
	//
	// DOWNLOWD_MARKRESULT_FLOW
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Job status. Possible values:
	//
	// - init: initialization
	//
	// - running: running
	//
	// - pause: pause
	//
	// - stop: stopped
	//
	// - succ: Succeeded
	//
	// - fail: failed
	//
	// example:
	//
	// succ
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s Job) String() string {
	return dara.Prettify(s)
}

func (s Job) GoString() string {
	return s.String()
}

func (s *Job) GetCreator() *SimpleUser {
	return s.Creator
}

func (s *Job) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Job) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Job) GetJobId() *string {
	return s.JobId
}

func (s *Job) GetJobResult() *JobJobResult {
	return s.JobResult
}

func (s *Job) GetJobType() *string {
	return s.JobType
}

func (s *Job) GetStatus() *string {
	return s.Status
}

func (s *Job) SetCreator(v *SimpleUser) *Job {
	s.Creator = v
	return s
}

func (s *Job) SetGmtCreateTime(v string) *Job {
	s.GmtCreateTime = &v
	return s
}

func (s *Job) SetGmtModifiedTime(v string) *Job {
	s.GmtModifiedTime = &v
	return s
}

func (s *Job) SetJobId(v string) *Job {
	s.JobId = &v
	return s
}

func (s *Job) SetJobResult(v *JobJobResult) *Job {
	s.JobResult = v
	return s
}

func (s *Job) SetJobType(v string) *Job {
	s.JobType = &v
	return s
}

func (s *Job) SetStatus(v string) *Job {
	s.Status = &v
	return s
}

func (s *Job) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.JobResult != nil {
		if err := s.JobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JobJobResult struct {
	// Return value link, which is the OSS path where the annotation results are stored.
	//
	// example:
	//
	// oss://****-hz-oss.oss-cn-hangzhou.aliyuncs.com/output/ocr_demo001.manifest
	ResultLink *string `json:"ResultLink,omitempty" xml:"ResultLink,omitempty"`
}

func (s JobJobResult) String() string {
	return dara.Prettify(s)
}

func (s JobJobResult) GoString() string {
	return s.String()
}

func (s *JobJobResult) GetResultLink() *string {
	return s.ResultLink
}

func (s *JobJobResult) SetResultLink(v string) *JobJobResult {
	s.ResultLink = &v
	return s
}

func (s *JobJobResult) Validate() error {
	return dara.Validate(s)
}
