// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteSmartJobRequest
	GetJobId() *string
}

type DeleteSmartJobRequest struct {
	// The IDs of the jobs to delete. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******042d5e4db6866f6289d1******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteSmartJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteSmartJobRequest) SetJobId(v string) *DeleteSmartJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteSmartJobRequest) Validate() error {
	return dara.Validate(s)
}
