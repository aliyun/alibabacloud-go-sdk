// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteLiveTranscodeJobRequest
	GetJobId() *string
}

type DeleteLiveTranscodeJobRequest struct {
	// The ID of the transcoding job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteLiveTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteLiveTranscodeJobRequest) SetJobId(v string) *DeleteLiveTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteLiveTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
