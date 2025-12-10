// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportJob(v *GetJobResp) *GetJobResponseBody
	GetImportJob() *GetJobResp
}

type GetJobResponseBody struct {
	// The details for obtaining the details of the migration task.
	ImportJob *GetJobResp `json:"ImportJob,omitempty" xml:"ImportJob,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetImportJob() *GetJobResp {
	return s.ImportJob
}

func (s *GetJobResponseBody) SetImportJob(v *GetJobResp) *GetJobResponseBody {
	s.ImportJob = v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	if s.ImportJob != nil {
		if err := s.ImportJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
