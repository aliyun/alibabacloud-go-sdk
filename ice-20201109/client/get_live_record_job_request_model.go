// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetLiveRecordJobRequest
	GetJobId() *string
}

type GetLiveRecordJobRequest struct {
	// The ID of the recording job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetLiveRecordJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveRecordJobRequest) SetJobId(v string) *GetLiveRecordJobRequest {
	s.JobId = &v
	return s
}

func (s *GetLiveRecordJobRequest) Validate() error {
	return dara.Validate(s)
}
