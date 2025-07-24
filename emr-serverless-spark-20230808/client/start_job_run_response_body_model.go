// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobRunId(v string) *StartJobRunResponseBody
	GetJobRunId() *string
	SetRequestId(v string) *StartJobRunResponseBody
	GetRequestId() *string
}

type StartJobRunResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// jr-54321
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *StartJobRunResponseBody) GetJobRunId() *string {
	return s.JobRunId
}

func (s *StartJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartJobRunResponseBody) SetJobRunId(v string) *StartJobRunResponseBody {
	s.JobRunId = &v
	return s
}

func (s *StartJobRunResponseBody) SetRequestId(v string) *StartJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartJobRunResponseBody) Validate() error {
	return dara.Validate(s)
}
