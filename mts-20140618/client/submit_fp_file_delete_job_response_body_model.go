// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpFileDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitFpFileDeleteJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitFpFileDeleteJobResponseBody
	GetRequestId() *string
}

type SubmitFpFileDeleteJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D127C68E-F1A1-4CE5-A874-8FF724881A12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFpFileDeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpFileDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFpFileDeleteJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitFpFileDeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitFpFileDeleteJobResponseBody) SetJobId(v string) *SubmitFpFileDeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitFpFileDeleteJobResponseBody) SetRequestId(v string) *SubmitFpFileDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFpFileDeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
