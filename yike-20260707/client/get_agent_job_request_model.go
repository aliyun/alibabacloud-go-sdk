// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAgentJobRequest
	GetJobId() *string
}

type GetAgentJobRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3B0E1A586AAC29742247******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAgentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentJobRequest) GoString() string {
	return s.String()
}

func (s *GetAgentJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAgentJobRequest) SetJobId(v string) *GetAgentJobRequest {
	s.JobId = &v
	return s
}

func (s *GetAgentJobRequest) Validate() error {
	return dara.Validate(s)
}
