// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetJobRequest
	GetClusterId() *string
	SetJobId(v string) *GetJobRequest
	GetJobId() *string
}

type GetJobRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job ID. You can call the ListJobs operation to query the job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobRequest) SetClusterId(v string) *GetJobRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
