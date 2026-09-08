// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetAgentJobResponseBodyJob) *GetAgentJobResponseBody
	GetJob() *GetAgentJobResponseBodyJob
	SetRequestId(v string) *GetAgentJobResponseBody
	GetRequestId() *string
}

type GetAgentJobResponseBody struct {
	// The task information.
	Job *GetAgentJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentJobResponseBody) GetJob() *GetAgentJobResponseBodyJob {
	return s.Job
}

func (s *GetAgentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentJobResponseBody) SetJob(v *GetAgentJobResponseBodyJob) *GetAgentJobResponseBody {
	s.Job = v
	return s
}

func (s *GetAgentJobResponseBody) SetRequestId(v string) *GetAgentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentJobResponseBodyJob struct {
	// The task ID.
	//
	// example:
	//
	// bc30a1080b21434f961a0d9a391b30b9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task output JSON string. Different tasks return different structures, which are defined by the business side.
	//
	// example:
	//
	// {\\"OssUri\\":\\"oss://ice-ai-saas/ice-ai-saas-prd/1123668546389636/210606863/generate/ag_3a506706a33f44008aec6274d2e38d58/\\"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The task status. Valid values:
	//
	// - Created
	//
	// - Queuing
	//
	// - Executing
	//
	// - Finished
	//
	// - Failed
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAgentJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetAgentJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetAgentJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetAgentJobResponseBodyJob) GetOutput() *string {
	return s.Output
}

func (s *GetAgentJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetAgentJobResponseBodyJob) SetJobId(v string) *GetAgentJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetAgentJobResponseBodyJob) SetOutput(v string) *GetAgentJobResponseBodyJob {
	s.Output = &v
	return s
}

func (s *GetAgentJobResponseBodyJob) SetStatus(v string) *GetAgentJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetAgentJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
