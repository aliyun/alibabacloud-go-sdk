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
	// 11A898F7-29D7-4AB3-B639-8BBDE671BBD5
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
