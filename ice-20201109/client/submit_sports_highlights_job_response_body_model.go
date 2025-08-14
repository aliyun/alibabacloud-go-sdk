// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSportsHighlightsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSportsHighlightsJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSportsHighlightsJobResponseBody
	GetRequestId() *string
}

type SubmitSportsHighlightsJobResponseBody struct {
	// The ID of the sports highlights job.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSportsHighlightsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSportsHighlightsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSportsHighlightsJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSportsHighlightsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSportsHighlightsJobResponseBody) SetJobId(v string) *SubmitSportsHighlightsJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSportsHighlightsJobResponseBody) SetRequestId(v string) *SubmitSportsHighlightsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSportsHighlightsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
