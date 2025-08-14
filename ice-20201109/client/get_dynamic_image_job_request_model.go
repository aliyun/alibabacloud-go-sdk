// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicImageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetDynamicImageJobRequest
	GetJobId() *string
}

type GetDynamicImageJobRequest struct {
	// The job ID.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDynamicImageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobRequest) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetDynamicImageJobRequest) SetJobId(v string) *GetDynamicImageJobRequest {
	s.JobId = &v
	return s
}

func (s *GetDynamicImageJobRequest) Validate() error {
	return dara.Validate(s)
}
