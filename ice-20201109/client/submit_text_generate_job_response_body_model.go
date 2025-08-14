// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTextGenerateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitTextGenerateJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitTextGenerateJobResponseBody
	GetRequestId() *string
}

type SubmitTextGenerateJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTextGenerateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTextGenerateJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTextGenerateJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTextGenerateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTextGenerateJobResponseBody) SetJobId(v string) *SubmitTextGenerateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitTextGenerateJobResponseBody) SetRequestId(v string) *SubmitTextGenerateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTextGenerateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
