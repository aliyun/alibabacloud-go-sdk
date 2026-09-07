// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetextJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoDetextJobRequest
	GetJobId() *string
}

type GetVideoDetextJobRequest struct {
	// The video text erasure task ID returned by SubmitVideoDetextJob.
	//
	// This parameter is required.
	//
	// example:
	//
	// vdt_0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoDetextJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetextJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoDetextJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoDetextJobRequest) SetJobId(v string) *GetVideoDetextJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoDetextJobRequest) Validate() error {
	return dara.Validate(s)
}
