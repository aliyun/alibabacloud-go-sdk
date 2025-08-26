// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAsyncJobResultRequest
	GetJobId() *string
}

type GetAsyncJobResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// E75FE679-0303-4DD1-8252-1143B4FA8A27
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultRequest) Validate() error {
	return dara.Validate(s)
}
