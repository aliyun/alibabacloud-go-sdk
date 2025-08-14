// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScreenMediaHighlightsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitScreenMediaHighlightsJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitScreenMediaHighlightsJobResponseBody
	GetRequestId() *string
}

type SubmitScreenMediaHighlightsJobResponseBody struct {
	// The ID of the task.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitScreenMediaHighlightsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitScreenMediaHighlightsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitScreenMediaHighlightsJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitScreenMediaHighlightsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitScreenMediaHighlightsJobResponseBody) SetJobId(v string) *SubmitScreenMediaHighlightsJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobResponseBody) SetRequestId(v string) *SubmitScreenMediaHighlightsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
