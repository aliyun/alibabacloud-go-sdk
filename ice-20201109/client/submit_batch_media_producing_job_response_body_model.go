// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchMediaProducingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitBatchMediaProducingJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitBatchMediaProducingJobResponseBody
	GetRequestId() *string
}

type SubmitBatchMediaProducingJobResponseBody struct {
	// The ID of the quick video production job.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitBatchMediaProducingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBatchMediaProducingJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitBatchMediaProducingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBatchMediaProducingJobResponseBody) SetJobId(v string) *SubmitBatchMediaProducingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitBatchMediaProducingJobResponseBody) SetRequestId(v string) *SubmitBatchMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBatchMediaProducingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
