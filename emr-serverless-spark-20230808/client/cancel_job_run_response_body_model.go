// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobRunId(v string) *CancelJobRunResponseBody
	GetJobRunId() *string
	SetRequestId(v string) *CancelJobRunResponseBody
	GetRequestId() *string
}

type CancelJobRunResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// jr-1a2bc3
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJobRunResponseBody) GetJobRunId() *string {
	return s.JobRunId
}

func (s *CancelJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelJobRunResponseBody) SetJobRunId(v string) *CancelJobRunResponseBody {
	s.JobRunId = &v
	return s
}

func (s *CancelJobRunResponseBody) SetRequestId(v string) *CancelJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelJobRunResponseBody) Validate() error {
	return dara.Validate(s)
}
