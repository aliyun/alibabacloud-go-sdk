// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateEdgeTranscodeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type CreateEdgeTranscodeJobResponseBody struct {
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

func (s CreateEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeTranscodeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeTranscodeJobResponseBody) SetJobId(v string) *CreateEdgeTranscodeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateEdgeTranscodeJobResponseBody) SetRequestId(v string) *CreateEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
