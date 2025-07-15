// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEdgeTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *StartEdgeTranscodeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *StartEdgeTranscodeJobResponseBody
	GetRequestId() *string
}

type StartEdgeTranscodeJobResponseBody struct {
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

func (s StartEdgeTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEdgeTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartEdgeTranscodeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *StartEdgeTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEdgeTranscodeJobResponseBody) SetJobId(v string) *StartEdgeTranscodeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *StartEdgeTranscodeJobResponseBody) SetRequestId(v string) *StartEdgeTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEdgeTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
