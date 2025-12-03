// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineJobsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineJobsResponseBody
	GetErrorMessage() *string
	SetJobs(v []*ListPipelineJobsResponseBodyJobs) *ListPipelineJobsResponseBody
	GetJobs() []*ListPipelineJobsResponseBodyJobs
	SetRequestId(v string) *ListPipelineJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineJobsResponseBody
	GetSuccess() *bool
}

type ListPipelineJobsResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPipelineJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineJobsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineJobsResponseBody) GetJobs() []*ListPipelineJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListPipelineJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineJobsResponseBody) SetErrorCode(v string) *ListPipelineJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetErrorMessage(v string) *ListPipelineJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetJobs(v []*ListPipelineJobsResponseBodyJobs) *ListPipelineJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListPipelineJobsResponseBody) SetRequestId(v string) *ListPipelineJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetSuccess(v bool) *ListPipelineJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineJobsResponseBodyJobs struct {
	// example:
	//
	// ss_saxsxsxs
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	JobName    *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// 123
	LastJobId *int64 `json:"lastJobId,omitempty" xml:"lastJobId,omitempty"`
	// example:
	//
	// {}
	LastJobParams *string `json:"lastJobParams,omitempty" xml:"lastJobParams,omitempty"`
}

func (s ListPipelineJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponseBodyJobs) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListPipelineJobsResponseBodyJobs) GetJobName() *string {
	return s.JobName
}

func (s *ListPipelineJobsResponseBodyJobs) GetLastJobId() *int64 {
	return s.LastJobId
}

func (s *ListPipelineJobsResponseBodyJobs) GetLastJobParams() *string {
	return s.LastJobParams
}

func (s *ListPipelineJobsResponseBodyJobs) SetIdentifier(v string) *ListPipelineJobsResponseBodyJobs {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetJobName(v string) *ListPipelineJobsResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetLastJobId(v int64) *ListPipelineJobsResponseBodyJobs {
	s.LastJobId = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetLastJobParams(v string) *ListPipelineJobsResponseBodyJobs {
	s.LastJobParams = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
