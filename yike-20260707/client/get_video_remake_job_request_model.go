// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRemakeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoRemakeJobRequest
	GetJobId() *string
}

type GetVideoRemakeJobRequest struct {
	// The task ID returned by SubmitVideoRemakeJob.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoRemakeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRemakeJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoRemakeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoRemakeJobRequest) SetJobId(v string) *GetVideoRemakeJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoRemakeJobRequest) Validate() error {
	return dara.Validate(s)
}
