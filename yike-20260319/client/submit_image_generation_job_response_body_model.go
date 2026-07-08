// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitImageGenerationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitImageGenerationJobResponseBody
	GetRequestId() *string
}

type SubmitImageGenerationJobResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// ******3B0E1A586AAC29742247******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitImageGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageGenerationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitImageGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageGenerationJobResponseBody) SetJobId(v string) *SubmitImageGenerationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitImageGenerationJobResponseBody) SetRequestId(v string) *SubmitImageGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageGenerationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
