// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchMediaProducingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetBatchMediaProducingJobRequest
	GetJobId() *string
}

type GetBatchMediaProducingJobRequest struct {
	// The ID of the quick video production job.
	//
	// example:
	//
	// ****b4549d46c88681030f6e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetBatchMediaProducingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *GetBatchMediaProducingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetBatchMediaProducingJobRequest) SetJobId(v string) *GetBatchMediaProducingJobRequest {
	s.JobId = &v
	return s
}

func (s *GetBatchMediaProducingJobRequest) Validate() error {
	return dara.Validate(s)
}
