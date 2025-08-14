// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartHandleJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetSmartHandleJobRequest
	GetJobId() *string
}

type GetSmartHandleJobRequest struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetSmartHandleJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobRequest) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetSmartHandleJobRequest) SetJobId(v string) *GetSmartHandleJobRequest {
	s.JobId = &v
	return s
}

func (s *GetSmartHandleJobRequest) Validate() error {
	return dara.Validate(s)
}
