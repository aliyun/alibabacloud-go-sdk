// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCreativeGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitVideoCreativeGenerationJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *SubmitVideoCreativeGenerationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoCreativeGenerationJobResponseBody
	GetRequestId() *string
}

type SubmitVideoCreativeGenerationJobResponseBody struct {
	// The error code. This parameter is returned when the job is in the Failed state.
	//
	// example:
	//
	// WorkflowTaskFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The job ID.
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

func (s SubmitVideoCreativeGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCreativeGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) SetErrorCode(v string) *SubmitVideoCreativeGenerationJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) SetJobId(v string) *SubmitVideoCreativeGenerationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) SetRequestId(v string) *SubmitVideoCreativeGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
