// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaProducingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetMediaProducingJobRequest
	GetJobId() *string
}

type GetMediaProducingJobRequest struct {
	// The ID of the media editing and production job.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaProducingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaProducingJobRequest) SetJobId(v string) *GetMediaProducingJobRequest {
	s.JobId = &v
	return s
}

func (s *GetMediaProducingJobRequest) Validate() error {
	return dara.Validate(s)
}
