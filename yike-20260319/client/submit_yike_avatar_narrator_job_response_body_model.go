// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAvatarNarratorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeAvatarNarratorJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeAvatarNarratorJobResponseBody
	GetRequestId() *string
}

type SubmitYikeAvatarNarratorJobResponseBody struct {
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// req_create_20260420_001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeAvatarNarratorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAvatarNarratorJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeAvatarNarratorJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeAvatarNarratorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeAvatarNarratorJobResponseBody) SetJobId(v string) *SubmitYikeAvatarNarratorJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeAvatarNarratorJobResponseBody) SetRequestId(v string) *SubmitYikeAvatarNarratorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeAvatarNarratorJobResponseBody) Validate() error {
	return dara.Validate(s)
}
