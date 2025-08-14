// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetLiveTranscodeJobRequest
	GetJobId() *string
}

type GetLiveTranscodeJobRequest struct {
	// The ID of the transcoding job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetLiveTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveTranscodeJobRequest) SetJobId(v string) *GetLiveTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *GetLiveTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
