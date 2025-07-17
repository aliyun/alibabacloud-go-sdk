// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobStatusRequest
	GetJobId() *string
}

type GetJobStatusRequest struct {
	// The ID of the asynchronous task that is generated after you call an asynchronous operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70ecdaec-bf21-4c11-8ecb-4f77453ceea8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
