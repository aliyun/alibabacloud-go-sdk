// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpDBDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitFpDBDeleteJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitFpDBDeleteJobResponseBody
	GetRequestId() *string
}

type SubmitFpDBDeleteJobResponseBody struct {
	// The ID of the job. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// d98459323c024947a104f6a50cbf****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4247B23C-26DE-529F-8D9F-FD6811AE979B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFpDBDeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpDBDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFpDBDeleteJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitFpDBDeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitFpDBDeleteJobResponseBody) SetJobId(v string) *SubmitFpDBDeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitFpDBDeleteJobResponseBody) SetRequestId(v string) *SubmitFpDBDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFpDBDeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
