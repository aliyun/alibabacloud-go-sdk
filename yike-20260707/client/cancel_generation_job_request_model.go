// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelGenerationJobRequest
	GetJobId() *string
}

type CancelGenerationJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ag_77b5f78b***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *CancelGenerationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelGenerationJobRequest) SetJobId(v string) *CancelGenerationJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
