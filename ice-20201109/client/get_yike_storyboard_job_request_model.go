// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeStoryboardJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeStoryboardJobRequest
	GetJobId() *string
}

type GetYikeStoryboardJobRequest struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeStoryboardJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeStoryboardJobRequest) SetJobId(v string) *GetYikeStoryboardJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeStoryboardJobRequest) Validate() error {
	return dara.Validate(s)
}
