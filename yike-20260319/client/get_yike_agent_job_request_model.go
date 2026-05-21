// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeAgentJobRequest
	GetJobId() *string
}

type GetYikeAgentJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeAgentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeAgentJobRequest) SetJobId(v string) *GetYikeAgentJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeAgentJobRequest) Validate() error {
	return dara.Validate(s)
}
