// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeAIAppJobRequest
	GetJobId() *string
}

type GetYikeAIAppJobRequest struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeAIAppJobRequest) SetJobId(v string) *GetYikeAIAppJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
