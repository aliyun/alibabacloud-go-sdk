// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateEdgeTranscodeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type UpdateEdgeTranscodeJobResponseBody struct {
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

func (s UpdateEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEdgeTranscodeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEdgeTranscodeJobResponseBody) SetJobId(v string) *UpdateEdgeTranscodeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobResponseBody) SetRequestId(v string) *UpdateEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
