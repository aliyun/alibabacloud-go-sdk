// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHighlightExtractionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitHighlightExtractionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitHighlightExtractionJobResponseBody
	GetRequestId() *string
}

type SubmitHighlightExtractionJobResponseBody struct {
	// The ID of the highlight extraction task.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitHighlightExtractionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHighlightExtractionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHighlightExtractionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitHighlightExtractionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHighlightExtractionJobResponseBody) SetJobId(v string) *SubmitHighlightExtractionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitHighlightExtractionJobResponseBody) SetRequestId(v string) *SubmitHighlightExtractionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHighlightExtractionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
