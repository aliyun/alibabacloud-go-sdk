// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeYikeStoryboardJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ResumeYikeStoryboardJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *ResumeYikeStoryboardJobResponseBody
	GetRequestId() *string
}

type ResumeYikeStoryboardJobResponseBody struct {
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeYikeStoryboardJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeYikeStoryboardJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeYikeStoryboardJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ResumeYikeStoryboardJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeYikeStoryboardJobResponseBody) SetJobId(v string) *ResumeYikeStoryboardJobResponseBody {
	s.JobId = &v
	return s
}

func (s *ResumeYikeStoryboardJobResponseBody) SetRequestId(v string) *ResumeYikeStoryboardJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeYikeStoryboardJobResponseBody) Validate() error {
	return dara.Validate(s)
}
