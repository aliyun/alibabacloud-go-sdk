// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitLiveTranscodeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitLiveTranscodeJobResponseBody
	GetRequestId() *string
}

type SubmitLiveTranscodeJobResponseBody struct {
	// The ID of the transcoding job.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitLiveTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveTranscodeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitLiveTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLiveTranscodeJobResponseBody) SetJobId(v string) *SubmitLiveTranscodeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitLiveTranscodeJobResponseBody) SetRequestId(v string) *SubmitLiveTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLiveTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
