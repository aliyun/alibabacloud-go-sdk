// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetextJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoDetextJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoDetextJobResponseBody
	GetRequestId() *string
}

type SubmitVideoDetextJobResponseBody struct {
	// The video text erasure job ID. Use this ID to call GetVideoDetextJob to query the job.
	//
	// example:
	//
	// vdt_0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which is used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// req-detext-20260820-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoDetextJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetextJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetextJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoDetextJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoDetextJobResponseBody) SetJobId(v string) *SubmitVideoDetextJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoDetextJobResponseBody) SetRequestId(v string) *SubmitVideoDetextJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoDetextJobResponseBody) Validate() error {
	return dara.Validate(s)
}
