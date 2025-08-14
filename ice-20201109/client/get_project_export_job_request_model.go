// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetProjectExportJobRequest
	GetJobId() *string
}

type GetProjectExportJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetProjectExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobRequest) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetProjectExportJobRequest) SetJobId(v string) *GetProjectExportJobRequest {
	s.JobId = &v
	return s
}

func (s *GetProjectExportJobRequest) Validate() error {
	return dara.Validate(s)
}
