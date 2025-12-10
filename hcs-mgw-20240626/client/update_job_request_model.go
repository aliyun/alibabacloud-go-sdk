// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportJob(v *UpdateJobInfo) *UpdateJobRequest
	GetImportJob() *UpdateJobInfo
}

type UpdateJobRequest struct {
	// The details for updating the task.
	ImportJob *UpdateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) GetImportJob() *UpdateJobInfo {
	return s.ImportJob
}

func (s *UpdateJobRequest) SetImportJob(v *UpdateJobInfo) *UpdateJobRequest {
	s.ImportJob = v
	return s
}

func (s *UpdateJobRequest) Validate() error {
	if s.ImportJob != nil {
		if err := s.ImportJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
