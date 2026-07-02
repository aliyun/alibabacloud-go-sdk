// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVideoCloneJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeVideoCloneJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeVideoCloneJobResponseBody
	GetRequestId() *string
}

type SubmitYikeVideoCloneJobResponseBody struct {
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// req_create_20260420_001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeVideoCloneJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVideoCloneJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeVideoCloneJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeVideoCloneJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeVideoCloneJobResponseBody) SetJobId(v string) *SubmitYikeVideoCloneJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeVideoCloneJobResponseBody) SetRequestId(v string) *SubmitYikeVideoCloneJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeVideoCloneJobResponseBody) Validate() error {
	return dara.Validate(s)
}
