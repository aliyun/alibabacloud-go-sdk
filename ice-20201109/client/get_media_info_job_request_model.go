// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetMediaInfoJobRequest
	GetJobId() *string
}

type GetMediaInfoJobRequest struct {
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaInfoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaInfoJobRequest) SetJobId(v string) *GetMediaInfoJobRequest {
	s.JobId = &v
	return s
}

func (s *GetMediaInfoJobRequest) Validate() error {
	return dara.Validate(s)
}
