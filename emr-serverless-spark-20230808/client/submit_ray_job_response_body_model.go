// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRayJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitRayJobResponseBody
	GetRequestId() *string
	SetSubmissionId(v string) *SubmitRayJobResponseBody
	GetSubmissionId() *string
}

type SubmitRayJobResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// rj-k7nm8ahl5te4tg91
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
}

func (s SubmitRayJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitRayJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitRayJobResponseBody) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *SubmitRayJobResponseBody) SetRequestId(v string) *SubmitRayJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitRayJobResponseBody) SetSubmissionId(v string) *SubmitRayJobResponseBody {
	s.SubmissionId = &v
	return s
}

func (s *SubmitRayJobResponseBody) Validate() error {
	return dara.Validate(s)
}
