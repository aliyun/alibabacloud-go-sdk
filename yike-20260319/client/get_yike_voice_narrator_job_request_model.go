// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeVoiceNarratorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeVoiceNarratorJobRequest
	GetJobId() *string
}

type GetYikeVoiceNarratorJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeVoiceNarratorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeVoiceNarratorJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikeVoiceNarratorJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeVoiceNarratorJobRequest) SetJobId(v string) *GetYikeVoiceNarratorJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeVoiceNarratorJobRequest) Validate() error {
	return dara.Validate(s)
}
