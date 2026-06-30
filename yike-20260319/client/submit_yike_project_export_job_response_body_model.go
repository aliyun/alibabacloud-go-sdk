// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeProjectExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikeProjectExportJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikeProjectExportJobResponseBody
	GetRequestId() *string
}

type SubmitYikeProjectExportJobResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// task_abc123def456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// req_create_20260420_001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikeProjectExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeProjectExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikeProjectExportJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikeProjectExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikeProjectExportJobResponseBody) SetJobId(v string) *SubmitYikeProjectExportJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikeProjectExportJobResponseBody) SetRequestId(v string) *SubmitYikeProjectExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikeProjectExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
