// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobRequest
	GetJobId() *string
}

type GetJobRequest struct {
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
