// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeStoryboardJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeStoryboardJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeStoryboardJobResponseBody
	GetRequestId() *string
}

type SubmitYikeStoryboardJobResponseBody struct {
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeStoryboardJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeStoryboardJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeStoryboardJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeStoryboardJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeStoryboardJobResponseBody) SetJobId(v string) *SubmitYikeStoryboardJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeStoryboardJobResponseBody) SetRequestId(v string) *SubmitYikeStoryboardJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeStoryboardJobResponseBody) Validate() error {
	return dara.Validate(s)
}
