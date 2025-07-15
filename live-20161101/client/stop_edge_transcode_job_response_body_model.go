// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *StopEdgeTranscodeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *StopEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type StopEdgeTranscodeJobResponseBody struct {
	// The ID of the edge transcoding task.
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

func (s StopEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopEdgeTranscodeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *StopEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopEdgeTranscodeJobResponseBody) SetJobId(v string) *StopEdgeTranscodeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *StopEdgeTranscodeJobResponseBody) SetRequestId(v string) *StopEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
