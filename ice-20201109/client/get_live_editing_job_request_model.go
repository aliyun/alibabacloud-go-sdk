// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetLiveEditingJobRequest
	GetJobId() *string
}

type GetLiveEditingJobRequest struct {
	// The ID of the live editing job.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetLiveEditingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobRequest) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveEditingJobRequest) SetJobId(v string) *GetLiveEditingJobRequest {
	s.JobId = &v
	return s
}

func (s *GetLiveEditingJobRequest) Validate() error {
	return dara.Validate(s)
}
