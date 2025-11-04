// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSegmentationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSegmentationJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSegmentationJobResponseBody
	GetRequestId() *string
}

type SubmitSegmentationJobResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSegmentationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSegmentationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSegmentationJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSegmentationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSegmentationJobResponseBody) SetJobId(v string) *SubmitSegmentationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSegmentationJobResponseBody) SetRequestId(v string) *SubmitSegmentationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSegmentationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
