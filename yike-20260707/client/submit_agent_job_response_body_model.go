// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAgentJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAgentJobResponseBody
	GetRequestId() *string
}

type SubmitAgentJobResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// 35ed5e9588184f2e8d862ee07437971a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAgentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAgentJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAgentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAgentJobResponseBody) SetJobId(v string) *SubmitAgentJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAgentJobResponseBody) SetRequestId(v string) *SubmitAgentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAgentJobResponseBody) Validate() error {
	return dara.Validate(s)
}
