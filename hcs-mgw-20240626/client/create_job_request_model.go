// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportJob(v *CreateJobInfo) *CreateJobRequest
	GetImportJob() *CreateJobInfo
}

type CreateJobRequest struct {
	// The details for creating the migration task.
	//
	// This parameter is required.
	ImportJob *CreateJobInfo `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetImportJob() *CreateJobInfo {
	return s.ImportJob
}

func (s *CreateJobRequest) SetImportJob(v *CreateJobInfo) *CreateJobRequest {
	s.ImportJob = v
	return s
}

func (s *CreateJobRequest) Validate() error {
	if s.ImportJob != nil {
		if err := s.ImportJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
