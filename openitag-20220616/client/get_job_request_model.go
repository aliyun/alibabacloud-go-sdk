// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobType(v string) *GetJobRequest
	GetJobType() *string
}

type GetJobRequest struct {
	// Task Type. Currently, only DOWNLOWD_MARKRESULT_FLOW is available.
	//
	// example:
	//
	// DOWNLOWD_MARKRESULT_FLOW
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *GetJobRequest) SetJobType(v string) *GetJobRequest {
	s.JobType = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
