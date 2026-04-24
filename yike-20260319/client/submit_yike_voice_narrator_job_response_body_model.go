// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVoiceNarratorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeVoiceNarratorJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeVoiceNarratorJobResponseBody
	GetRequestId() *string
}

type SubmitYikeVoiceNarratorJobResponseBody struct {
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// req_create_20260420_001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeVoiceNarratorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVoiceNarratorJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeVoiceNarratorJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeVoiceNarratorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeVoiceNarratorJobResponseBody) SetJobId(v string) *SubmitYikeVoiceNarratorJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeVoiceNarratorJobResponseBody) SetRequestId(v string) *SubmitYikeVoiceNarratorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeVoiceNarratorJobResponseBody) Validate() error {
	return dara.Validate(s)
}
