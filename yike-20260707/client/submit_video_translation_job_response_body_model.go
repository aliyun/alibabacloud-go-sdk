// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoTranslationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoTranslationJobResponseBody
	GetRequestId() *string
}

type SubmitVideoTranslationJobResponseBody struct {
	// The video translation job ID, used to call GetVideoTranslationJob to query the job.
	//
	// example:
	//
	// vtj_0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// req-vt-20260820-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoTranslationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoTranslationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoTranslationJobResponseBody) SetJobId(v string) *SubmitVideoTranslationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) SetRequestId(v string) *SubmitVideoTranslationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
